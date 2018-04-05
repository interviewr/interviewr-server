// Code generated by go-bindata. DO NOT EDIT.
// sources:
// migrations/1_init.sql
// migrations/migrations.go

package repository


import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataMigrations1initsql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x41\x4b\xc3\x40\x10\x85\xef\xf3\x2b\xde\xad\x0d\x1a\x88\x42\x4e" +
	"\x3d\xc5\x36\xa0\x18\xdb\xba\x26\x42\x8f\x6b\x76\x68\x07\x37\xbb\x4b\x36\x52\xf1\xd7\x4b\x2a\xd5\x16\x3c\xf4\x3a" +
	"\xdf\xf7\x60\xde\x4b\x53\x5c\x75\xb2\xed\xf5\xc0\x68\x02\xa5\x29\x5e\x9e\x2b\x88\x43\xe4\x76\x10\xef\x30\x69\xc2" +
	"\x04\x12\xc1\x9f\xdc\x7e\x0c\x6c\xb0\xdf\xb1\xc3\xb0\x93\x88\x9f\xdc\x28\x49\x84\x0e\xc1\x0a\x1b\x9a\xab\xb2\xa8" +
	"\x4b\xd4\xc5\x5d\x55\xc2\xf7\x5b\xed\xe4\xeb\x20\xd1\x94\x00\x31\x78\x2d\xd4\xfc\xbe\x50\xd3\xdb\x3c\x4b\xb0\x56" +
	"\x0f\x4f\x85\xda\xe0\xb1\xdc\x5c\x13\xe0\x74\xc7\x7f\x42\x96\x60\xb9\xaa\xb1\x6c\xaa\x6a\x84\xdc\x69\xb1\x27\xf1" +
	"\x73\x6a\x38\xb6\xbd\x84\xc3\x3b\x47\xe7\x26\xcf\x92\x91\x59\xdf\xea\x33\x90\x67\x09\x25\x33\x22\x3a\xad\xbf\xf0" +
	"\x7b\x77\x1c\xe0\xb7\xfd\x78\xbc\xa8\x7f\xef\xad\x65\x83\x37\xdd\xbe\xd3\x42\xad\xd6\xff\x2c\x30\xa3\xef\x00\x00" +
	"\x00\xff\xff\xbf\x88\x9d\x43\x6e\x01\x00\x00")

func bindataMigrations1initsqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataMigrations1initsql,
		"migrations/1_init.sql",
	)
}



func bindataMigrations1initsql() (*asset, error) {
	bytes, err := bindataMigrations1initsqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "migrations/1_init.sql",
		size: 366,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1522857656, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataMigrationsMigrationsgo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindataMigrationsMigrationsgoBytes() ([]byte, error) {
	return bindataRead(
		_bindataMigrationsMigrationsgo,
		"migrations/migrations.go",
	)
}



func bindataMigrationsMigrationsgo() (*asset, error) {
	bytes, err := bindataMigrationsMigrationsgoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "migrations/migrations.go",
		size: 0,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1522913763, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"migrations/1_init.sql":    bindataMigrations1initsql,
	"migrations/migrations.go": bindataMigrationsMigrationsgo,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"migrations": {Func: nil, Children: map[string]*bintree{
		"1_init.sql": {Func: bindataMigrations1initsql, Children: map[string]*bintree{}},
		"migrations.go": {Func: bindataMigrationsMigrationsgo, Children: map[string]*bintree{}},
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
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
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