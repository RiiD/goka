// Code generated by go-bindata. DO NOT EDIT.
// sources:
// web/templates/common/base.go.html (141B)
// web/templates/common/head.go.html (1.28kB)
// web/templates/common/menu.go.html (369B)
// web/templates/monitor/details.go.html (8.148kB)
// web/templates/monitor/index.go.html (2.283kB)
// web/templates/monitor/menu.go.html (299B)
// web/templates/query/index.go.html (2.479kB)
// web/templates/index/index.go.html (477B)

package templates

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _webTemplatesCommonBaseGoHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\x4a\x2c\x4e\x55\xaa\xad\xe5\xb2\x51\x74\xf1\x77\x0e\x89\x0c\x70\x55\xc8\x28\xc9\xcd\xb1\xe3\xaa\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\xca\x48\x4d\x4c\x51\x52\xd0\x03\xa9\x82\x48\xda\x24\xe5\xa7\x54\xa2\xaa\xc9\x4d\xcd\x2b\x85\xa8\x41\x16\x4d\xce\xcf\x2b\x49\xcd\x2b\x81\x6a\xd6\x87\x68\xb3\xd1\x87\x59\x91\x9a\x97\x52\x5b\xcb\x05\x08\x00\x00\xff\xff\xfd\x8f\xc0\x67\x8d\x00\x00\x00")

func webTemplatesCommonBaseGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webTemplatesCommonBaseGoHtml,
		"web/templates/common/base.go.html",
	)
}

func webTemplatesCommonBaseGoHtml() (*asset, error) {
	bytes, err := webTemplatesCommonBaseGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/common/base.go.html", size: 141, mode: os.FileMode(0644), modTime: time.Unix(1588163175, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x75, 0x44, 0xa2, 0x23, 0x83, 0xc9, 0x36, 0xe6, 0xa0, 0xdb, 0x49, 0xde, 0xb1, 0x3c, 0xfd, 0x55, 0x6, 0x81, 0x42, 0xa5, 0xff, 0x8d, 0x4c, 0xcf, 0x62, 0xf, 0x43, 0x37, 0x80, 0x88, 0x58, 0xbb}}
	return a, nil
}

var _webTemplatesCommonHeadGoHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x4d\x73\xf2\x36\x18\xbc\xe7\x57\x78\x74\xe1\x90\x5a\xc2\x98\x16\x92\xc1\xe9\xa4\xa4\xe5\x23\x5f\x04\x12\x08\xbd\x74\x84\xf4\xd8\x96\x23\x4b\x8e\x25\x03\x2e\xe5\xbf\x77\x1c\x4a\xc2\xa4\x79\x3f\x72\x78\x6f\x7e\xd6\xab\xf5\xee\xfa\x19\x6d\x36\x1c\x42\xa1\xc0\x41\x31\x50\x8e\xb6\xdb\xa3\x4e\xf5\x70\x76\xe4\x38\x8e\xd3\x49\xc1\x52\x87\xc5\x34\x37\x60\x03\x54\xd8\xd0\x6d\xa3\xc3\x57\xb1\xb5\x99\x0b\xcf\x85\x58\x06\x68\xed\x16\xd4\x65\x3a\xcd\xa8\x15\x0b\x09\xc8\x61\x5a\x59\x50\x36\x40\x02\x02\xe0\x11\xec\x4f\x5a\x61\x25\x9c\x6d\x36\x38\xa3\x11\xfc\xf5\x32\x6d\xb7\x1d\xb2\x83\x0f\xc4\x15\x4d\x21\x40\x1c\x0c\xcb\x45\x66\x85\x56\x07\x92\xe8\xff\xc4\xa5\x80\x55\xa6\x73\x7b\xc0\x5a\x09\x6e\xe3\x80\xc3\x52\x30\x70\x5f\x86\x9f\x1c\xa1\x84\x15\x54\xba\x86\x51\x09\x81\xb7\x17\x92\x42\x3d\x39\x71\x0e\x61\x50\xab\x42\x99\x53\x42\x42\xad\xac\xc1\x91\xd6\x91\x04\x9a\x09\x83\x99\x4e\x09\x33\xe6\xd7\x90\xa6\x42\x96\xc1\x6d\x06\xea\x78\x42\x95\x39\xee\x6a\xc5\x41\x19\xe0\xa7\x7e\xbd\xfe\xcf\x2b\x5e\x73\x72\x90\x41\xcd\xd8\x52\x82\x89\x01\x6c\xcd\xb1\x65\x06\x41\xcd\xc2\xda\x56\x4a\xb5\xc3\x8f\x57\x5c\xf4\xc6\x45\x3b\x37\x68\xef\x26\xa5\x6b\xc6\x15\x5e\x68\x6d\x8d\xcd\x69\x56\x0d\x95\xa1\x57\x80\xf8\xd8\xc7\xad\x4a\xf6\x0d\xc3\xa9\x50\x98\x19\x83\x1c\xa1\x2c\x44\xb9\xb0\x65\x80\x4c\x4c\xfd\x76\xd3\xfd\x6d\x3a\x17\x62\x32\xf8\x03\x2e\x3d\xde\x4b\x87\xe3\xf3\xa7\x92\x15\xfd\xf3\xfe\x38\xf2\x1b\xb7\xe9\x03\x5b\xad\x5a\x5a\xf9\xe3\x39\x8f\x9a\x53\x7a\x3c\x4a\x27\xf7\xe6\x6f\x72\xf9\x4b\x7b\xb9\xe0\xbf\x27\x71\xb3\x40\x0e\xcb\xb5\x31\x3a\x17\x91\x50\x01\xa2\x4a\xab\x32\xd5\x85\x41\x3f\x3a\x94\x6b\x63\x48\xe1\x6b\xd1\xf2\x7e\xa9\x6f\x3c\x31\x36\xd3\xc7\x69\x53\x5d\xd4\x87\x85\x95\xaa\x47\x8d\xec\x0e\x8b\x6e\xab\x58\x25\xbc\x98\x9d\x4c\xa6\xf9\xd5\x72\x3c\xd7\x7a\x94\x35\x16\xb3\x79\x94\x46\xc3\xbb\xc1\xe3\x4a\x92\x49\xf6\xad\x68\xbb\x8d\x74\x4c\xce\x02\x44\x08\x4d\xe8\xfa\xfd\x9a\x54\x18\x91\x62\x61\x48\xf2\x5c\x40\x5e\x12\x0f\x7b\x1e\xae\xff\x37\xbd\x78\x4f\x0c\x3a\xeb\x90\x9d\xd4\x07\xba\x9f\xad\x28\x79\xff\xdb\x93\x0f\xab\xb9\x67\x3f\x0f\xee\xc4\xa2\xde\x68\x3d\x2f\xcb\x64\x72\x1d\xf6\x93\xdb\x6b\x7a\xf5\x14\x16\xb3\xe9\xfa\xcf\xf5\xc3\x48\x75\x87\xe7\x2d\xd9\x48\xbb\xb3\x9b\x41\xd6\x3b\x49\x7b\xdd\x8b\xf6\xaa\x77\x33\x60\xa3\x8b\xd6\xfd\x9a\x7e\xb9\x9a\xef\xc8\xc2\xb8\x4a\x0c\x66\x52\x17\x3c\x94\x34\x87\x77\x55\x49\xcd\xa9\x89\x71\x62\x48\x13\x7b\x2d\xec\xed\x81\x4f\xb4\xc5\xfd\xc4\x60\x9d\x47\x84\xfb\x78\xd9\xfc\xe0\x64\x87\xec\xae\xb7\xcd\x06\x14\xdf\x6e\x8f\xfe\x0d\x00\x00\xff\xff\x25\x6d\x7e\xc6\x00\x05\x00\x00")

func webTemplatesCommonHeadGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webTemplatesCommonHeadGoHtml,
		"web/templates/common/head.go.html",
	)
}

func webTemplatesCommonHeadGoHtml() (*asset, error) {
	bytes, err := webTemplatesCommonHeadGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/common/head.go.html", size: 1280, mode: os.FileMode(0644), modTime: time.Unix(1588163175, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc8, 0xd2, 0x19, 0x80, 0x2f, 0xb7, 0x33, 0x91, 0x41, 0x9, 0x21, 0x7f, 0x21, 0xc6, 0xb8, 0x4a, 0xe2, 0x3a, 0xe6, 0x79, 0x81, 0x2e, 0x82, 0xa2, 0x79, 0x89, 0xe2, 0xb9, 0x26, 0xe2, 0xac, 0xe8}}
	return a, nil
}

var _webTemplatesCommonMenuGoHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\x41\x4e\xc4\x30\x0c\x45\xf7\x3d\x85\x65\x36\xb0\x60\x7a\x81\xb4\x0b\x38\x08\x72\x13\x57\x8d\x6a\x1c\x94\x36\x91\x50\x94\xbb\x23\x98\x36\xc0\xcc\x2a\xb1\xfe\x8b\xde\x8f\x4b\x71\x3c\x7b\x65\xc0\x77\xd6\x84\xb5\x76\x46\x29\x83\x15\xda\xb6\x01\x95\xf2\x44\x11\xae\xc7\xb3\xe3\x99\x92\xec\x38\x76\x00\xc6\xf9\x46\xd9\xa0\x3b\x79\xe5\xf8\x93\xfc\xcf\x8e\xa7\x0b\x93\x6b\x39\x80\xa1\x9b\x7c\x8a\xa4\x0e\x61\x89\x3c\x0f\xf8\x80\x10\xf4\x55\xbc\x5d\x07\x5c\xfc\xb6\x87\xf8\x79\x99\xc8\xae\x8f\x4f\x38\xbe\x90\x5d\x4d\x4f\x87\xa8\x77\x3e\xdf\x3b\x6d\x10\xa1\x8f\x8d\xcf\xde\xe7\xfc\xab\x4f\xf2\xc7\x7f\x62\x4a\xb9\x11\xa5\x4c\x12\xec\x7a\xdd\xca\xdb\xf7\x07\x59\x77\x84\x4b\xad\x0d\x60\x75\x6d\x32\x7d\x92\x9b\x4a\xc7\xc5\xf4\x4a\x79\xec\x4a\x01\x56\x07\xb5\x76\x5f\x01\x00\x00\xff\xff\x42\x5b\x78\xaf\x71\x01\x00\x00")

func webTemplatesCommonMenuGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webTemplatesCommonMenuGoHtml,
		"web/templates/common/menu.go.html",
	)
}

func webTemplatesCommonMenuGoHtml() (*asset, error) {
	bytes, err := webTemplatesCommonMenuGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/common/menu.go.html", size: 369, mode: os.FileMode(0644), modTime: time.Unix(1588163175, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xad, 0x11, 0x1a, 0x8d, 0x3a, 0x6, 0x26, 0x97, 0x5e, 0x76, 0xcd, 0xf0, 0x55, 0x23, 0x55, 0x6e, 0xbe, 0x21, 0xfa, 0x77, 0x4d, 0xfd, 0x4, 0xf5, 0x34, 0x95, 0xe2, 0x44, 0xb9, 0xd3, 0x4b, 0x9d}}
	return a, nil
}

var _webTemplatesMonitorDetailsGoHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x59\xdf\x6f\xdb\x38\xf2\x7f\xf7\x5f\x31\x5f\x7d\x1f\x2a\xa1\xb6\x6c\x27\x6f\x75\x6c\xe0\xae\xc5\x5d\xbb\xd7\x76\x8b\x6b\xb1\x7d\xc8\x05\x05\x2d\x8e\x6d\xa6\x12\xa9\x23\x69\x3b\x5e\x43\xff\xfb\x81\xa4\x24\x93\x8a\x1d\x67\x7b\x09\xb0\xc5\xe5\x21\xb0\x39\x3f\x39\x9c\xf9\xcc\x90\xde\xef\x29\x2e\x18\x47\x88\x32\xc1\x35\x72\x1d\x55\x55\xef\x8a\xb2\x0d\x64\x39\x51\x6a\x1a\x49\xb1\x8d\x66\x3d\x00\x7f\xcd\xb0\x12\xc6\x51\x5a\x4a\x97\x96\x0f\x0a\x3a\x18\x5f\x34\xb4\xd5\x78\xb6\xdf\xa7\x9a\xe9\x1c\xab\xea\x6a\xb8\x1a\xdf\x97\x29\x09\xc7\x1c\xec\xff\x01\xc5\x05\x59\xe7\xba\x96\x3e\xcd\xb7\x42\x42\x19\x5f\xb6\x7c\xc6\xd2\xe5\xec\x0b\x99\xe7\x08\x4a\x13\xcd\x94\x66\x99\xba\x1a\xae\x2e\x5b\x4d\x43\xca\x36\x27\xd5\x0e\xe6\x82\xee\x7c\x6d\xda\xaa\xaa\x59\xdc\x17\xfb\x7f\xa0\xb4\x64\x25\x52\x8f\xd7\x70\x1b\x7f\xfc\x15\xb3\x26\xc3\x05\xcb\x06\x36\x12\xc6\xa8\xd4\x4c\x33\xc1\x81\xaf\x8b\x39\xca\x08\xa4\xd8\xaa\x92\xf0\x69\x74\x11\xcd\x3e\x35\xd4\xab\xa1\x5e\x3d\xa0\xe5\xb3\x26\x7a\xad\x40\x2c\x40\xaf\x10\x4a\x29\x32\x54\x4a\x74\x94\x19\x26\x3c\xa3\xe8\xa3\x75\xc2\x28\x2a\x50\x29\xb2\x44\x05\x39\x59\x2e\x19\x5f\xc2\x1c\x57\x8c\x53\x78\xfb\xf5\x43\xa8\xf6\xd7\xc5\x42\xa1\x1e\xbc\x27\xcb\x33\xba\xdf\xb2\xe5\x0a\xb6\x44\xa3\x84\x82\xc8\xef\x10\x0b\x2b\xd9\x78\xcd\xf1\x4e\x37\x56\x61\x8e\xc6\x64\x26\xb8\x5a\x17\x48\x93\xd0\xe2\xdb\xaf\x1f\x4e\x99\xf2\xd9\xbe\x4a\xa6\x71\xf0\xcf\x07\x36\x9d\x89\xdc\x71\x5f\x46\xb3\x77\xbc\x5c\xeb\xb3\x8c\x66\xbf\x6b\x7d\x94\xf3\x6a\xd8\x3d\xe8\xe3\x27\x3f\x93\xa7\x3d\x9a\xcd\x77\x1a\xd5\x29\x22\xc5\x9c\xec\x4e\x11\x7f\x40\x6b\xd7\x63\xc3\xd1\x49\xdf\x2b\x6d\xea\x01\x18\xf5\x32\xf5\x37\x86\xdb\xa8\x23\x67\xb8\xbc\xa2\x19\xda\x12\x39\x52\x71\xf5\xc7\xde\x7e\xcf\x16\x80\xff\x86\x54\x22\xa7\x28\xbf\xec\x4a\x84\xe8\x90\xb7\x55\xf5\x8c\xe0\xf0\x5a\x0c\x7e\x11\x8c\x23\x05\x0b\x13\x3f\x13\x38\x3c\x1a\x10\xbe\x88\x92\x65\x75\x15\xdd\xda\xcd\x46\x0e\x14\x7f\x10\x48\x9e\x10\x3c\xfe\x5c\x80\x51\x9b\x32\x28\x11\x38\xbe\x95\x4c\x6b\xe4\xa0\x05\xe4\x22\x23\xb9\x6d\x26\x08\x25\x4a\x50\x98\x09\x4e\x1f\xc6\x97\x3f\x56\x5a\xe6\x84\x9e\xa0\xaa\xf6\x7b\xe4\xb4\xaa\x7a\x6e\x4d\x65\x92\x95\x1a\xf4\xae\xc4\x69\xa4\xf1\x4e\x0f\x6f\xc9\x86\xb8\xd5\x68\xd6\xab\xb5\x6c\x88\x84\x9c\x28\xdd\x26\x16\x4c\x81\x5e\xa6\x76\xcf\x71\x32\xe9\x70\x99\xba\x31\x99\xa0\x4e\x72\x29\x9b\x40\x1f\x48\x09\x53\xd8\x8f\x5e\x41\x24\x31\x13\x1b\x94\xa6\x0c\xfb\x30\x7e\x15\x95\x12\x4b\x52\x7f\xbd\x30\xf4\x35\xe7\xe6\x5b\x35\xf1\x5d\x72\xb0\xf0\x06\x35\x61\xb9\x31\xb6\x58\xf3\xcc\x78\x17\xb7\x25\xa1\x92\x7d\xaf\x0d\x8d\x11\x59\x97\x94\x68\x6c\xf7\xf1\xc9\xe2\x80\x27\x49\x89\x26\xc9\xde\x0b\xb0\x11\x6a\xd5\xbd\xa3\x66\x4b\x44\x93\xeb\xd1\xcd\xa4\xc3\xa4\x9a\x1d\x1b\xf2\xf8\x66\xd2\x3b\x42\x5f\x1b\x86\x76\xf3\xd7\x56\x24\xb5\x15\x97\xba\x9a\x0a\xc5\xd8\x22\xee\xb0\xe4\x39\xd2\xc0\x3d\x38\xe8\x8d\x94\xa3\x47\xbe\x67\x55\xc7\x0b\x57\x1c\xef\xc9\xb2\x76\xa4\xd1\xfd\x76\x5b\xc0\x20\x58\x71\x05\x08\x03\x18\x4f\x7a\xbe\x53\xc3\x21\x30\xce\x34\x23\x39\xfb\x1d\x4d\x89\x15\x30\xdf\x41\x0d\xb6\x1d\x6b\xa6\x3e\xd0\x56\xcd\x14\x46\xdd\x80\x31\xd3\x4b\x1f\x24\xfe\xd5\x74\xa4\x63\x54\x61\xbb\xeb\x29\x59\x47\x3d\x29\x6c\x55\xbf\x31\x5d\x12\xa6\x30\x18\x4f\xc2\xbd\x29\x4d\xa4\x86\x2d\xd3\x2b\x20\x90\xad\x88\x6d\x01\xa2\x44\x49\x6c\xde\x0b\x6e\x51\xc5\xce\x01\x06\x0a\x5c\xc4\x8e\x59\x78\x6d\x64\x61\x0a\xdf\x52\xab\xa5\x3e\x49\x2b\x98\x74\x6c\xb2\x05\x6c\x11\xa8\x80\x15\xd9\x20\x6c\x48\xbe\x46\x65\x30\x65\x89\xda\x5a\x23\x9b\x25\x2c\xa4\x28\xfa\x90\xa3\x7e\xa1\x40\x93\xef\x08\x4c\xa7\xc7\xb4\x68\xc9\x8c\xc3\x1c\x08\x07\x2c\x4a\xbd\x83\x9c\x29\xdd\x07\xa6\x61\x2b\xd6\x39\x85\x4c\xa2\x89\x1b\x81\x8f\xe4\x63\x98\x6b\x07\xb7\x53\xc5\x7e\xc7\x38\x49\xad\x2b\x71\x02\x33\x18\x75\xb2\x2e\x08\xa2\x27\xb8\xc8\x89\xfe\x40\xca\x38\xb2\xb4\x28\x49\x0b\x24\xdc\xd3\x34\x84\xf1\xc8\xfe\x85\x49\xda\x49\x30\xb7\xcd\x6c\x85\xd9\xf7\x7a\x5b\x36\x32\x04\x4a\x89\x1b\x26\xd6\xca\x56\x99\x49\x4f\xc2\x29\x28\x17\xa6\xc2\x84\x4c\xaf\x90\xc9\x50\x59\x26\xa4\xc4\x4c\xd7\x71\xed\x1c\x95\xc1\xab\x06\xab\x02\x84\x4b\x97\xa8\x63\xbd\x62\x2a\xe9\x96\x64\x2b\xd2\x09\x89\x35\x55\xee\xec\x89\x39\x30\xd0\x02\xc8\x46\x30\x0a\x85\xa0\x6c\xb1\x33\xbd\x87\x69\x13\xad\x9c\x64\x18\xc8\x1a\x57\xb2\xb5\x6c\x3c\xf9\xe5\xf3\xaf\x1f\xd3\x92\x48\x85\xb1\xfd\x68\x06\x05\xbe\x64\x8b\x9d\xcb\xa2\x24\x48\x20\x6b\x99\xb2\xc5\xc2\xd8\x6b\x03\xe4\x3c\x60\xbc\x6e\x44\xea\x9e\x39\xcd\x0a\x7c\x63\xa4\xa6\x10\x73\xdc\xc2\x1b\xa2\x31\x6e\x7c\x48\x3f\x8a\x6d\x02\x03\x68\x09\xed\xa6\x2d\xa5\x39\xc7\x74\x14\x3a\xe2\xd7\xfb\x41\x57\x07\x4f\x0e\xaa\x7c\x82\x51\xd9\xb8\x14\x44\x1c\x5a\x70\x76\x27\x68\x18\x62\x31\xbf\xfd\xcd\x7c\xe9\x83\x92\x99\xfd\x94\xc0\xbe\xd3\xb8\x25\xea\xb5\xe4\xf7\x96\x01\x5e\x8b\x35\xd7\xaf\xa0\xd1\x91\xda\xef\x06\xfd\x6a\x55\x6e\xa1\x7f\x4f\xce\x22\x8a\x27\xe7\x10\xc6\x93\xb3\x0b\x5d\xb9\x2a\x0c\x51\xd5\xbb\x77\x12\x4b\x29\xd6\x25\x52\x87\x29\x06\x30\x0a\x94\x4b\xfc\xca\xf4\xea\x10\x43\x4b\xec\x7b\xb1\xab\x17\xda\x90\x74\x52\xc2\x07\xd7\x06\x81\x7c\x3b\xc9\xa1\x52\xed\x6e\xa3\x24\x55\xeb\x22\x28\xd4\xc3\x69\xdc\x53\xdc\x60\xeb\x19\xcd\x96\xed\x41\xcd\xa7\x62\xe1\xae\x4f\xa7\x82\xe1\xa8\x7e\x34\x9a\x95\x53\xe1\x08\x1a\x46\xc7\x6b\x27\xfb\xc3\x01\x09\xbb\xcd\x39\xdd\xe7\x43\xe2\xa7\xca\xe1\x73\x88\x4d\xaa\xc6\xa6\xbe\x2b\xf3\x10\xa2\xea\xbc\x7f\x71\xa5\xe9\xec\xc5\x4b\x6f\x76\x79\xf9\xe2\x6a\xa8\xe9\xec\x5f\xfc\xc5\xcb\x60\x07\x35\xa7\x1b\x24\xce\x30\xb5\xf3\xc3\x23\x94\x79\xb3\xc5\x19\xee\x16\x37\x52\x2d\xfe\xc6\xee\x90\xc6\x17\xc9\x19\x91\x36\xc1\xff\xa0\x88\x3d\x80\x56\x66\xf4\x28\x19\xdb\xcb\x02\x3b\x50\xa8\x87\xc3\xd4\xe6\xdb\xe3\xdd\xf3\x12\xe9\xa8\x7f\x87\xd4\xa8\xbc\xf3\x7e\xdc\x1d\xd9\xfd\xb5\x05\x06\xd3\xc3\x4c\xab\xd2\xbf\x9b\xa5\x89\xa7\x11\x73\x85\x8f\x11\x6c\x13\x52\x05\xd2\x87\xeb\x05\xb8\xe6\x64\x0d\xb8\x7b\x74\xef\x9e\xca\xf7\x4c\xb9\x42\xd7\xe2\x13\x61\x52\xb9\xd2\x49\x52\x25\xa4\x8e\xdb\xb9\x9c\xf4\xe7\xc9\xbe\xce\x6c\xdb\x18\xdf\x71\x1d\x9b\x31\xdc\x74\xa9\x76\x61\x6e\x16\x26\x95\x57\xfc\xc3\x61\x3d\xf2\xf7\x01\xb9\xb9\x23\x9a\x81\x41\x62\x21\x36\x68\x67\x88\xc0\x1f\xea\x2e\x2c\x0a\x73\xcc\x74\x1c\xfd\x7f\xf8\x98\x91\xd4\x84\xbf\xe4\x79\x1c\xa5\x2d\x6d\x2e\xee\xa2\x24\x35\xba\xe2\x76\x3f\x7d\xef\x42\x91\xec\x9b\x8a\xa4\xe6\xd6\x00\xbe\x77\x34\x5d\xe9\x22\x8f\x8f\x5d\x4a\x7c\xae\xe7\x77\x29\xb5\xc1\x89\x93\x94\x94\x25\x72\x1a\x47\x5a\x46\x49\x6a\x9f\x2c\x90\xc6\x51\xa0\xb9\x0f\x5a\xae\x31\x79\x9c\xeb\x29\xde\x31\x1d\x27\xa9\x0b\x79\x9c\x3c\xf6\x51\x07\x82\xfb\x9a\xb9\x51\x3e\xd7\x55\xed\xf2\x28\xf9\xa7\xbd\xa9\x85\x4a\x3a\x17\xb0\x3f\xcb\x94\xdd\xbe\x10\xfc\x6f\x4e\xd9\xea\xa9\x46\xec\xa3\x59\xf0\x98\xf9\xfa\x78\xda\x85\x47\xf3\x2c\x43\x46\xfd\x3a\xf2\xb3\x8e\x22\x41\x0f\xf6\x12\xe3\xb5\xa8\x1f\x8b\x0f\x2d\x31\x80\xb1\x5b\xc1\x78\xdb\xe8\x9a\x61\xf0\xd0\xf2\xbc\x9e\xea\xf4\x24\x1e\x5c\x6f\xee\x5f\x6f\xea\xe8\x7f\x4b\x0b\x52\x3a\xfa\xf5\xf8\xc6\x13\x61\x9c\xa3\xac\x2f\x48\xdf\x71\xd7\xbd\x1b\xd9\x37\x0a\xfa\x0f\x34\x77\x77\x93\x7e\x9f\x6d\xe2\xd7\x8a\x46\x37\x49\x9a\x09\x9e\x11\x1d\x1b\xd1\x30\xfb\x6a\xc3\xd7\x6d\xc7\x6d\x65\xac\xa1\xbe\x53\xdb\x87\x83\x03\x01\xba\xfa\xbd\xaf\xba\xdf\xe2\x61\x1e\x7a\xca\x16\xb6\xc5\xc3\xff\x4d\xc1\x76\xf6\xce\x36\x6a\x67\x2c\xcb\xc0\x72\x9c\x4a\xec\x96\x73\x7c\xe3\x9e\x23\xf1\xb5\x28\x4a\x22\x31\x9e\x5f\x8f\x6f\x02\xa7\x26\xbd\xa7\x99\x1d\xda\xd7\xda\xf3\x3d\xba\x49\x8e\x53\x2d\xfa\xe2\xa1\xa9\xa1\x6d\x8d\x27\x06\x86\x27\xf7\xe3\xbf\x1c\x15\x8e\xfa\xdb\x9d\x12\x26\xbd\xe6\x99\xda\x1c\xc2\x52\x80\xc6\xa2\xcc\x0d\xe8\x21\xa7\xa6\x57\x19\xf4\xcf\x44\x51\x0a\x8e\x5c\x03\x53\xc0\xc5\xe1\x77\x88\xe6\x04\x0f\x63\xf2\x96\x71\x2a\xb6\x06\xd1\xde\x19\xe7\x37\x24\x3f\xa4\x9d\x9f\x55\xf4\x32\xbd\x55\x82\xc7\xd1\x7e\x9f\xce\x89\xc2\x6f\x25\xd1\xab\xaa\x1a\x9a\xf0\x0c\xf7\x7b\x6f\x64\xa9\x2a\xf3\x7d\x43\xa4\x4a\x19\xbd\xab\xaa\xa8\x1f\xbe\x46\xb7\xbb\xab\xfa\x70\x31\x1a\x8d\x0e\x79\x65\x1a\x18\xc9\x73\xd7\xa3\xec\x3b\x6a\xbe\xeb\x3d\x83\xf9\xfa\xcd\xdf\x3d\xe5\xdb\x1f\xe5\xdd\xfb\x7f\xe7\x67\x80\xff\x04\x00\x00\xff\xff\x86\xe7\xae\xf6\xd4\x1f\x00\x00")

func webTemplatesMonitorDetailsGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webTemplatesMonitorDetailsGoHtml,
		"web/templates/monitor/details.go.html",
	)
}

func webTemplatesMonitorDetailsGoHtml() (*asset, error) {
	bytes, err := webTemplatesMonitorDetailsGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/monitor/details.go.html", size: 8148, mode: os.FileMode(0644), modTime: time.Unix(1588163175, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xca, 0x3, 0x84, 0x14, 0x2f, 0x75, 0xc6, 0x8d, 0x27, 0x1e, 0xf8, 0xef, 0x4e, 0x16, 0x9f, 0x4c, 0xd3, 0x12, 0x31, 0xeb, 0xfa, 0xd5, 0x29, 0xf3, 0xcf, 0x73, 0xd8, 0xf7, 0x26, 0xc5, 0x70, 0xb7}}
	return a, nil
}

var _webTemplatesMonitorIndexGoHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x95\xc1\x8e\xda\x30\x10\x86\xef\xfb\x14\xa3\x88\xf6\x54\x12\x15\xed\x69\x1b\x90\xda\xaa\x42\x6c\xd1\x06\xb5\xa1\xd7\xca\xe0\x81\x58\xf5\xda\x96\xe3\x40\x2b\xcb\xef\x5e\x39\x90\x6c\x48\xa0\x39\xd0\x5d\xb5\x17\xb0\x3d\xbf\xc7\xdf\x4c\xc6\x1e\x6b\x29\x6e\x98\x40\x08\xd6\x52\x18\x14\x26\x70\xee\x26\xa6\x6c\x07\x6b\x4e\xf2\x7c\x1c\x68\xb9\x0f\x26\x37\x00\xcd\x35\x2f\x25\x4c\xa0\x2e\x2d\x6d\x1b\x1f\x3e\xd2\xe1\xdb\xd1\xd1\x06\x10\x67\xa3\xc9\x42\xcb\x35\xe6\xb9\xd4\x79\x1c\x65\xa3\xca\x62\xed\x60\x45\x72\xfc\xae\x88\xc9\xe0\x6e\x0c\x61\x3d\x73\xae\x96\x68\x22\xb6\x08\x03\x26\x28\xfe\x7c\x03\x03\x55\x79\x2a\x37\xd4\xb3\xbc\xde\xd1\x87\xd3\x56\x28\x22\x90\x43\xf9\x3b\xa4\xb8\x21\x05\x37\x27\xda\x33\xea\x61\x86\x84\x32\xb1\x6d\xe9\x7c\xa8\xb7\xa7\x42\xc3\x0c\xc7\x60\x12\x13\xc8\x34\x6e\xc6\x41\x33\x62\xe7\xa2\x1a\x3f\xb2\xf6\x10\xa1\x73\xc1\xc4\xda\xa7\x28\xc3\xa9\x26\x2a\x0b\xa7\x5a\x16\xca\xb9\x38\x22\x9d\x23\xa3\xec\xb6\x85\x1b\x51\xb6\xeb\x8b\x60\x25\xe9\xaf\x0e\xbe\xb5\xb0\x67\x26\x83\xc1\xd6\x1f\xea\x13\xdc\x06\x81\x46\x9a\x8f\xae\x0b\xde\x76\xd3\xf8\x6a\x48\xb7\x38\xa3\xfe\xbb\xf9\x51\xe9\xb1\xf4\x1d\xce\x84\x2a\xcc\x57\xa3\x91\x3c\xe6\x1d\x9f\x00\x31\x67\x93\xd9\xc3\x62\x99\xde\xf9\x22\xf1\x7b\xc3\x54\x2a\xb6\x76\x0e\x5e\x6b\xa2\xf5\x3b\xb0\x56\x69\x26\xcc\x06\x82\x57\x69\x70\x70\x1f\x7e\x94\x14\xd7\x3e\x4b\x9c\x9d\x63\x42\x41\xcf\x1c\xd5\xcf\x7a\x2f\x99\x30\x29\x59\x71\xbc\x84\x7a\x9f\xcc\x1e\xfe\x09\xd4\xb9\x94\x3f\x0a\xf5\x47\xd6\x79\x92\x7c\x5e\x2e\x5e\x02\xb6\xaa\x25\x79\x24\x6a\x70\x4e\xeb\xc5\x0b\x94\xd3\x2f\xc9\x72\x91\xbe\xff\x30\xff\x54\x92\x3e\x39\xe9\xe1\x6d\x08\xaf\xa2\xe6\x52\xaa\xd3\xbc\xaa\x43\xb5\x5e\xce\xea\x21\xa7\x7e\x63\x0f\x63\x29\x79\xa6\x02\x48\x0a\xd3\x7b\xb1\x92\x65\xfa\x02\x37\x2b\x8e\xba\x2f\x83\xb5\x80\x82\xb6\x1f\x91\xce\x83\xd5\x5a\x38\x99\x36\x0f\xab\x0d\xc7\xc1\xd5\x6d\xea\x1b\xc3\xfd\xd5\x1d\x6a\xc7\x70\x5f\x6a\xfd\xe0\xff\xec\x4b\x9e\xbc\xdd\x92\xfc\x5a\x55\x2a\xcf\xda\x87\xfe\x5a\x39\x1c\xff\x2a\xc9\xef\x00\x00\x00\xff\xff\x33\x4f\x5a\x19\xeb\x08\x00\x00")

func webTemplatesMonitorIndexGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webTemplatesMonitorIndexGoHtml,
		"web/templates/monitor/index.go.html",
	)
}

func webTemplatesMonitorIndexGoHtml() (*asset, error) {
	bytes, err := webTemplatesMonitorIndexGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/monitor/index.go.html", size: 2283, mode: os.FileMode(0644), modTime: time.Unix(1588163175, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1b, 0x17, 0x90, 0x50, 0xe6, 0xb2, 0xf, 0xf3, 0x84, 0x13, 0x19, 0x6e, 0x3d, 0xf8, 0xa6, 0xb4, 0x84, 0xe4, 0xbd, 0x9a, 0x77, 0x42, 0x20, 0x72, 0x74, 0xd3, 0xaf, 0xeb, 0xcf, 0x8c, 0x84, 0xf9}}
	return a, nil
}

var _webTemplatesMonitorMenuGoHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8e\x41\xca\x83\x30\x10\x46\xf7\x9e\x62\x08\x2e\xfe\x1f\x4a\xb2\x2f\x31\x5b\x2f\xd0\xbd\xa4\x3a\xd6\x80\x4d\x42\xa2\x6d\x61\x98\xbb\x97\x54\x70\x25\xa5\xab\x99\xc5\xe3\x7b\x8f\x68\xc0\xd1\x79\x04\x71\x47\xbf\x76\x7d\xf0\x0b\xfa\x45\x30\x57\x44\xc9\xfa\x1b\x42\xed\xfc\x80\xaf\x13\xd4\x31\x85\x1e\xce\x0d\xc8\xf2\x60\xce\x21\x65\xe6\x0a\x40\xcf\xce\x68\x0b\x53\xc2\xb1\x11\x44\xb5\xbc\xda\x8c\x5d\xb4\xcb\xc4\xac\x76\x56\x11\x6d\x4b\xcc\xc2\x10\x7d\xd6\x64\x9b\x6c\x9c\x64\x9b\xc2\x1a\x99\xe1\x6f\x87\xff\xb5\xb2\x46\xab\xd9\x99\x8a\x08\xfd\xf0\x35\xe7\xe1\xf0\xf9\x43\x49\xc1\x0e\x23\x2e\x21\xba\xbe\xe8\x0b\x71\x68\xde\xee\x3b\x00\x00\xff\xff\xfc\x37\x76\xe3\x2b\x01\x00\x00")

func webTemplatesMonitorMenuGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webTemplatesMonitorMenuGoHtml,
		"web/templates/monitor/menu.go.html",
	)
}

func webTemplatesMonitorMenuGoHtml() (*asset, error) {
	bytes, err := webTemplatesMonitorMenuGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/monitor/menu.go.html", size: 299, mode: os.FileMode(0644), modTime: time.Unix(1588163175, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd5, 0x53, 0xa1, 0x29, 0x98, 0xfe, 0x80, 0x2d, 0x4d, 0x92, 0x4b, 0xed, 0x4e, 0x17, 0x97, 0xf0, 0x64, 0xe2, 0xc3, 0x6, 0x32, 0xad, 0x9b, 0xa5, 0xd2, 0x13, 0x4, 0xe9, 0x7, 0x5a, 0xd1, 0x93}}
	return a, nil
}

var _webTemplatesQueryIndexGoHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\x4d\x6f\xe3\x36\x10\xbd\xe7\x57\x4c\x08\xa3\x39\xb4\xb2\xd0\xe4\xd6\x50\xea\xa5\x39\x15\x28\x50\xf4\xd0\x63\x40\x89\x23\x89\x28\x4d\xaa\xe4\x30\x8e\x21\xe8\xbf\x2f\xf4\x19\x5b\x96\xed\x2c\x16\x0b\xec\x45\xc9\x68\x3e\x38\xef\xcd\xe3\x58\x4d\x23\xb1\x50\x06\x81\xe5\xd6\x10\x1a\x62\x6d\x7b\xc7\xa5\x7a\x83\x5c\x0b\xef\x13\xe6\xec\x9e\xa5\x77\x00\x00\xc7\x6f\xbb\x60\xa1\x0c\xba\xd1\xb7\xf4\x7f\x64\xad\x79\x73\xab\xa3\x9d\x8c\x7e\x7d\x5c\xc4\x00\xf0\xea\x31\xfd\x3b\xa0\x3b\xc0\x9f\x78\xf0\x3c\xae\x1e\x97\x11\x4d\xa3\x0a\xd8\xa2\x73\xd6\xb5\xed\x32\xfb\xe8\x0c\xa1\xd1\x11\xf4\xcf\x48\x0a\x53\xa2\x9b\x0c\xe5\x77\xca\x7b\x95\x69\x64\xe0\xac\xc6\x31\xf6\xac\x17\x00\x9e\x05\x22\x6b\x80\x0e\x35\x26\x6c\x30\xd8\x0c\x42\x5b\x8f\x0c\xa4\x20\x31\xd5\x9c\x2b\x71\x5f\x0b\x93\xfe\x44\x6a\x87\xfe\x99\xc7\xbd\xc5\xe3\xa1\xc0\xca\x31\x9e\x9c\x35\x65\xfa\xd2\x81\xba\xe7\xf1\x68\x42\xd3\x0c\x38\xb7\x2f\xeb\x68\x63\xa9\xde\xce\xe9\x41\x23\xcf\x42\x07\xd2\xf6\xc2\x19\x65\xca\x4f\xd3\x36\xc6\xff\xf0\xbc\xfd\x3b\xf4\x79\xca\xdc\xd8\xfc\xb7\x73\x37\xb1\xe7\x6d\x70\x39\xfa\x15\x3f\x2f\xac\xdb\x81\x35\x3e\x64\x3b\x45\x09\xdb\x2b\x23\xed\x7e\xab\x6d\x2e\x48\x59\x03\x09\x3c\x34\xcd\x36\x13\x1e\x5f\x6b\x41\x55\xdb\xc6\x4d\xb3\xf5\xa8\x31\x27\x94\xaf\x43\xdd\xb6\x8d\x1f\xe0\x67\xf0\x28\x5c\x5e\x6d\xdf\x84\x0e\xf8\x0c\x0e\x29\x38\x03\x85\xd0\x1e\x9f\x57\x78\x5e\x0e\x4f\x99\x3a\x50\x54\x3a\x1b\x6a\x38\xfa\x3f\xd2\xe5\x85\xe4\x2b\x05\xa2\x8c\xcc\x95\xac\x9b\x63\xce\xc8\x40\x46\x26\x92\x58\x88\xa0\x09\xa4\xb3\xb5\xb4\x7b\x13\x91\x2d\x4b\x3d\x09\x60\x30\x12\x36\x79\x59\xba\x46\x4d\x2f\x8b\x59\x3f\xc2\x61\x27\x95\x5b\xea\x38\xe9\x34\xe8\x29\x7d\xee\x63\x87\x26\xdc\x00\x08\xfd\xf0\x37\xf3\xe4\xe0\xb7\x04\x8e\xe7\xf8\x89\x64\xd7\xad\x1e\xd8\x28\x23\xf1\xfd\x17\xd8\x0c\x88\xfa\x3a\x97\xf5\x74\xd6\xbd\x56\x29\x17\x50\x39\x2c\x12\x76\xdc\x4f\xaf\xa4\xcd\xc4\x52\xc7\xdd\x6c\xf0\x58\xa4\x3c\xd6\xea\x33\x00\xd7\x55\x7f\xd2\x41\x1c\xf4\x15\x05\xad\x5d\xa6\x23\x77\xaf\x2a\x50\x32\x61\x83\xbc\xd9\xa8\x18\xc2\x77\x9a\xf5\xd2\xdd\xa1\xa8\xfb\x59\x71\x56\x33\x70\xf8\x7f\x50\x0e\xe5\xf7\x97\xed\x70\x69\x2f\xc9\x96\xa5\xff\xf4\x2d\xdf\x96\xd9\x15\x12\x2e\xb8\x78\xdc\x61\x3e\x7f\xdf\x34\xa8\x3d\xae\xed\x99\xeb\x7b\x7a\xb1\x95\xff\xb2\x30\x4a\x0c\x0a\x1b\x8c\xbc\x87\x3f\x94\x84\x83\x0d\x50\x58\x57\x22\x01\x59\x10\x44\x22\xaf\x80\x2a\xdc\xfd\x7e\xa1\xcb\x35\x79\x2c\x42\x17\xe6\xb0\x2e\xfb\x15\x76\x94\xc7\x2b\x17\xaf\x7f\x2b\xd4\xc2\xa0\x86\xfe\x39\x6f\x8b\xc1\xf2\x21\xcf\xd1\xfb\x2b\x5f\x12\x43\x5c\x85\x42\x76\x04\xac\x70\x5c\x3d\x9d\x86\x92\x22\x8d\xfd\x96\xf9\x0f\x0f\xdd\x35\xa9\x9e\xd2\x6b\xd8\xd6\x0f\xcc\xac\x3c\xac\x9d\x56\x3b\xec\x4a\x8f\xd8\x79\xdc\xd9\x5f\xc5\xdc\x07\xd7\xa3\x6b\xfc\x33\x79\xbe\x04\x00\x00\xff\xff\x05\xda\xbe\x9a\xaf\x09\x00\x00")

func webTemplatesQueryIndexGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webTemplatesQueryIndexGoHtml,
		"web/templates/query/index.go.html",
	)
}

func webTemplatesQueryIndexGoHtml() (*asset, error) {
	bytes, err := webTemplatesQueryIndexGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/query/index.go.html", size: 2479, mode: os.FileMode(0644), modTime: time.Unix(1588163175, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4d, 0x7, 0xb0, 0x5b, 0x3e, 0x3a, 0xfb, 0x8, 0xd8, 0xf3, 0xeb, 0xbc, 0x13, 0xc2, 0xd3, 0xd, 0xc8, 0x7e, 0x61, 0x63, 0x31, 0x24, 0xaf, 0x96, 0xd2, 0x5e, 0x8f, 0xb2, 0x58, 0x82, 0xd3, 0xe6}}
	return a, nil
}

var _webTemplatesIndexIndexGoHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xcf\x4a\xc4\x30\x10\x87\xef\x7d\x8a\x21\xec\xd1\x6e\x51\xf6\x24\x6d\x0f\x3e\x80\xf8\x0a\x43\x67\xba\x09\xa4\xd3\xa5\x8d\xab\x30\xcc\xbb\x8b\x4b\xac\xa6\x0a\x5e\x92\xcc\x7c\x5f\xfe\xfc\xa2\x4a\x3c\x06\x61\x70\xc3\x2c\x89\x25\x39\xb3\xaa\xa5\x70\x85\x21\xe2\xba\x76\x6e\x99\xdf\x5c\x5f\x01\xfc\xec\x7d\xaa\x18\x84\x97\x1b\xd9\xb3\x58\x4f\x54\xdf\x3f\x64\x06\xa0\xba\xa0\x9c\x19\x0e\x41\x88\xdf\xef\xe0\x30\xcc\xd3\x65\x16\x96\x04\x8f\x1d\x1c\xb7\x6a\x35\xcb\x3b\xfe\x3b\x71\x6f\x5c\x50\x38\xc2\x6d\xac\x89\x47\x7c\x8d\xa9\x70\xff\xb0\x6b\xcf\x48\x41\xce\x3b\x0f\xa0\xf5\xa7\x52\x4c\x21\x45\x76\x7d\x8b\xe0\x17\x1e\x3b\xa7\xfa\x1d\xe0\xf8\x84\x2b\xbf\x60\xf2\x66\x8d\xeb\x0b\xf2\x8c\x13\x9b\xb5\x0d\xfe\xba\xa0\xf1\xa7\xbe\x2a\x5f\xd7\x50\xb8\x16\xe1\xca\x46\x51\xaa\xb2\x50\xfe\xaa\x0d\xe4\x45\x9e\xbe\x94\x8f\x00\x00\x00\xff\xff\xcd\x13\xfe\x16\xdd\x01\x00\x00")

func webTemplatesIndexIndexGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webTemplatesIndexIndexGoHtml,
		"web/templates/index/index.go.html",
	)
}

func webTemplatesIndexIndexGoHtml() (*asset, error) {
	bytes, err := webTemplatesIndexIndexGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/index/index.go.html", size: 477, mode: os.FileMode(0644), modTime: time.Unix(1588163175, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf9, 0x5b, 0x7c, 0xfd, 0x89, 0x8e, 0x47, 0x41, 0xe7, 0x99, 0xac, 0x38, 0x42, 0x38, 0x90, 0x60, 0x3b, 0x85, 0xc, 0x61, 0xe0, 0xc1, 0x2e, 0xc5, 0x69, 0x51, 0xde, 0xd5, 0x61, 0xb2, 0x2f, 0xf7}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"web/templates/common/base.go.html":     webTemplatesCommonBaseGoHtml,
	"web/templates/common/head.go.html":     webTemplatesCommonHeadGoHtml,
	"web/templates/common/menu.go.html":     webTemplatesCommonMenuGoHtml,
	"web/templates/monitor/details.go.html": webTemplatesMonitorDetailsGoHtml,
	"web/templates/monitor/index.go.html":   webTemplatesMonitorIndexGoHtml,
	"web/templates/monitor/menu.go.html":    webTemplatesMonitorMenuGoHtml,
	"web/templates/query/index.go.html":     webTemplatesQueryIndexGoHtml,
	"web/templates/index/index.go.html":     webTemplatesIndexIndexGoHtml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"web": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"common": &bintree{nil, map[string]*bintree{
				"base.go.html": &bintree{webTemplatesCommonBaseGoHtml, map[string]*bintree{}},
				"head.go.html": &bintree{webTemplatesCommonHeadGoHtml, map[string]*bintree{}},
				"menu.go.html": &bintree{webTemplatesCommonMenuGoHtml, map[string]*bintree{}},
			}},
			"index": &bintree{nil, map[string]*bintree{
				"index.go.html": &bintree{webTemplatesIndexIndexGoHtml, map[string]*bintree{}},
			}},
			"monitor": &bintree{nil, map[string]*bintree{
				"details.go.html": &bintree{webTemplatesMonitorDetailsGoHtml, map[string]*bintree{}},
				"index.go.html":   &bintree{webTemplatesMonitorIndexGoHtml, map[string]*bintree{}},
				"menu.go.html":    &bintree{webTemplatesMonitorMenuGoHtml, map[string]*bintree{}},
			}},
			"query": &bintree{nil, map[string]*bintree{
				"index.go.html": &bintree{webTemplatesQueryIndexGoHtml, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
