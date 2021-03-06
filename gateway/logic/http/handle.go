// Copyright 2018 github.com/xiaoenai. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/xiaoenai/tp-micro/clientele"

	ws "github.com/fasthttp-contrib/websocket"
	"github.com/henrylee2cn/goutil"
	tp "github.com/henrylee2cn/teleport"
	"github.com/henrylee2cn/teleport/codec"
	"github.com/henrylee2cn/teleport/plugin/proxy"
	"github.com/valyala/fasthttp"
	"github.com/xiaoenai/tp-micro/gateway/logic"
	"github.com/xiaoenai/tp-micro/gateway/logic/hosts"
	"github.com/xiaoenai/tp-micro/gateway/logic/socket"
)

const (
	// SEQ the packet sequence query parameter specified by the client
	SEQ = "seq_"
)

var (
	allowCross bool
	gwHostsUri string
)

func handler(ctx *fasthttp.RequestCtx) {
	(&requestHandler{ctx: ctx}).handle()
}

type requestHandler struct {
	ctx    *fasthttp.RequestCtx
	errMsg []byte
}

var rerrInternalServerError = tp.NewRerror(tp.CodeInternalServerError, tp.CodeText(tp.CodeInternalServerError), "")

var (
	gwWsUri    string
	wsUpgrader = ws.Custom(wsHandler, 4096, 4096)
)

type wsFastHttpConn struct {
	*ws.Conn
	rio, wio    sync.Mutex
	frameReader io.Reader
}

// Read reads data from the connection.
// Read can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetReadDeadline.
func (w *wsFastHttpConn) Read(msg []byte) (n int, err error) {
	w.rio.Lock()
	defer w.rio.Unlock()
again:
	if w.frameReader == nil {
		_, w.frameReader, err = w.Conn.NextReader()
		if err != nil {
			return 0, err
		}
		if w.frameReader == nil {
			goto again
		}
	}
	n, err = w.frameReader.Read(msg)
	if err == io.EOF {
		w.frameReader = nil
		goto again
	}
	return n, err
}

// Write implements the io.Writer interface:
// it writes data as a frame to the WebSocket connection.
func (w *wsFastHttpConn) Write(msg []byte) (n int, err error) {
	w.wio.Lock()
	defer w.wio.Unlock()
	err = w.Conn.WriteMessage(ws.TextMessage, msg)
	return len(msg), err
}

// SetDeadline sets the connection's network read & write deadlines.
func (w *wsFastHttpConn) SetDeadline(t time.Time) error {
	err := w.Conn.SetReadDeadline(t)
	if err != nil {
		return err
	}
	return w.Conn.SetWriteDeadline(t)
}

func wsHandler(conn *ws.Conn) {
	socket.OuterServeConn(&wsFastHttpConn{Conn: conn})
}

func (r *requestHandler) handle() {
	var ctx = r.ctx
	var uri = goutil.BytesToString(ctx.Path())
	// websocket
	if uri == gwWsUri {
		err := wsUpgrader.Upgrade(ctx)
		if err != nil {
			tp.Debugf("upgrade websocket fail: %s", err.Error())
		}
		return
	}
	var h = r.Header()
	var contentType = goutil.BytesToString(h.ContentType())
	var bodyCodec = GetBodyCodec(contentType, codec.ID_PLAIN)
	var acceptBodyCodec = GetBodyCodec(goutil.BytesToString(h.Peek("Accept")), bodyCodec)
	var query = r.ctx.QueryArgs()
	var bodyBytes = ctx.Request.Body()
	var reply []byte
	var label proxy.ProxyLabel
	label.Uri = uri

	// set real ip
	if xRealIp := h.Peek("X-Real-IP"); len(xRealIp) > 0 {
		label.RealIp = string(xRealIp)
	} else if xForwardedFor := h.Peek("X-Forwarded-For"); len(xForwardedFor) > 0 {
		label.RealIp = string(bytes.Split(xForwardedFor, []byte{','})[0])
	}
	if len(label.RealIp) == 0 {
		label.RealIp = ctx.RemoteAddr().String()
	}
	start := time.Now()
	defer func() {
		if p := recover(); p != nil {
			r.replyError(rerrInternalServerError.Copy().SetReason(fmt.Sprint(p)))
		}
		r.runlog(start, &label, goutil.BytesToString(query.Peek(SEQ)), bodyBytes, &reply)
	}()

	// cross
	if allowCross && r.crossDomainFilter() {
		return
	}

	// gw hosts
	if label.Uri == gwHostsUri {
		switch acceptBodyCodec {
		case codec.ID_PROTOBUF:
			b, _ := codec.ProtoMarshal(hosts.GwHosts())
			r.ctx.Success("application/x-protobuf", b)
		default:
			b, _ := json.Marshal(hosts.GwHosts())
			r.ctx.Success("application/json", b)
		}
		return
	}

	// verify access token
	accessToken, settings, rerr := logic.HttpHooks().OnRequest(r, bodyBytes, logic.AuthFunc())
	if rerr != nil {
		r.replyError(rerr)
		return
	}

	// set header
	h.VisitAll(func(key, value []byte) {
		settings = append(settings, tp.WithAddMeta(string(key), string(value)))
	})

	// set body codec
	settings = append(settings, tp.WithBodyCodec(bodyCodec))

	// set accept body codec
	if acceptBodyCodec != bodyCodec {
		settings = append(settings, tp.WithAcceptBodyCodec(acceptBodyCodec))
	}

	// set session id
	if accessToken == nil {
		label.SessionId = ctx.RemoteAddr().String()
	} else {
		label.SessionId = accessToken.SessionId()
		if info := accessToken.AddedQuery(); info != nil {
			info.VisitAll(func(key, value []byte) {
				query.AddBytesKV(key, value)
			})
		}
	}

	settings = append(settings, tp.WithAddMeta(tp.MetaRealIp, label.RealIp))

	// set seq
	if seqBytes := query.Peek(SEQ); len(seqBytes) > 0 {
		settings = append(settings, tp.WithSeq(clientele.GetSeq(label.RealIp+"@"+goutil.BytesToString(seqBytes))))
	} else {
		settings = append(settings, tp.WithSeq(clientele.GetSeq(label.RealIp)))
	}

	if query.Len() > 0 {
		label.Uri += "?" + query.String()
	}

	callcmd := logic.
		ProxySelector(&label).
		Call(label.Uri, bodyBytes, &reply, settings...)

	// fail
	if rerr := callcmd.Rerror(); rerr != nil {
		callcmd.InputMeta().VisitAll(func(key, value []byte) {
			k := goutil.BytesToString(key)
			v := goutil.BytesToString(value)
			ctx.Response.Header.Add(k, v)
		})
		r.replyError(rerr)
		return
	}

	// succ

	var hasRespContentType bool
	callcmd.InputMeta().VisitAll(func(key, value []byte) {
		k := goutil.BytesToString(key)
		v := goutil.BytesToString(value)
		if k == "Content-Type" {
			hasRespContentType = true
			ctx.Response.Header.Set(k, v)
		} else {
			ctx.Response.Header.Add(k, v)
		}
	})
	if !hasRespContentType {
		ctx.Response.Header.Add(
			"Content-Type",
			GetContentType(callcmd.InputBodyCodec(), contentType),
		)
	}
	ctx.SetBody(reply)
}

var optionsBytes = []byte("OPTIONS")

func (r *requestHandler) crossDomainFilter() bool {
	r.ctx.Response.Header.Set("Access-Control-Allow-Origin", string(r.ctx.Request.Header.Peek("Origin")))
	r.ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	r.ctx.Response.Header.Set("Access-Control-Allow-Methods", string(r.ctx.Request.Header.Peek("Access-Control-Request-Method")))
	// r.ctx.Response.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	r.ctx.Response.Header.Set("Access-Control-Allow-Headers", string(r.ctx.Request.Header.Peek("Access-Control-Request-Headers")))
	// r.ctx.Response.Header.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	r.ctx.Response.Header.Set("Access-Control-Max-Age", "-1")
	// r.ctx.Response.Header.Set("Access-Control-Max-Age", "172800")
	if !bytes.Equal(r.ctx.Method(), optionsBytes) {
		return false
	}
	r.ctx.SetStatusCode(204)
	return true
}

func (r *requestHandler) replyError(rerr *tp.Rerror) {
	var statusCode int
	if rerr.Code < 200 {
		// Internal communication error
		statusCode = 500
	} else if rerr.Code < 600 {
		// Custom HTTP error
		statusCode = int(rerr.Code)
	} else {
		// Business error
		statusCode = 299
	}
	r.errMsg, _ = rerr.MarshalJSON()
	r.ctx.SetStatusCode(statusCode)
	r.ctx.SetContentType("application/json")
	r.ctx.SetBody(r.errMsg)
}

// QueryArgs returns the query arguments object of request.
func (r *requestHandler) QueryArgs() *fasthttp.Args {
	return r.ctx.QueryArgs()
}

// Header returns the header object of request.
func (r *requestHandler) Header() *fasthttp.RequestHeader {
	return &r.ctx.Request.Header
}
