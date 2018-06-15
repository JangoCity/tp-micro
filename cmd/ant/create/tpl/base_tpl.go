// Code generated by go-bindata.
// sources:
// .ant_gen_lock
// .gitignore
// README.md
// api/handler.go
// api/router.go
// args/const.go
// args/type.go
// args/var.go
// rerrs/rerrs.go
// sdk/rpc.go
// sdk/rpc_test.go
// DO NOT EDIT!

package tpl

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/xiaoenai/ants/cmd/ant/info"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _Ant_gen_lock = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xca\xb1\x0d\x02\x31\x0c\x05\xd0\x9e\x29\x3e\x37\xc0\x0d\x40\x4f\x0b\x4d\x06\x48\x14\xec\xc4\x52\x62\x4b\xd8\x40\xd8\x1e\x71\xed\xd3\xbb\xdd\xd3\xf5\x72\x42\xea\xe2\xc8\x45\x03\x8d\x34\xa3\xda\x9c\x45\x1f\x30\x1d\x5f\x54\x7b\xd3\xd3\xc1\x32\xc8\xf1\x91\xe8\x88\x4e\xd8\xf6\x46\xba\x37\xdb\xe0\x2f\x66\x59\x10\x3e\xfc\xdf\x40\x4b\x3c\xfc\xfc\x0b\x00\x00\xff\xff\x56\x3f\x50\x79\x5d\x00\x00\x00")

func Ant_gen_lockBytes() ([]byte, error) {
	return bindataRead(
		_Ant_gen_lock,
		".ant_gen_lock",
	)
}

func Ant_gen_lock() (*asset, error) {
	bytes, err := Ant_gen_lockBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".ant_gen_lock", size: 93, mode: os.FileMode(420), modTime: time.Unix(1529040143, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _Gitignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8e\xc1\x0a\xc2\x30\x10\x44\xef\xf3\x29\x7b\x18\x50\x50\xfc\x17\x91\xd0\xa6\x69\xac\xd8\x6e\x4c\x52\xa9\x3d\xf8\xed\x92\xa6\x5e\x66\x77\x1e\xcb\xec\x08\x15\xc2\x06\xc2\xa4\x30\xda\x3e\x60\xb2\x4b\x19\xc2\xeb\xe9\x7c\x79\xbf\x6e\xd8\x27\x75\x2e\xd4\x7a\x3d\xd0\x6b\xdd\x8e\xb4\x30\xd6\xab\xe9\x5c\x3f\x4f\x7f\xe3\x35\x7f\x82\x4b\xe5\x6a\xf3\x6e\x09\x1a\x33\xa5\x26\x8f\xcd\x30\xd5\x00\xb7\xb8\xaa\x5f\x08\xf7\xa7\x21\x6a\x0f\x61\x6c\x22\x84\xeb\x10\x20\xf4\x6b\xe1\xa9\x83\xb0\x1d\x8b\xda\xde\x17\x12\xf2\x02\xe1\x53\x3d\x64\xd2\xfb\x1c\xb6\x82\xbf\x00\x00\x00\xff\xff\xbf\xaa\x91\x77\xcf\x00\x00\x00")

func GitignoreBytes() ([]byte, error) {
	return bindataRead(
		_Gitignore,
		".gitignore",
	)
}

func Gitignore() (*asset, error) {
	bytes, err := GitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".gitignore", size: 207, mode: os.FileMode(420), modTime: time.Unix(1523605812, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _readmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xa8\xae\x0e\x08\xf2\xf7\x8a\xf7\x73\xf4\x75\xad\xad\xe5\x02\x04\x00\x00\xff\xff\x40\x0d\xb5\xec\x10\x00\x00\x00")

func readmeMdBytes() ([]byte, error) {
	return bindataRead(
		_readmeMd,
		"README.md",
	)
}

func readmeMd() (*asset, error) {
	bytes, err := readmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "README.md", size: 16, mode: os.FileMode(420), modTime: time.Unix(1523605812, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _apiHandlerGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x48\x2c\xc8\xe4\x02\x04\x00\x00\xff\xff\x0c\x0c\x0a\x62\x0c\x00\x00\x00")

func apiHandlerGoBytes() ([]byte, error) {
	return bindataRead(
		_apiHandlerGo,
		"api/handler.go",
	)
}

func apiHandlerGo() (*asset, error) {
	bytes, err := apiHandlerGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api/handler.go", size: 12, mode: os.FileMode(420), modTime: time.Unix(1529044481, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _apiRouterGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8b\x41\x0e\x02\x21\x0c\x45\xd7\xf6\x14\xcd\xac\x66\x5c\x0c\x89\x47\xf1\x06\x88\x15\x88\x40\x9b\x52\x16\xc6\x78\x77\x83\xba\x70\xf9\xdf\x7f\x4f\x7c\xb8\xfb\x48\xe8\x25\x03\xe4\x2a\xac\x86\x2b\x1c\x4c\x70\x89\xd9\xd2\xb8\xec\x81\xab\x4b\xd4\xf4\x51\x88\x4e\xa1\x39\xa3\x42\x53\x5b\x60\x03\x70\x0e\xc3\xe8\xc6\xf5\xcc\xc3\x08\x95\x62\xee\x46\xda\x7f\x14\x93\x6f\xd7\x32\xb7\x31\xea\x54\x74\x87\xdb\x68\xe1\xbf\x5a\xbf\x07\x1e\x4d\xf6\x0f\xd0\x0d\x9f\xf0\x82\x77\x00\x00\x00\xff\xff\xd3\xb5\x7a\x55\x9b\x00\x00\x00")

func apiRouterGoBytes() ([]byte, error) {
	return bindataRead(
		_apiRouterGo,
		"api/router.go",
	)
}

func apiRouterGo() (*asset, error) {
	bytes, err := apiRouterGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api/router.go", size: 155, mode: os.FileMode(420), modTime: time.Unix(1523605812, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _argsConstGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x48\x2c\x4a\x2f\xe6\xe2\x4a\xce\xcf\x2b\x2e\x51\xd0\xd0\xe4\x02\x04\x00\x00\xff\xff\xf5\x9a\x10\x2f\x17\x00\x00\x00")

func argsConstGoBytes() ([]byte, error) {
	return bindataRead(
		_argsConstGo,
		"args/const.go",
	)
}

func argsConstGo() (*asset, error) {
	bytes, err := argsConstGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "args/const.go", size: 23, mode: os.FileMode(420), modTime: time.Unix(1529039639, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _argsTypeGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x48\x2c\x4a\x2f\xe6\xe2\x2a\xa9\x2c\x48\x55\xd0\xd0\xe4\x02\x04\x00\x00\xff\xff\x61\x1b\x80\x25\x16\x00\x00\x00")

func argsTypeGoBytes() ([]byte, error) {
	return bindataRead(
		_argsTypeGo,
		"args/type.go",
	)
}

func argsTypeGo() (*asset, error) {
	bytes, err := argsTypeGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "args/type.go", size: 22, mode: os.FileMode(420), modTime: time.Unix(1529039644, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _argsVarGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x48\x2c\x4a\x2f\xe6\xe2\x2a\x4b\x2c\x52\xd0\xd0\xe4\x02\x04\x00\x00\xff\xff\xa5\xca\xdc\xfb\x15\x00\x00\x00")

func argsVarGoBytes() ([]byte, error) {
	return bindataRead(
		_argsVarGo,
		"args/var.go",
	)
}

func argsVarGo() (*asset, error) {
	bytes, err := argsVarGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "args/var.go", size: 21, mode: os.FileMode(420), modTime: time.Unix(1529039650, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rerrsRerrsGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xca\xb1\xce\x82\x30\x10\x00\xe0\x99\x7b\x8a\x4b\x27\x48\x08\x2d\xff\xfe\x4f\x4e\x2e\xc6\xf0\x06\x27\x5e\xa0\x11\xda\x7a\x3d\x30\xbe\xbd\xc1\xa0\x93\xf3\xf7\x25\xea\x6f\x34\x30\x0a\x8b\x64\x00\x3f\xa7\x28\x8a\x25\x14\x9a\xd0\x0c\x5e\xc7\xe5\xd2\xf4\x71\xb6\x23\x07\x79\x4e\xcc\x7f\x7d\xb0\xca\x13\x6f\xcd\x40\x05\xb0\x92\x6c\xdd\x5a\xec\x58\xe4\x18\x56\x9a\xfc\xf5\x4c\x42\x33\x2b\x0b\xb2\x48\x14\x28\x7e\xda\x3f\x6a\x6a\x4e\xfc\xe8\xde\xa9\x6c\x9d\x73\xae\xad\xd1\xec\x11\xbf\xd3\xd4\x68\x0e\x31\x28\xf9\x90\xd1\xef\x2a\x7c\x5f\x38\x2b\xa6\xcf\xca\xa6\x82\x0a\x5e\x01\x00\x00\xff\xff\x8b\x81\xb9\xe7\xd1\x00\x00\x00")

func rerrsRerrsGoBytes() ([]byte, error) {
	return bindataRead(
		_rerrsRerrsGo,
		"rerrs/rerrs.go",
	)
}

func rerrsRerrsGo() (*asset, error) {
	bytes, err := rerrsRerrsGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rerrs/rerrs.go", size: 209, mode: os.FileMode(420), modTime: time.Unix(1523605812, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sdkRpcGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x4e\xc9\xe6\x02\x04\x00\x00\xff\xff\x36\xfa\x03\xb1\x0c\x00\x00\x00")

func sdkRpcGoBytes() ([]byte, error) {
	return bindataRead(
		_sdkRpcGo,
		"sdk/rpc.go",
	)
}

func sdkRpcGo() (*asset, error) {
	bytes, err := sdkRpcGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sdk/rpc.go", size: 12, mode: os.FileMode(420), modTime: time.Unix(1523605812, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sdkRpc_testGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x4e\xc9\xe6\x02\x04\x00\x00\xff\xff\x36\xfa\x03\xb1\x0c\x00\x00\x00")

func sdkRpc_testGoBytes() ([]byte, error) {
	return bindataRead(
		_sdkRpc_testGo,
		"sdk/rpc_test.go",
	)
}

func sdkRpc_testGo() (*asset, error) {
	bytes, err := sdkRpc_testGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sdk/rpc_test.go", size: 12, mode: os.FileMode(420), modTime: time.Unix(1523605812, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	".ant_gen_lock":   Ant_gen_lock,
	".gitignore":      Gitignore,
	"README.md":       readmeMd,
	"api/handler.go":  apiHandlerGo,
	"api/router.go":   apiRouterGo,
	"args/const.go":   argsConstGo,
	"args/type.go":    argsTypeGo,
	"args/var.go":     argsVarGo,
	"rerrs/rerrs.go":  rerrsRerrsGo,
	"sdk/rpc.go":      sdkRpcGo,
	"sdk/rpc_test.go": sdkRpc_testGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	".ant_gen_lock": &bintree{Ant_gen_lock, map[string]*bintree{}},
	".gitignore":    &bintree{Gitignore, map[string]*bintree{}},
	"README.md":     &bintree{readmeMd, map[string]*bintree{}},
	"api": &bintree{nil, map[string]*bintree{
		"handler.go": &bintree{apiHandlerGo, map[string]*bintree{}},
		"router.go":  &bintree{apiRouterGo, map[string]*bintree{}},
	}},
	"args": &bintree{nil, map[string]*bintree{
		"const.go": &bintree{argsConstGo, map[string]*bintree{}},
		"type.go":  &bintree{argsTypeGo, map[string]*bintree{}},
		"var.go":   &bintree{argsVarGo, map[string]*bintree{}},
	}},
	"rerrs": &bintree{nil, map[string]*bintree{
		"rerrs.go": &bintree{rerrsRerrsGo, map[string]*bintree{}},
	}},
	"sdk": &bintree{nil, map[string]*bintree{
		"rpc.go":      &bintree{sdkRpcGo, map[string]*bintree{}},
		"rpc_test.go": &bintree{sdkRpc_testGo, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	data = bytes.Replace(data, projNameTpl, projNameBytes, -1)
	data = bytes.Replace(data, projPathTpl, projPathBytes, -1)
	if strings.HasSuffix(name, ".go") {
		data, _ = format.Source(data)
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	fmt.Printf("generate %s\n", string(projPathBytes)+"/"+_filePath(dir, name))
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		if name == ".ant_gen_lock" {
			return nil
		}
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		if child == "logic" {
			continue
		}
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

var (
	projNameBytes []byte
	projPathBytes []byte
	projNameTpl   = []byte("{{PROJ_NAME}}")
	projPathTpl   = []byte("{{PROJ_PATH}}")
)

// Create creates base files.
func Create() {
	projNameBytes = []byte(info.ProjName())
	projPathBytes = []byte(info.ProjPath())
	RestoreAssets("./", "")
}
