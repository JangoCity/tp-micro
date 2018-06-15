# Ants [![GitHub release](https://img.shields.io/github/release/xiaoenai/ants.svg)](https://github.com/xiaoenai/ants/releases) [![report card](https://goreportcard.com/badge/github.com/xiaoenai/ants)](http://goreportcard.com/report/xiaoenai/ants) [![github issues](https://img.shields.io/github/issues/xiaoenai/ants.svg)](https://github.com/xiaoenai/ants/issues?q=is%3Aopen+is%3Aissue) [![github closed issues](https://img.shields.io/github/issues-closed-raw/xiaoenai/ants.svg)](https://github.com/xiaoenai/ants/issues?q=is%3Aissue+is%3Aclosed) [![view teleport](https://img.shields.io/badge/based%20on-teleport-00BCD4.svg)](https://github.com/henrylee2cn/teleport) [![view tp-micro](https://img.shields.io/badge/based%20on-tp--micro-00BCD4.svg)](https://github.com/henrylee2cn/tp-micro) [![view Go网络编程群](https://img.shields.io/badge/官方QQ群-Go网络编程(42730308)-27a5ea.svg)](http://jq.qq.com/?_wv=1027&k=fzi4p1)


Ants is a highly available micro service platform based on [TP-Micro](https://github.com/henrylee2cn/tp-micro) and [Teleport](https://github.com/henrylee2cn/teleport).

[简体中文](https://github.com/xiaoenai/ants/blob/master/README_ZH.md)


## Install


```
go version ≥ 1.9
```

```sh
go get -u -f -d github.com/xiaoenai/ants/...
cd $GOPATH/src/github.com/xiaoenai/ants/cmd/ant
go install
```

## Feature

- Support auto service-discovery
- Supports custom service linker
- Support load balancing
- Support NIO and connection pool
- Support custom protocol
- Support custom body codec
- Support plug-in expansion
- Support heartbeat mechanism
- Detailed log information, support print input and output details
- Support for setting slow operation alarm thresholds
- Support for custom log
- Support smooth shutdown and update
- Support push handler
- Support network list: `tcp`, `tcp4`, `tcp6`, `unix`, `unixpacket` and so on
- Client support automatically redials after disconnection
- Circuit breaker for overload protection


## Project Management

### Install Ant Command

```sh
cd $GOPATH/src/github.com/xiaoenai/ants/cmd/ant
go install
```

### Generate project

`ant gen` command help:

```
NAME:
     ant gen - Generate an ant project

USAGE:
     ant gen [command options] [arguments...]

OPTIONS:
     --template value, -t value    The template for code generation(relative/absolute)
     --app_path value, -p value  The path(relative/absolute) of the project
```

example: `ant gen -t ./__ant__tpl__.go -p ./myant` or default `ant gen myant`

- template file `__ant__tpl__.go` demo:

```go
// package __ANT__TPL__ is the project template
package __ANT__TPL__

// __API__PULL__ register PULL router:
//  /home
//  /math/divide
type __API__PULL__ interface {
	Home(*struct{}) *HomeResult
	Math
}

// __API__PUSH__ register PUSH router:
//  /stat
type __API__PUSH__ interface {
	Stat(*StatArg)
}

// MODEL create model
type __MODEL__ struct {
	DivideArg
	User
}

// Math controller
type Math interface {
	// Divide handler
	Divide(*DivideArg) *DivideResult
}

// HomeResult home result
type HomeResult struct {
	Content string // text
}

type (
	// DivideArg divide api arg
	DivideArg struct {
		// dividend
		A float64
		// divisor
		B float64 `param:"<range: 0.01:100000>"`
	}
	// DivideResult divide api result
	DivideResult struct {
		// quotient
		C float64
	}
)

// StatArg stat handler arg
type StatArg struct {
	Ts int64 // timestamps
}

// User user info
type User struct {
	Id   int64
	Name string
	Age  int32
}
```

- The template generated by `ant gen` command.

```
├── .ant_gen_lock
├── .gitignore
├── README.md
├── __ant__tpl__.go
├── api
│   ├── handler.gen.go
│   ├── handler.go
│   ├── router.gen.go
│   └── router.go
├── args
│   ├── const.go
│   ├── type.gen.go
│   ├── type.go
│   └── var.go
├── config
│   └── config.yaml
├── config.go
├── demo
├── log
│   └── PID
├── logic
│   ├── model
│   │   ├── divide_arg.gen.go
│   │   ├── init.go
│   │   └── user.gen.go
│   └── tmp_code.gen.go
├── main.go
├── rerrs
│   └── rerrs.go
└── sdk
    ├── rpc.gen.go
    ├── rpc.gen_test.go
    ├── rpc.go
    └── rpc_test.go
```

Desc:

- This `ant gen` command only covers files with the ".gen.go" suffix if the `.ant_gen_lock` file exists
- Add `.gen` suffix to the file name of the automatically generated file
- `tmp_code.gen.go` is temporary code used to ensure successful compilation!<br>When the project is completed, it should be removed!

[Generated Default Sample](https://github.com/henrylee2cn/tp-micro/tree/master/examples/sample)

### Run project

`ant run` command help:

```
NAME:
     ant run - Compile and run gracefully (monitor changes) an any existing go project

USAGE:
     ant run [options] [arguments...]
 or
     ant run [options except -app_path] [arguments...] {app_path}

OPTIONS:
     --watch_exts value, -x value  Specified to increase the listening file suffix (default: ".go", ".ini", ".yaml", ".toml", ".xml")
     --notwatch value, -n value    Not watch files or directories
     --app_path value, -p value    The path(relative/absolute) of the project
```

example: `ant run -x .yaml -p myant` or `ant run`

[More Ant Command](https://github.com/xiaoenai/ants/tree/master/cmd/ant)

## Demo


- server

```go
package main

import (
        micro "github.com/henrylee2cn/tp-micro"
        tp "github.com/henrylee2cn/teleport"
)

// Args args
type Args struct {
        A int
        B int `param:"<range:1:>"`
}

// P handler
type P struct {
        tp.PullCtx
}

// Divide divide API
func (p *P) Divide(args *Args) (int, *tp.Rerror) {
        return args.A / args.B, nil
}

func main() {
        srv := micro.NewServer(micro.SrvConfig{
                ListenAddress: ":9090",
        })
        srv.RoutePull(new(P))
        srv.Listen()
}
```

- client

```go
package main

import (
        micro "github.com/henrylee2cn/tp-micro"
    tp "github.com/henrylee2cn/teleport"
)

func main() {
        cli := micro.NewClient(
                micro.CliConfig{},
                micro.NewStaticLinker(":9090"),
        )
        defer   cli.Close()

        type Args struct {
                A int
                B int
        }

        var reply int
        rerr := cli.Pull("/p/divide", &Args{
                A: 10,
                B: 2,
        }, &reply).Rerror()
        if rerr != nil {
                tp.Fatalf("%v", rerr)
        }
        tp.Infof("10/2=%d", reply)
        rerr = cli.Pull("/p/divide", &Args{
                A: 10,
                B: 0,
        }, &reply).Rerror()
        if rerr == nil {
                tp.Fatalf("%v", rerr)
        }
        tp.Infof("test binding error: ok: %v", rerr)
}
```


[More](https://github.com/henrylee2cn/tp-micro/tree/master/samples)

## License

Ants is under Apache v2 License. See the [LICENSE](https://github.com/xiaoenai/ants/raw/master/LICENSE) file for the full license text
