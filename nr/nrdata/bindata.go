// Code generated by go-bindata.
// sources:
// nr_data.json
// sum_data.json
// DO NOT EDIT!

package nrdata

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


func nr_dataJsonBytes() ([]byte, error) {
	return bindataRead(
		_nr_dataJson,
		"nr_data.json",
	)
}

func nr_dataJson() (*asset, error) {
	bytes, err := nr_dataJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nr_data.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sum_dataJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x94\xbd\x6a\x1c\x41\x10\x84\xdf\x65\xe2\x0b\xaa\xff\xbb\xee\x55\x8c\x11\xc6\x3e\x2c\x05\x96\xe0\xee\xe4\x44\xe8\xdd\xcd\x64\xb7\xd6\xc2\x4e\xb2\xc1\x32\x50\x1f\xd5\x5f\xf7\xb7\x8f\x71\xbb\xff\xb8\xde\x9f\x9e\x2f\x2f\xbf\x9f\xef\xe3\x3c\xd4\x50\x00\xc6\x69\x5c\x5e\x7f\x3d\xfe\xf6\x32\xe1\x38\x8d\xbf\x97\xeb\xed\xe5\xed\x75\x9c\x87\x8c\xd3\xb8\xbd\xff\x19\xe7\x8f\xf1\xf2\xfa\xf4\xf6\x7e\xbf\xcd\x9f\xda\xe1\x6a\x42\xa2\x3c\xbd\x52\xaa\x34\x11\xee\xf3\xf5\xcf\xb7\xeb\x65\x3e\x62\x37\x34\x72\x7c\x9e\xc6\xe5\x7a\x1d\xe7\x31\x3e\x4f\x3b\x28\x5e\xa6\x3b\x28\x5d\x69\x2b\x28\x88\x88\x2c\x57\x6a\x28\x8c\xda\x14\x83\x74\xf7\x23\x8a\x3a\x04\x25\x47\x28\x5d\xe9\x5f\x51\x5c\x8b\xb1\x82\x42\x50\x3d\x85\x65\x25\xa9\x62\x40\x27\x43\x99\x8f\x28\xdd\x9a\x28\x3f\x40\x99\x99\xb9\x83\x92\xad\xb5\x82\xd2\x91\xf4\x19\xdd\x0a\x5a\xa9\x24\x84\x99\xd8\xb4\xd2\x6a\x15\xa6\x47\x28\xd9\xda\x5f\x51\x02\x1d\x5c\x73\x45\x14\xc9\x0c\x31\x35\xc0\xaa\xcd\xbc\x55\x7d\xdb\x4a\xd2\x15\x07\x28\x81\xce\x1d\x6d\xc3\x9b\x2b\xda\xba\xab\x54\x7a\xb1\x1a\x4e\x40\xa2\x09\xaf\xed\x80\x94\x59\xd5\x7d\x34\xa0\x99\xb9\xa3\x6d\x34\x75\x45\x5b\xb5\x32\x64\x45\xa4\xc1\x0d\x64\x87\xd7\x9c\x0f\x1e\x51\xcc\x28\x6d\x47\x1b\x34\x33\x77\xb4\x4d\x65\xac\x68\xab\x44\x07\x44\x2c\x92\x95\x4c\xb6\x80\xd9\x51\x9b\x56\xdc\x7b\xae\xf8\x01\xca\xcc\xdc\xd1\x36\x93\xbd\xa2\xad\x79\x23\x2a\xd3\xe6\x47\x58\x9d\x41\x6b\xba\x6f\x5b\x01\x44\x79\x88\x92\xf3\xfe\x7c\x41\x29\x81\x2c\x69\x0b\x8a\xa9\x2a\x55\x54\xa6\x77\x05\x9d\x92\x8a\x3e\xa2\x44\x87\x7b\xf5\x01\x4a\x09\x74\x47\xdb\x0a\xc4\x8a\xb6\xe9\x31\x8f\x17\xcb\x3b\xb2\x44\x3c\x3a\x54\xb0\xb9\x70\xea\x41\x06\x8f\x76\x79\x46\xee\x58\x5b\x44\xaf\x58\xeb\x4e\x86\x05\xa0\x89\x04\x3c\x12\xec\x94\x6d\x27\xda\x6d\x5a\x47\xd2\xce\xc8\x1d\x69\xdb\x44\x56\xa4\x0d\x51\x7a\xaa\xf7\x34\x35\x0a\x51\xe5\xec\xff\x48\xa4\xdd\x68\x9b\xe9\x7c\xff\x17\x00\x00\xff\xff\xce\xea\x4c\x5d\x0d\x07\x00\x00")

func sum_dataJsonBytes() ([]byte, error) {
	return bindataRead(
		_sum_dataJson,
		"sum_data.json",
	)
}

func sum_dataJson() (*asset, error) {
	bytes, err := sum_dataJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sum_data.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"nr_data.json":  nr_dataJson,
	"sum_data.json": sum_dataJson,
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
	"nr_data.json":  &bintree{nr_dataJson, map[string]*bintree{}},
	"sum_data.json": &bintree{sum_dataJson, map[string]*bintree{}},
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
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