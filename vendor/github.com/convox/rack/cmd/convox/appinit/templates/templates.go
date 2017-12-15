// Code generated by go-bindata.
// sources:
// appinit/templates/Dockerfile
// appinit/templates/dockerignore
// appinit/templates/gitignore
// DO NOT EDIT!

package templates

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

var _appinitTemplatesDockerfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xc1\x4b\xc3\x30\x18\xc5\xef\xf9\x2b\xde\xa1\x0c\x05\xd7\xdc\x07\x3d\xe9\x44\x91\xad\x52\x50\xf1\x98\xa5\x9f\x6d\x68\x9b\x84\x9a\x0c\x24\xcb\xff\x2e\x69\xb6\x83\x0a\x9e\xf2\x25\xbc\xdf\x7b\xf9\xde\x7d\x53\xef\xd0\xd3\x6c\x06\xcf\x25\xb5\x62\x66\xac\x79\xd9\x43\xb6\xe0\x6e\xb2\x58\xad\xd0\x29\x07\x39\x1a\x4d\xe8\x9d\xb3\x9f\x1b\xce\x3b\xe5\x7a\x7f\x28\xa5\x99\xf8\x99\xcc\xc7\xfa\xe0\xd5\xd8\x5a\x21\x87\x75\x08\x28\x07\xa5\x5b\xc4\xc8\xb6\xfb\x57\x3c\xd4\xbb\x6d\xc5\x85\xb5\xec\xad\x6e\x9e\xee\x1e\x1b\x2c\x97\x10\xa0\x3e\x50\x92\x3e\xaa\xd9\xe8\x89\xb4\xbb\x00\x21\x60\x16\xba\x23\x14\x03\x7d\xdd\xa0\x38\x8a\xd1\x13\x36\xd5\x6f\x71\x08\x8b\x02\x31\x56\x69\xcc\xb2\x18\x13\x4f\x39\xfe\xc7\x94\xd2\xec\x4c\xd2\x4c\x56\x8d\x94\xf9\x3f\x0f\x67\x39\xbb\xad\x9f\xdf\x51\xe6\x9f\x2e\xad\x18\xef\xac\x77\x55\x71\x95\xba\xf9\x77\x69\x7e\x50\x9a\x5f\x4c\x93\xc1\x52\x27\x97\x42\xf6\x74\x8d\xd3\x09\x24\x7b\x83\x22\x1b\xb2\xef\x00\x00\x00\xff\xff\xf0\x17\xf3\xb4\x85\x01\x00\x00")

func appinitTemplatesDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_appinitTemplatesDockerfile,
		"appinit/templates/Dockerfile",
	)
}

func appinitTemplatesDockerfile() (*asset, error) {
	bytes, err := appinitTemplatesDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "appinit/templates/Dockerfile", size: 389, mode: os.FileMode(420), modTime: time.Unix(1475420665, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _appinitTemplatesDockerignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\x4b\xce\xcf\x2b\xcb\xaf\xe0\xd2\x4b\xcd\x2b\xe3\xd2\x4b\xcf\x2c\x01\x13\x99\xe9\x79\xf9\x45\xa9\x5c\x5c\xd5\xd5\x0a\x99\x69\x0a\x7a\x10\xae\x5b\x66\x4e\x6a\xb1\x42\x6d\x2d\x48\xb4\x28\x31\x2f\x3d\x55\x41\x25\x2d\x33\x27\x55\xc1\xca\x16\x5d\x45\x75\x35\x54\xaa\xb6\x56\xa1\xba\x5a\x21\x35\x2f\x05\xaa\x0d\xca\x02\x04\x00\x00\xff\xff\xc9\x75\xf6\x46\x75\x00\x00\x00")

func appinitTemplatesDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_appinitTemplatesDockerignore,
		"appinit/templates/dockerignore",
	)
}

func appinitTemplatesDockerignore() (*asset, error) {
	bytes, err := appinitTemplatesDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "appinit/templates/dockerignore", size: 117, mode: os.FileMode(420), modTime: time.Unix(1475420665, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _appinitTemplatesGitignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\x4b\xce\xcf\x2b\xcb\xaf\xe0\xd2\x4b\xcd\x2b\xe3\x02\x04\x00\x00\xff\xff\x70\x1b\xbc\x48\x0d\x00\x00\x00")

func appinitTemplatesGitignoreBytes() ([]byte, error) {
	return bindataRead(
		_appinitTemplatesGitignore,
		"appinit/templates/gitignore",
	)
}

func appinitTemplatesGitignore() (*asset, error) {
	bytes, err := appinitTemplatesGitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "appinit/templates/gitignore", size: 13, mode: os.FileMode(420), modTime: time.Unix(1475420665, 0)}
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
	"appinit/templates/Dockerfile": appinitTemplatesDockerfile,
	"appinit/templates/dockerignore": appinitTemplatesDockerignore,
	"appinit/templates/gitignore": appinitTemplatesGitignore,
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
	"appinit": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"Dockerfile": &bintree{appinitTemplatesDockerfile, map[string]*bintree{}},
			"dockerignore": &bintree{appinitTemplatesDockerignore, map[string]*bintree{}},
			"gitignore": &bintree{appinitTemplatesGitignore, map[string]*bintree{}},
		}},
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
