// Code generated by go-bindata. DO NOT EDIT.
// sources:
// migrations/1_add_organization_table.sql
// migrations/2_add_user_table.sql
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

var _bindataMigrations1addorganizationtablesql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x41\x4b\xc3\x40\x10\x85\xef\xf3\x2b\xde\xad\x0d\x9a\x8b\xb4\xbd" +
	"\xf4\x14\xdb\x80\x62\x6c\xeb\x9a\x08\x3d\xae\xbb\x43\x3b\xb8\xd9\x5d\xb2\x91\x8a\xbf\x5e\x52\xb1\x54\x50\xe8\x75" +
	"\xbe\x37\xcc\x7c\x2f\xcf\x71\xd5\xca\xae\xd3\x3d\xa3\x89\x94\xe7\x78\x7e\xaa\x20\x1e\x89\x4d\x2f\xc1\x63\xd4\xc4" +
	"\x11\x24\x81\x3f\xd8\xbc\xf7\x6c\x71\xd8\xb3\x47\xbf\x97\x84\xef\xbd\x21\x24\x09\x3a\x46\x27\x6c\x69\xa1\xca\xa2" +
	"\x2e\x51\x17\xb7\x55\x89\xd0\xed\xb4\x97\xcf\x63\x88\xc6\x04\x88\xc5\x4b\xa1\x16\x77\x85\x1a\xdf\x4c\xa7\x19\x36" +
	"\xea\xfe\xb1\x50\x5b\x3c\x94\xdb\x6b\x02\xbc\x6e\xf9\x14\x98\x4d\x32\xac\xd6\x35\x56\x4d\x55\x0d\x90\x5b\x2d\xee" +
	"\x5f\x6a\x39\x99\x4e\xe2\xf1\x9d\xf3\x13\x03\x73\xc1\xe8\x5f\x60\x36\xc9\x28\x9b\x13\xd1\xb9\xfe\x32\x1c\xfc\x4f" +
	"\x01\x27\xfb\x61\x78\x91\x7f\x17\x9c\x63\x8b\x57\x6d\xde\x68\xa9\xd6\x9b\x3f\x1a\x98\xd3\x57\x00\x00\x00\xff\xff" +
	"\xf1\x50\x03\x79\x6e\x01\x00\x00")

func bindataMigrations1addorganizationtablesqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataMigrations1addorganizationtablesql,
		"migrations/1_add_organization_table.sql",
	)
}



func bindataMigrations1addorganizationtablesql() (*asset, error) {
	bytes, err := bindataMigrations1addorganizationtablesqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "migrations/1_add_organization_table.sql",
		size: 366,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1535920294, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataMigrations2addusertablesql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x41\x4b\xc3\x40\x10\x85\xef\xf3\x2b\xde\xad\x2d\x1a\x10\xb1\x5e" +
	"\x7a\x8a\x6d\x40\x31\xb6\x75\x4d\x84\x1e\xb7\xc9\xd0\x0e\x6e\x76\xc3\xee\xd6\xfa\xf3\x65\x29\x96\x54\x10\xbc\xbe" +
	"\xef\xbd\x61\xf8\xb2\x0c\x57\x9d\xec\xbc\x8e\x8c\xba\xa7\x2c\xc3\xdb\x6b\x09\xb1\x08\xdc\x44\x71\x16\xa3\xba\x1f" +
	"\x41\x02\xf8\x8b\x9b\x43\xe4\x16\xc7\x3d\x5b\xc4\xbd\x04\x9c\x76\xa9\x24\x01\xba\xef\x8d\x70\x4b\x73\x55\xe4\x55" +
	"\x81\x2a\x7f\x28\x0b\x1c\x02\x7b\x1a\x13\x20\x2d\xde\x73\x35\x7f\xcc\xd5\xf8\x76\x7a\x33\xc1\x5a\x3d\xbd\xe4\x6a" +
	"\x83\xe7\x62\x73\x4d\x80\x71\x3b\xb1\xe7\xc6\xfd\xdd\x04\xcb\x55\x85\x65\x5d\x96\x89\x5a\xdd\xf1\x9f\x90\x3b\x2d" +
	"\x66\x70\x7c\x7a\x89\x8d\x6b\x4e\x2f\x0e\xf6\x29\xdf\x8a\xbb\x18\xa5\x4c\x7f\xea\xa8\x7d\xed\xcd\x2f\x32\x99\x11" +
	"\xd1\x50\xd4\xc2\x1d\xed\x8f\xaa\xb3\xa7\x14\xfe\xcb\x94\x77\xc6\x70\x8b\xad\x6e\x3e\x68\xa1\x56\xeb\x81\xab\x19" +
	"\x7d\x07\x00\x00\xff\xff\xe4\xbf\xeb\xf5\x90\x01\x00\x00")

func bindataMigrations2addusertablesqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataMigrations2addusertablesql,
		"migrations/2_add_user_table.sql",
	)
}



func bindataMigrations2addusertablesql() (*asset, error) {
	bytes, err := bindataMigrations2addusertablesqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "migrations/2_add_user_table.sql",
		size: 400,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1535920526, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataMigrationsMigrationsgo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\x5b\x8f\xdb\xc8\xd1\x7d\x16\x7f\x45\x7f\x02\x76\x21\x7d\x99\x68" +
	"\x78\xbf\x08\xf0\xcb\xda\x0e\xe0\x07\x7b\x81\xc4\x79\x4a\x05\x46\x93\xdd\xad\x25\x56\x12\xc7\x24\xb5\xa9\x19\xc3" +
	"\xff\x3d\x38\x5d\x2d\x8d\x76\x62\x8f\xd7\x40\x80\x18\xd8\x59\x91\xec\xaa\x3a\x75\xe9\x73\xea\xf6\x56\xbd\x1c\x8c" +
	"\x55\x3b\x7b\xb4\xa3\x9e\xad\x51\xed\xbd\xda\x0d\x7f\x6e\xfb\xa3\xd1\xb3\xde\xa8\x57\x3f\xab\x77\x3f\xbf\x57\xaf" +
	"\x5f\xbd\x79\xbf\x89\x6e\x6f\xd5\x34\x9c\xc6\xce\x4e\x5b\xfc\x3e\xf4\xbb\x51\xcf\xfd\x70\x9c\x6e\x93\x0f\xda\x98" +
	"\x0f\xc3\xb8\xd3\xc7\xfe\xc1\xbf\xfb\x30\xeb\x76\x6f\x37\xd3\xc7\xfd\x93\xa3\xa9\x3f\x7a\x9a\xec\xf8\xd5\x23\x8f" +
	"\x3f\x37\xbb\x21\x8a\xee\x74\xf7\xab\xde\x59\x35\xda\xbb\x61\xea\xe7\x61\xbc\x8f\xa2\xa8\x3f\xdc\x0d\xe3\xac\x56" +
	"\xd1\x62\xd9\xde\xcf\x76\x5a\x46\x8b\x65\x37\x1c\xee\x46\x3b\x4d\xb7\xbb\x87\xfe\x0e\x2f\xdc\x61\xc6\xff\xfa\x41" +
	"\xfe\xde\xf6\xc3\x69\xee\xf7\x78\x18\xbc\xc1\x9d\x9e\x7f\xb9\x75\xfd\xde\xe2\x07\x5e\x4c\xf3\xd8\x1f\x77\xfe\xdb" +
	"\xdc\x1f\xec\x32\x5a\x47\x91\x3b\x1d\x3b\x15\x0a\xf2\x57\xab\xcd\x0a\x3f\xd4\x3f\xfe\x89\xb0\x37\xea\xa8\x0f\x56" +
	"\x89\xd9\x5a\xad\xce\x6f\xed\x38\x0e\xe3\x5a\x7d\x8a\x16\xbb\x07\xff\xa4\xb6\x2f\x14\x50\x6d\xde\xd9\x7f\xc1\x89" +
	"\x1d\x57\x1e\x36\x9e\x7f\x3a\x39\x67\x47\xef\x76\xbd\x8e\x16\xbd\xf3\x06\xff\xf7\x42\x1d\xfb\x3d\x5c\x2c\x46\x3b" +
	"\x9f\xc6\x23\x1e\x6f\x94\x3b\xcc\x9b\xd7\xf0\xee\x56\x4b\x38\x52\x3f\x7c\xdc\xaa\x1f\x7e\x5b\x0a\x12\x1f\x6b\x1d" +
	"\x2d\x3e\x47\xd1\xe2\x37\x3d\xaa\xf6\xe4\x94\xc4\x91\x20\xd1\xe2\x83\xc0\x79\xa1\xfa\x61\xf3\x72\xb8\xbb\x5f\xfd" +
	"\xd8\x9e\xdc\x8d\xda\x3d\xac\xa3\x45\xb7\x7f\x7d\x46\xba\x79\xb9\x1f\x26\xbb\x5a\x47\xff\x2d\x3c\x70\x23\xfe\xbf" +
	"\xe2\xc8\x8e\xa3\xe0\x0e\x2f\xdb\x93\xdb\xfc\x04\xe8\xab\xf5\x0d\x4e\x44\x9f\xa3\x28\x9a\xef\xef\xac\xd2\xd3\x64" +
	"\x67\xd4\xfc\xd4\xcd\x70\xe3\x13\x0c\x0d\x89\x16\xfd\xd1\x0d\x4a\xa1\xa9\x6f\x8e\x6e\x78\xcd\xb0\xf3\x66\x8f\xaf" +
	"\x54\x7f\x9c\xed\xe8\x74\x67\x61\x3e\x4c\x9b\xbf\x84\x4f\xd1\xe2\xed\xab\xe2\xe5\x2f\xb6\xfb\x75\x3a\x1d\x56\xeb" +
	"\xd0\xd7\x8b\x87\x30\x04\xe7\xd3\x57\x10\xfc\x14\x84\x7f\xc1\x68\x31\xf5\x0f\x97\x77\xfd\x71\x2e\xf3\x68\x71\xc0" +
	"\x5d\x0b\xff\x42\xd8\xb7\x83\xb1\xfe\xc3\xfb\x3e\xb8\xc0\xe0\x6d\xf0\x14\x2d\x0e\xa6\xe8\x02\x9a\x2b\x2c\x7e\x20" +
	"\x57\xae\x7f\x8a\x67\xad\xde\xe9\x83\xbd\xc0\x06\xae\x50\x4b\xd7\x6f\x80\x30\xfa\xfc\x8c\xed\xdf\xfa\x07\xd8\x7a" +
	"\xa4\xbf\x37\x45\x22\xcf\x9a\x22\x87\xd5\xfa\x3a\xa3\xdf\x3b\x40\xda\xdf\x72\x80\x84\x57\xeb\xc7\xe4\xff\xc3\x83" +
	"\xaf\xc8\xb3\x4e\xbe\xd0\xba\x27\x5e\x1e\xcb\xf9\xac\xa7\x37\xd3\xab\x7e\x5c\xad\x55\x3b\x0c\xfb\x6b\x0f\x7a\x3f" +
	"\x7d\xa3\x86\xf7\x93\x94\x50\xa6\xeb\xd3\xe7\x2b\xeb\x30\xc2\xb8\x95\x1f\x82\xe1\xdb\x0b\xcf\x25\xda\x98\x6b\xf6" +
	"\xf4\xcc\x38\x7d\xdc\xab\x17\x61\xae\xc1\x73\xc4\x89\x23\xae\x5b\xe2\xb8\x26\x8e\xe3\x2f\xff\xe7\x70\xa6\x23\x6e" +
	"\x62\xe2\x3c\x21\xce\x5b\xe2\x2e\x23\xce\x63\xe2\x24\x26\xae\x0b\x62\xeb\x88\x5d\x46\x9c\xb6\xc4\xc6\x12\x6b\x43" +
	"\x1c\x1b\xe2\x46\x8b\xff\x36\x27\x6e\xcd\x52\xfd\xc9\x47\x75\x39\x71\x92\x13\x9b\x96\xb8\x8e\x89\xcb\x94\xb8\xec" +
	"\x88\x6d\x2b\x16\x40\x93\x19\x62\x6d\x89\xdb\x96\x38\xcf\x88\x33\x78\xa9\x89\x4d\x43\x5c\x18\xe2\x36\x25\x6e\x12" +
	"\xe2\x5a\x13\xb7\x8e\xb8\xb0\xc4\x45\x4a\xdc\x26\xc4\x45\x4e\x5c\xc4\xc4\xb6\x26\xae\x8a\x73\xd4\xd6\x12\x67\x15" +
	"\x71\xd7\x11\x57\x1d\x71\xea\x88\x3b\x47\x5c\x25\xc4\xa6\x20\xee\xb4\x44\x34\x59\x88\x9e\x11\xd7\x0d\x71\x93\x13" +
	"\xdb\x8a\xb8\x82\x37\xe4\xa6\x89\x53\xe4\x6e\xe5\x7b\x6e\x82\xaf\x84\xb8\xcc\x88\x4d\x4e\xdc\xe5\xe7\xa8\x49\x42" +
	"\x9c\xe6\xc4\x75\x42\x9c\x39\x62\x53\x13\xb7\x1d\xb1\xab\x24\x67\x1f\x1d\xef\x50\xd1\x4a\x72\x69\x2a\xe2\x3a\x97" +
	"\xaa\xb6\xf0\x1e\x7c\xc4\x0d\x71\xa6\x89\xf3\x92\x38\x0d\xf6\x65\x43\xac\x93\x80\x3e\x3d\x47\x4d\x51\x8d\x84\x38" +
	"\x81\xc7\x8a\xb8\x28\x04\xab\x41\x55\x8c\xf4\x03\x51\x90\x3f\x30\xd7\x35\x71\x57\x12\xc7\x79\xf8\x5d\x48\x9f\xe1" +
	"\x39\x29\x89\xab\x4a\xfa\x9c\x68\x62\xe3\x88\xf3\x8e\x58\x57\xc4\x09\x10\x95\xe7\xa8\x56\x13\x3b\x1b\x7a\x10\x13" +
	"\x17\x2d\x71\xd6\x49\x05\xd1\xeb\x12\x93\x96\x4a\xfe\xa5\x25\x76\x8d\xcc\x41\x53\x4b\x15\xb3\x94\x58\x77\xc4\xa6" +
	"\x24\xce\x0a\xe2\xa2\x94\xf7\x40\xef\xa7\x29\xf8\x4c\x0d\xb1\xb5\xe7\xa8\x05\x72\xd0\xc4\x19\x7a\x85\x7e\x58\x62" +
	"\x9b\x12\xbb\x84\xb8\x31\x32\x97\x49\x46\x1c\x67\xc4\x55\x26\x7d\xc2\x64\x78\xbb\x18\xf8\x89\xbb\x86\x38\xad\x89" +
	"\x9b\x56\xce\x9a\x84\xb8\x6d\x24\x1b\xa0\x4a\xd0\xb1\x8e\x38\x77\xe7\xa8\x71\x22\x3d\x70\xc8\x2b\x91\xe9\xc0\x3c" +
	"\x56\x4e\xaa\xde\x74\xa1\xb2\x2d\x71\x81\x5e\x19\xb9\x19\x65\x4d\xac\x1b\xc9\x11\xd1\x30\x13\xa8\x2a\x6a\x80\xc9" +
	"\xc3\xd9\xa7\xb7\xcf\x5d\xa2\xba\x50\x59\x9f\x4b\x23\x55\x04\x12\x39\xbd\x7c\xb2\x5c\x7c\x9b\x0e\x82\x1e\x7e\x69" +
	"\xcf\x38\xab\xe6\xd5\x9e\x12\x2d\x16\x7f\x9c\x69\x6e\xa2\xc5\x62\xf9\x07\x37\xba\xe5\x4d\xb4\x58\x7b\x3d\xfe\x4e" +
	"\xfc\x80\xfe\xff\x5e\xbe\xaf\xa1\x7b\xfd\xbe\x6c\x49\xdf\x5b\x8c\x6f\x2d\x28\x97\xbd\xc2\x2f\x06\x8f\x01\xce\xac" +
	"\x8d\xf3\x10\xc7\xad\xfa\x9e\xf4\xbd\xba\x6f\x55\x56\x96\x78\xb8\x12\x96\xad\x5a\xfa\xef\xd0\xbc\xed\xb5\x24\xae" +
	"\xf2\x34\x5e\x87\x2f\xd0\xb2\xad\x68\xdd\xdf\x8f\x3d\xaf\x92\x22\x2b\x9a\x34\x4e\x9b\xfc\x46\xf9\x43\xc0\xab\x01" +
	"\xf6\x47\x5f\xad\x4f\xbe\x44\x5b\x15\x2a\x85\x4c\xb6\xfe\xef\xd5\xbe\xa4\x6f\x9e\x97\x98\x54\x1b\x83\xad\xfb\x7f" +
	"\x20\x2d\x20\x00\x5c\x1a\x9c\xf1\x64\x73\x21\x83\x4a\x8b\x24\xe0\xba\xc1\x4b\x86\x8b\x0c\x02\x2b\x84\x4a\x40\xab" +
	"\x20\x6e\x90\x62\x17\x08\x31\xb6\x72\x91\xaa\x52\xa2\x5b\x2b\xd7\xd3\x69\x89\x5c\x16\xc4\x29\xc8\xa5\x14\x69\xf1" +
	"\x51\x3b\x08\xca\x85\xf8\x02\x51\x83\x08\x5c\x2d\xe2\x14\x77\x72\x99\x41\x41\xb6\x13\x8b\xda\x4a\xbe\xad\x16\xfa" +
	"\x4c\x3b\x89\x78\x21\xc7\x46\xf2\x41\xcd\x0c\xe8\x26\x17\x81\x00\x05\x43\x8e\x60\x97\x5c\xc8\x00\x15\x03\x9d\x22" +
	"\x22\x2a\xdd\x04\xa9\xb4\xb9\x58\x74\x95\x08\x19\x28\xb3\xcb\x05\x21\x08\x1e\xd4\x84\x5c\x41\x43\x5e\x56\x12\xf1" +
	"\x8c\x2c\x6a\x43\x5c\xc5\xd2\x09\x90\x25\xa8\x17\xfe\x8a\x8b\x8c\x42\xc8\x52\x2d\x44\x07\xca\x8c\x5b\xa1\x47\x20" +
	"\xa9\xda\xd0\x97\x4c\x24\x32\x0d\xa4\x07\x4f\x20\x74\xd8\x40\x7e\x21\xb7\x40\x0d\x14\xe8\x58\x16\x68\xb9\xd0\x82" +
	"\x18\x48\x11\xb5\xd4\x97\xa8\x99\x88\x30\x56\x05\x78\xf3\xbd\x8c\xa5\x3a\x7e\x35\x48\xe4\x3b\x44\xcc\x21\x6a\xc8" +
	"\xb5\x6b\x83\xf8\x15\xd2\x47\xac\x0e\xe8\x23\xc4\x10\xd1\x70\xce\x0b\x85\x93\x29\xcc\xbc\xbc\x9c\xa3\x96\xa5\x54" +
	"\x03\xeb\x82\x9f\xac\x46\x2a\x84\x5e\x41\x66\x20\xfb\x98\x1e\x57\xca\x84\x40\x1a\xfd\x42\x82\x3a\x80\xe8\x0b\x91" +
	"\x4a\xe4\x0d\x71\xd4\xc8\xd9\x88\x04\x77\x61\x6d\x80\xbc\x40\xb6\x92\xe4\x1c\xd5\x04\x92\xf7\x8b\x04\xc4\x27\x58" +
	"\xd4\x4e\x96\x0f\xac\x0a\x5e\x7c\x73\x11\x28\xe4\x08\x89\xad\x42\xfe\x40\x8c\x79\xc0\x4d\xc1\x5c\x67\x67\xd9\x49" +
	"\x44\x52\xb1\x66\xa1\x8b\x1a\x28\x9b\xcb\xcd\x41\xc5\xaa\xa7\xc2\x23\x7d\xc0\x52\x02\x2b\x57\x48\x95\xbe\x2e\x39" +
	"\xff\x0e\x00\x00\xff\xff\x5a\xb3\xa8\xc1\x00\x10\x00\x00")

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
		size: 8192,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1535925552, 0),
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
	"migrations/1_add_organization_table.sql": bindataMigrations1addorganizationtablesql,
	"migrations/2_add_user_table.sql":         bindataMigrations2addusertablesql,
	"migrations/migrations.go":                bindataMigrationsMigrationsgo,
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
		"1_add_organization_table.sql": {Func: bindataMigrations1addorganizationtablesql, Children: map[string]*bintree{}},
		"2_add_user_table.sql": {Func: bindataMigrations2addusertablesql, Children: map[string]*bintree{}},
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
