// Code generated by go-bindata.
// sources:
// config.yaml
// templates/api.gotmpl
// templates/config.gotmpl
// templates/parameter.gotmpl
// DO NOT EDIT!

package main

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

var _configYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\x4f\x4b\xc3\x30\x18\xc6\xef\x7e\x8a\x97\x9e\xb6\x83\x9d\xe7\xde\x04\x15\x14\x94\x81\xbb\x6f\xaf\xed\xd3\x18\x97\x7f\x24\x99\xa2\x25\xdf\x5d\x9a\x6e\x93\x4d\xd1\x0d\xd9\xed\x6d\xd3\xdf\xd3\x5f\x9e\x44\xf1\xbb\x5d\xc5\xea\x8c\x88\x9d\x53\xb2\xe6\x28\xad\xe9\x1f\x89\xce\xc9\xb0\x46\x45\xb5\x35\xad\x14\x2b\x8f\xfc\x96\x28\xd8\x95\xaf\x51\x51\xd1\x75\x54\xce\xb4\xbb\x92\x9e\x52\x9a\x44\x68\xa7\x38\x22\x4c\x06\xa0\x14\x36\x6a\xa7\x8a\x35\x15\xd9\x0b\xc4\x4c\x2d\xba\x8e\x5e\xac\x34\x37\x52\x61\xca\xf1\x99\xca\x59\x5e\xa4\xf2\x11\xfe\x15\x7e\xca\xf5\x92\x05\x28\xa5\x45\x4a\x1b\xbe\x95\x0a\xf3\x41\xa8\xd8\xfe\xa0\xd8\x11\x85\x7e\x42\xd3\xa0\x99\x07\x87\x7a\x4f\x96\x43\x40\xac\xc2\x1b\x0b\x01\x7f\x17\xac\xb9\xee\x3f\x3e\x81\xdb\x8e\xc4\x37\xc5\x90\x43\x0e\x2d\x92\x9d\x3c\x61\x8b\x43\x7a\xbf\xa2\x6d\x03\x15\x76\x4f\xbd\x41\x2b\x8d\xec\x6f\xc3\x8f\x4d\x66\xe4\x18\xa9\xfb\x1e\xf8\xd3\x69\x48\x18\x05\xc3\x4b\xf9\x01\x1a\x39\x0e\x35\xab\x7e\x2c\x1f\x58\x63\x3c\x1e\xd0\xb5\xb7\x75\xf0\xf9\xc2\xee\xb9\x3b\xf6\xac\x11\xe1\xc3\xa1\x4d\x6f\x89\xdf\xfb\x96\x2d\x89\x48\x23\x05\xd3\x6f\x4c\x84\x31\x5d\x50\x4a\x07\x1e\x44\x79\x39\xbd\xdd\xce\x9b\x61\xc0\xa1\x02\x8e\x08\xda\x83\x4d\xf3\xef\x42\xe7\x5f\x95\xe5\x6e\x3f\x03\x00\x00\xff\xff\x82\x4f\xcf\x1e\x17\x04\x00\x00")

func configYamlBytes() ([]byte, error) {
	return bindataRead(
		_configYaml,
		"config.yaml",
	)
}

func configYaml() (*asset, error) {
	bytes, err := configYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.yaml", size: 1047, mode: os.FileMode(420), modTime: time.Unix(1508518995, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesApiGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x58\x5d\x73\xdb\xb6\xd2\xbe\x26\x7f\xc5\x96\x93\xa4\xa4\x47\xa6\xda\x5e\x6a\x5e\x5f\xe4\x75\xd2\xc4\xa7\x49\xea\xb1\x9c\x93\x8b\x4e\xa7\x85\xc9\x15\x85\x63\x0a\x60\x01\xd0\xb2\xc2\xe1\x7f\x3f\xb3\x0b\x90\xa2\x64\xc5\x3d\xf5\x85\x45\x02\xfb\x85\xdd\x67\x3f\xc0\x46\x14\xf7\xa2\x42\xe8\x3a\xc8\x5f\x5f\x5f\x5d\x87\xd7\xbe\x8f\x63\xb9\x69\xb4\x71\x90\xc6\x51\x52\x68\xe5\xf0\xd1\x25\x71\x94\x28\x74\xf3\xb5\x73\x0d\x3d\x6b\xeb\xff\xcf\xad\xac\x94\xa8\xe9\xc5\xee\x6c\x21\x6a\x7e\x74\x72\x83\xbc\xe4\x8c\x54\x95\x4d\xe2\x38\x4a\x2a\x5d\x0b\x55\xe5\xda\x54\xf3\xc7\xb9\x16\xad\x5b\xff\xe4\xd7\xa5\x5b\xb7\x77\x79\xa1\x37\xf3\x4a\xaa\x73\xd2\x67\xe4\xdd\xbc\x69\x8c\x5e\x25\x4f\xf7\x2b\xad\x64\x41\x4f\x49\x1c\x55\x52\x79\x41\x30\xa5\xfa\x2a\x6a\xa1\x4a\xcd\xd4\x83\x9e\x03\x31\x06\x9d\x6d\x7f\xe0\x7d\xbb\x15\x55\x85\x66\xbe\x91\x65\x59\xe3\x56\x18\xfc\x7b\x5a\xd1\xc8\x24\x8e\x6a\x5d\x1d\x68\xb5\xd2\xb4\x8d\x45\x35\xaf\x75\x65\x5a\xf2\x4e\xd7\x19\xa1\x2a\x84\xfc\x0d\xae\x44\x5b\xbb\x2b\x76\xaa\xed\xfb\xae\x6b\x8c\x54\x6e\x05\xc9\xcb\xbf\x12\xc8\xfb\x9e\x68\x51\x95\xfe\xc1\x33\xbd\xb8\xc7\xdd\x0c\x5e\x3c\x88\xba\x45\x58\x5c\x40\x3e\xe1\xa6\xbd\xbe\xa7\xb8\x4d\xe5\x78\xda\xa9\xb0\x2c\x8e\xe7\x73\xb8\xd1\xad\x43\x0b\x25\xae\xa4\x42\x0b\xa2\xae\xc1\xad\x11\x8c\x5f\xd6\x2b\x7e\x5b\xa2\x79\x40\x03\x16\xcd\x83\x2c\x30\x8f\xdd\xae\xc1\x81\xd3\x3a\xd3\x16\x0e\xba\x38\x3a\xab\xa4\xca\xdf\xaa\x4a\x2a\x9c\x1c\xef\xd7\x06\x8d\x70\x52\x2b\x36\x0e\x1a\x41\x30\x90\x5f\x11\xf2\x4f\x62\x43\x78\x9a\x48\xf0\x22\x58\xb0\x79\x67\x74\xdb\xc4\x51\xd4\x75\x20\x57\x90\xbf\x6e\xdd\x5a\x1b\xf9\x15\x4b\xe8\x7b\x7a\x01\x22\x7d\x2f\x54\x59\xa3\xf9\xb9\x55\x05\x8c\x4e\x8a\xae\xb5\x75\xf0\x54\x54\x1f\x0f\x24\x3d\x1f\xbd\xd0\x6a\x25\xab\xd6\xe0\x17\xac\xeb\x5f\x94\xde\x2a\x40\x25\xee\x6a\x72\x83\x2a\xf7\xdb\x16\xe6\xf9\x16\xeb\xfa\xfc\x3e\xd0\x94\x8d\x96\xca\xd9\x3c\x5e\x91\xde\xd4\xc0\x99\xf7\x45\x76\x42\x64\xba\x46\x51\xbb\x35\x5b\x48\xe4\x69\x06\x77\x5a\xd7\x19\x1d\x77\x3b\xea\x5d\x5c\x80\xc9\xd9\xcc\x34\x99\x2a\x4b\xb2\x38\x22\xbf\x8c\x94\xf9\xbb\xb7\xb7\x69\x32\xb7\xc5\x1a\x37\xe2\xbc\x94\xb6\xd0\x0f\x68\x76\xc9\xcc\x0b\x2f\xdc\xa3\x3f\xf8\xa5\xcf\x4b\x56\x13\x45\x23\x1d\x29\x9a\xb8\x3b\x8a\x96\x2c\xe8\xf3\xcd\x07\x00\x9f\x8d\xf0\xe7\x7f\xac\x56\x8b\xc4\x6b\xf8\xa3\x35\x75\xf2\xe7\x84\xf2\x96\x22\x7f\x92\x92\x30\x11\x48\x3f\x5f\xb1\x40\xfa\x3b\x24\x6d\xe5\x5e\x60\x7f\x64\xc0\x02\x20\x99\x87\x24\xca\x89\x3c\x99\x1d\xe9\x5d\x40\x12\xf6\xcf\x7f\xca\x7f\xf0\xdb\x14\xee\xa8\x70\x8f\xf9\xbf\x96\xbf\x7e\x4a\xa9\xfc\xe4\x4b\x27\x5c\x6b\x7f\xfd\x65\x06\xaf\xc6\x73\x67\x71\x14\xf5\xd9\x09\x47\xfa\xe8\x24\x33\xf0\x0f\x01\x4f\x93\xa0\x65\x19\x01\x27\x8e\xcc\xe0\xf9\x03\x0b\x9f\xf3\x3a\x59\xb5\xe4\xe3\x1f\xdb\xe5\x9d\x92\x2e\xbd\x28\xb2\x9c\xb5\x64\x01\x97\x07\xa6\x80\xb4\x9c\x82\x7e\x11\xde\xdf\xde\x5e\xc3\x3a\x6c\xb5\x16\x4b\x58\x69\xc3\x04\x53\xd8\x84\x63\x91\x30\x4e\x64\xca\x20\x0f\xed\x32\x80\xf6\x5b\xc7\x3d\xc4\xe8\x93\x14\x8b\x23\x4f\xca\x40\xda\x73\xc5\x91\x41\xd7\x1a\xf5\x2c\x08\xc3\x09\x8e\x11\xf8\xde\x2f\x93\xc2\x01\x26\x21\x28\x53\x98\x78\xaa\x45\xd0\xb9\x4b\xb3\x10\x7d\xfa\x2f\x57\xf0\x9d\x5f\xcf\x83\x30\xcf\x73\x0a\x16\x4b\x5f\xc2\x3e\x2b\xf1\x20\x64\x4d\x1e\x99\xc1\x2b\xcf\x9c\xb1\x44\xc0\xda\xe2\x33\x02\x18\x57\x53\x06\x82\x87\x8f\x5b\x90\x3d\x44\x4c\x2a\x87\x66\x25\x0a\x04\xb7\x16\x0e\x36\xad\x75\x70\x87\x20\x37\x4d\x8d\x1b\x54\x0e\x4b\x90\x0a\xb4\x29\xd1\x80\xd3\xd0\x18\xfd\x20\x4b\x24\x41\x77\xad\xa5\x62\x6c\xa1\xd6\x95\x2c\xc6\x08\x9f\xac\xc2\xa3\xd2\x51\x5b\x17\x07\x6f\xed\x42\x1c\xff\x49\x31\x7e\x12\xbc\x50\x7e\xaf\x85\x11\x1b\x0b\x7d\x3f\x83\xc6\x3f\x9e\x75\x5d\x1e\x86\x82\xbe\xcf\x4f\x0a\xf3\x4c\x5d\x47\x25\x13\xfa\x3e\x83\x33\xd1\xc8\xfc\x06\x6d\xa3\x95\xc5\x69\x35\x66\x4c\x56\x52\xc9\xaf\x78\x2d\xdc\x3a\x6d\x84\x5b\x87\x34\xc9\x86\x1a\xd2\x8d\x28\x0b\x73\x43\x7e\x83\x4d\x2d\x0a\x4c\x8f\xdf\x89\x7b\x06\x49\x97\xcc\x20\x59\x24\x33\x38\xff\x31\x9b\x41\xd2\xd3\xab\x7f\x3b\x6e\x00\xa1\x95\x4d\x2a\xfe\xa4\x07\x7e\xcb\xfd\x6c\xf3\x91\x88\x34\xec\x0e\x61\x99\x85\xb4\xe3\x86\x45\xb1\x98\x81\xd3\xf7\xa8\xa8\x3e\x0e\xc7\x0b\xcd\x83\xce\x87\xdc\x3c\x29\x45\x28\x02\x9f\x70\x9b\x66\xc3\x62\xfe\xd9\x62\xca\x2d\x0d\x7d\x49\x4b\xb3\xc3\xbd\xfd\x8c\x92\x7f\xe0\x19\xe3\x83\xa6\xea\x42\x64\x31\xa5\xc8\x7e\xdf\xc2\x77\x17\xa0\x64\x0d\x5d\x0c\xe1\x8f\xce\xf8\xc7\x0c\x1e\xb8\x11\x31\x54\xa6\xe4\x7b\x3a\xfa\x9b\xe8\x7c\xc8\xc6\x9d\x3e\xf6\xff\xe3\x28\xb8\x6d\x71\x01\xaf\xfc\xc9\x3a\x3f\x13\x2c\x02\x27\xd1\x9c\x04\xa4\x67\x3c\x0d\xa5\x69\x27\x87\x8b\x10\x9a\xa1\x63\x72\x97\x64\x9c\x2a\x84\xfc\x23\xba\xb5\x2e\x21\x79\xf7\xf6\x36\x81\xff\x5d\xea\xb1\x13\x19\xff\xca\x51\xe3\xb1\xe9\x38\x77\x49\x55\xe2\xe3\x0c\x5e\xb8\x5d\xc3\x73\xd7\xa5\x56\xb6\xdd\xa0\xfd\x88\xa5\xe4\x26\xc5\x89\x25\x57\x81\xb2\xef\x67\x01\xe5\x49\xd7\x11\x13\x3f\xf0\x42\x96\xc5\x63\x66\x9c\x1c\x72\xc6\x9a\xcd\xe0\xe9\xfc\x28\xf4\x82\x8f\xc3\xe6\x2f\x2e\x20\x3d\x3a\x52\x06\x3c\xff\x10\x9d\xc5\xa2\x35\xd2\xed\xde\xd0\x60\x27\xd9\xc3\x6c\xf0\xf2\xc4\xfa\xc0\x14\x62\x32\x90\x00\xd7\x88\x17\x25\xae\x88\x93\xcf\x73\x5a\xee\xe0\xce\x71\x5a\xc3\xbf\x98\x6d\x94\xc4\x3d\x1c\xaf\x8d\x6e\x6c\xce\x13\x44\x12\x26\x6f\xcf\xc4\x67\xba\x1d\xf2\x62\x71\x31\xe6\x48\xcc\xc5\xfd\x70\xfb\xe2\x02\x92\xc4\x17\xe8\xa3\x0d\x38\x9c\x9c\xbf\x6d\x41\xe0\x60\xdd\xa3\x7e\x86\xc8\xc4\xbd\x7d\xcf\xe1\x00\xce\x45\x6f\x2e\x2f\xa4\xa4\x79\x82\x93\x65\xa1\x1b\xb4\xbc\xf3\x04\x25\x96\xf6\xbc\xdb\x99\xea\x19\x6c\x30\xe9\x04\x1d\xdc\xdf\x82\xda\xb7\x61\xe2\xf4\x5d\x69\x30\x7f\x71\xe8\x18\xdf\x10\xe9\x7f\xe6\xe3\xe0\xa1\xf5\xcc\xf1\x8e\xc1\xff\x9c\x1b\x8e\x64\x8e\xa3\x74\xf4\x6c\x7a\xf1\x20\x3e\x66\xeb\x69\x9a\x49\x0e\xff\xad\xac\xbc\xeb\x42\x7a\xf7\x7d\x3a\x69\x17\x47\x57\x26\x5a\x23\x17\xc2\x41\x8b\x1a\xbd\x78\xb2\xef\x0d\x65\xfd\xe4\x2e\xd5\xd0\xf1\x7a\x11\x9a\x90\x37\x76\xd2\xfa\xd1\x8c\xb7\xa8\x67\xbb\x35\x2d\x8e\x03\x50\xa8\xfd\xa1\x07\xc4\x91\xef\x27\x70\x76\xc9\xbf\x71\x64\x3d\xc3\x99\x9f\x40\xf8\x25\x8e\x6e\xa5\xab\x87\x29\x3c\x8e\xfe\x8d\xc6\x4a\x3d\x74\xc5\x60\xd0\x27\xdc\x06\x5d\x9c\xa5\x7c\x1c\x0b\x02\x14\x6e\x4f\x37\xb2\x91\x21\xb5\x0f\xc5\xbe\x7d\x1d\xda\x33\x3b\x68\x0b\xbf\xfd\x7e\x34\x1f\x66\x70\x16\x64\x77\xdc\x72\xbe\xf3\xdc\xf9\x1b\xbc\x6b\xb9\x7f\xd3\x95\x3c\x5f\xa2\xfb\xa8\xcb\xa1\x99\xd5\x28\x2c\xd2\x7b\x98\xb3\x45\x23\xb9\x75\x78\x41\xc4\xe3\x5d\xb3\x78\xda\x6a\x1f\x8a\xd9\xa8\x82\x30\xfa\x46\x5a\x1e\x71\x07\xab\xc7\x44\x3f\x30\x9b\x33\xcb\x13\x0c\x32\x69\x85\x7d\xba\x80\xa4\xeb\x20\xbf\x52\x2b\x9d\x7b\x27\xf7\x3d\xdf\x33\x82\x8f\xa7\xfb\x83\xdb\x3d\x05\x99\x3e\x9f\x87\x82\x0d\xfc\x75\x02\x28\x66\xfb\xdb\x22\x0d\x7b\x25\x3b\x62\xa3\x4b\x64\xff\x3c\x71\x0f\xf3\xe5\x37\x58\x49\xeb\xd0\xa4\x3c\x2d\xf9\x9c\xf0\x2d\x74\x46\x6d\x3b\x1b\xb5\x59\x74\xe0\xbf\x27\xd0\x4f\xe5\xe7\xc8\x5b\x7c\x74\x3f\x6b\xb3\x11\xce\xa1\x81\xad\x74\x6b\x50\x1a\x0a\x5d\x6b\x63\xf9\xbb\x04\x05\x60\x24\x48\x5f\xd1\xca\x01\x4f\x17\xfc\x78\xc9\x2c\x0b\x70\xa6\xc5\x3e\xf3\xa1\xc9\x03\x1e\x2f\xe0\xd5\x04\x91\x64\xfa\xeb\xb2\x34\x8b\x71\x46\x18\xa2\x52\x96\x06\xad\x25\x0f\x06\x98\x04\x92\xa7\x27\xa3\x40\xa3\x28\x6f\xe5\x06\x75\xeb\x16\x00\x3f\xfe\x00\x67\xe0\xe4\x06\xa9\x86\x6b\x55\x12\xc5\x17\x23\x1d\x8e\x24\x27\x28\xfa\xf8\x00\x78\xe3\x2d\x7c\x80\x06\x7b\x79\xa2\xfc\xc4\x7d\xdd\x3e\x14\xe1\x1a\xb1\xdb\x7b\x7a\xa4\x0b\x17\x90\xaf\x63\x5c\xe3\xa9\xb8\xe9\x9d\xf2\xeb\x93\x4b\xe5\x54\x32\xf9\x33\x14\x12\xd1\xc8\x90\xb4\x37\xad\x02\xd3\xaa\x27\xf5\x83\xb2\xd8\xc1\x56\xd6\x35\xd4\x84\x0c\x05\x5a\x01\x4a\xb7\x46\xe3\x2f\x84\xda\xff\x2e\xa1\xc4\x06\x55\x49\xd3\xb2\x56\x24\x91\x04\x85\x14\x6e\x84\xa5\xeb\xa2\xd3\xfb\x54\x1f\x3e\x60\x88\x21\x6f\x33\x32\x21\xcd\x00\x8d\xd1\x9c\xc4\xa7\xe1\x42\x77\xa2\x3d\x5c\xe8\x5e\x4d\xab\x94\x15\xab\x34\xe1\xd2\xa1\x2a\xf8\xfe\xa5\x85\x73\x78\x69\xbf\x27\x63\x85\x07\x02\xbc\xb4\xc9\x0c\x84\xcf\x2d\x7a\x08\x49\x44\x8f\xfe\xa0\x0c\x99\x2c\x8e\x38\x8c\xc3\x5c\xff\x5e\xd8\x6b\x83\x2b\xf9\x98\x1e\x90\xcd\x20\x69\x95\x7c\x5c\x24\xfe\x82\x39\xb8\x33\xbf\x69\xd5\x67\x25\x1f\x53\x0e\x1f\xc9\x11\x21\xce\xf9\x95\xe2\xf1\x05\xd9\x69\x07\x3c\x41\xec\x07\x76\xef\x6b\x55\xb2\x3f\x82\x84\xe7\x89\x6e\x3f\x2c\xd3\x51\xc1\xed\x87\xe5\x25\x1a\xf7\xb3\xf4\xc7\xdb\xaf\xfe\x82\x3b\x5a\x1c\x6e\x1e\xcb\x75\xeb\x4a\xbd\x55\x3e\xa8\x95\x11\x05\xae\xda\xba\xde\x81\x1d\x36\x9e\x40\xe0\x44\xb0\x06\x29\xd3\x88\x1d\x1b\x3b\xd2\x84\x2f\xb5\xf9\xff\x8b\xe2\xbe\x32\xba\x55\x25\x5d\x0d\x46\xe0\x7d\x91\x6e\xbd\x94\xd5\xf0\xd5\xe1\x1b\x30\xe4\x72\xb2\xbc\x7a\x77\xfb\xf6\xe6\xa3\xff\x0c\x41\xa1\x16\xad\xd3\x1b\xe1\x64\x21\xea\x7a\x17\x8f\x65\xb0\xcc\xe1\x76\x8d\x7b\xd6\x3d\x7c\xe9\xba\x21\x46\x39\xfe\x1b\x31\x7f\x73\x3b\xe1\x89\x01\xc6\x5b\xbc\x1b\x3d\x31\x9f\xc3\x95\xfb\xde\x42\xa3\xad\x95\x54\x70\x9d\x06\xdd\xd0\x30\x4a\x16\x30\xd4\x41\xa8\x1d\xa8\x76\x73\x47\xc6\x0f\x4e\x25\x0f\xfa\x91\x75\xbb\x96\xc5\x9a\x6d\x62\x7b\x1f\xb1\x68\x1d\x82\x56\x08\x77\x3b\xfe\x11\x2b\x2a\x9f\x41\x71\x38\xc2\x5a\x58\xb8\x43\x54\x7b\x81\xb6\x2d\x0a\xb4\x96\x2d\x3e\x9d\x4d\x87\x7e\x4d\x47\xce\x3c\xcf\xc3\xa7\x16\x0e\xdd\x24\x82\x56\x56\x97\xfc\x99\x64\x23\xee\x31\x2d\xd6\x42\x81\xb6\xf9\x92\x9d\x34\x83\x1f\x33\xa6\x50\xa2\xce\x3f\x69\x27\x57\xbb\x94\xe9\x67\x10\xbe\xaf\xe7\xc1\xab\x54\x60\x2a\x3d\x7c\xce\x21\xac\xff\xdf\x39\x53\x52\x01\xdc\xa3\x82\x10\x9e\x12\x2d\x1a\x43\x3a\x39\x7b\x68\x95\xa6\x78\x63\xf6\x37\xc5\x68\xb2\xc2\xa5\xff\xad\x31\xfe\x98\x97\xb5\xb6\xa1\xb2\x0e\xe8\x43\x63\xe2\xf0\x69\x24\x8e\xc2\xdd\x72\xa5\xf6\x97\xcb\xd1\x0b\xc4\x14\x34\xaf\x58\xed\x09\xbd\x27\x85\x86\x35\x25\xeb\x3d\x84\x29\xe9\x3d\x70\xe9\x82\x0c\x42\x01\x95\x07\xb0\xba\xb8\x47\xb7\x0f\xce\xeb\xeb\xab\x0c\xc6\x1a\xb1\xf7\x3b\xd5\x28\xef\x82\x49\x8d\x89\xa3\x95\xac\x79\x84\xa7\xed\xdf\x6a\x54\xe9\x50\x74\x16\xbf\xc7\x91\xb6\xf9\x0d\x6e\xf4\x03\xa6\x2b\xce\xee\x38\xf2\x08\x47\x33\x83\x70\x2e\x85\x2e\x54\x0c\xcf\x99\xcc\xc0\xd3\x9e\x38\xea\xf4\xa4\x7d\x08\xca\xc4\x20\x5f\x94\x06\x0d\xff\x40\xc2\xc0\x92\x73\xac\x4e\x87\xf7\x98\xf3\xc0\xc3\xf3\x39\x3c\xc8\xcd\x02\x56\xee\xa2\xd2\xf1\x7f\x03\x00\x00\xff\xff\x65\x1a\xe4\x95\x1b\x1a\x00\x00")

func templatesApiGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesApiGotmpl,
		"templates/api.gotmpl",
	)
}

func templatesApiGotmpl() (*asset, error) {
	bytes, err := templatesApiGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/api.gotmpl", size: 6683, mode: os.FileMode(420), modTime: time.Unix(1517452009, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConfigGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\xc1\x6e\xe3\x36\x10\x3d\x8b\x5f\x31\xd5\x21\x96\x8a\x95\x1c\xec\x69\xb1\x80\x0f\xee\x66\x17\x0d\xe2\x83\x51\xbb\xdd\x33\x4d\x0d\x25\xc2\x14\x29\x90\x23\x07\x46\xe2\x7f\x2f\x48\x4b\x89\x6c\xa7\x3d\xac\x0c\x18\xd0\x90\x7c\x8f\xf3\xe6\x3d\x75\x5c\xec\x79\x8d\xf0\xf2\x02\xe5\x72\xfd\xb8\x1e\x5e\x4f\x27\xc6\x54\xdb\x59\x47\x90\xb1\x24\x95\x2d\xa5\x8c\x25\x69\x6d\xbb\x7d\x5d\x2a\x33\xe7\x1a\x05\x35\xb6\xe5\x7e\xbe\x57\xa6\xee\x94\x29\x0f\x9f\x53\x96\x33\x26\xac\xf1\xf1\x50\x85\x92\xf7\x9a\x96\x55\xe5\xd0\x7b\x58\x40\xfa\xf5\xcb\xfd\x97\xfb\xb8\x69\x3e\x87\x6f\xd6\x48\x55\x43\x85\x52\x19\xf4\x40\x0d\x82\x38\x97\x6c\x47\xca\x1a\x0f\xd2\xba\x58\x5e\xae\x1f\xc1\xa3\x3b\xa0\x2b\x19\x1d\x3b\x1c\x8f\x7a\x72\xbd\x20\x78\x61\xc9\xc8\xf1\xfe\x78\x72\xca\xd4\x2c\x79\xc0\x5d\x5f\xc3\xf4\xd9\x59\xab\x59\xf2\x68\x3c\x8a\xde\xe1\x9f\xdb\xed\x7a\x5a\x5f\xf6\xd4\x3c\x28\xcf\x77\x1a\xab\x69\x7d\xbb\xda\x7c\x43\x47\x3f\x94\xc6\x2b\x86\xed\x6a\xf3\x84\xc7\xc9\xc2\xdb\xca\x4f\xd4\xfa\xc9\xd8\x67\xf3\x06\x38\x60\xd9\x3d\x9a\xbf\xff\x5a\x8d\x1b\x4f\x8c\xc9\xde\x08\xc8\x04\xfc\x7e\x6e\x2d\x87\x9f\x8a\x9a\x87\xb3\x80\x3f\x34\xaf\x7d\x96\x87\x3e\x47\xa9\x43\x29\x4b\xab\xd0\x5b\xfa\x09\xd2\xef\x26\xc0\x43\x7c\x07\x6d\xeb\x5a\x99\x1a\xb8\xa9\xa0\xeb\x9c\x95\xd0\x22\x39\x25\x7c\x99\xe6\xe5\x1f\xd6\xea\x7f\xb8\xcb\xee\x44\x19\x95\xc9\xaf\x31\xf9\x59\xc9\x80\x3a\x8a\x4a\x16\xb4\xf2\x84\x06\xac\xf9\x04\x58\xd6\x25\xc4\x41\x82\x75\x70\x5f\xc6\x5f\x7c\x0f\xf8\x2c\x49\x86\x5b\x67\x97\xe3\xcf\xcb\x4d\x6c\x76\x20\x1f\xab\xd7\xf4\x6a\x18\x4b\xd1\x10\x75\xe1\x12\x1b\x74\x07\x25\x10\xac\xd1\x47\x08\xc3\xba\xea\x62\x3a\xc7\x1b\x34\xd2\xbe\x10\xe8\xa8\x90\x4a\x63\x40\x5b\x73\x6a\x42\x3f\xdb\xd5\x06\xc2\x38\x21\x2c\x40\xef\xb1\x82\xe7\x06\x4d\x74\x59\xd0\x2e\xa0\x6d\x86\x7e\x2e\xee\x3d\xb1\xc1\x87\x6c\x7b\x3c\x7e\x48\xf6\x84\xc7\x5f\xe1\x1a\x8c\x75\x43\x55\x9d\x0d\x55\x3c\xa3\xd6\xc5\x3e\x58\x2c\x10\x0e\x36\x03\xde\x93\x6d\x39\x29\x01\xf3\xf2\x7d\x07\x38\xf4\xb6\x77\x02\xfd\x40\x36\x51\xf1\xc6\xa9\xff\xc9\xc8\x7b\x6a\xae\xb8\x9a\x18\x54\xae\xf5\xff\x31\x4c\x73\x75\xab\x5c\x08\x44\xd1\x3b\x7d\x9e\x38\xc1\x5b\x42\xa2\x5c\x64\xe1\xc0\xb5\xaa\x38\x21\xd8\x40\xf8\x19\xe2\x09\xff\xa1\x6a\xc3\xd1\x3c\xc4\x6a\x3e\x87\x35\x77\x1e\xa1\x0b\xff\x7e\xf8\xc0\xf4\x8e\x87\x0f\x0c\x48\x67\x5b\x10\xb6\x6d\xb9\xa9\xb4\x32\x08\x32\x24\xad\xbc\x0d\x63\xc4\xc8\x72\x40\xe7\xac\x9b\xe6\x70\x58\x60\x2c\x51\x12\x7e\xbb\x74\x23\xdc\xdd\x41\x76\xe1\x18\x58\x2c\x20\x4d\xe1\xf5\x15\xa6\xc3\x3d\x57\x63\xbc\x13\x87\xd4\x3b\x03\xb2\xa5\xf2\x7b\xe0\x92\x59\x3a\x2b\x8a\x0b\x1b\xcf\x62\xb0\x87\xea\x68\xb7\x19\xb4\xbd\x27\xd8\x21\xf8\x0e\x85\x92\x6a\x34\xd9\xac\x28\x2e\x22\xb5\x90\x5c\x7b\x9c\xa5\x39\x4b\x4e\x6c\xa4\x33\x4a\xb3\x53\xd0\xea\xa0\xda\xaf\x20\x69\x51\x5b\xf6\x6f\x00\x00\x00\xff\xff\x53\x94\x48\x4d\x17\x06\x00\x00")

func templatesConfigGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesConfigGotmpl,
		"templates/config.gotmpl",
	)
}

func templatesConfigGotmpl() (*asset, error) {
	bytes, err := templatesConfigGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/config.gotmpl", size: 1559, mode: os.FileMode(420), modTime: time.Unix(1517451889, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesParameterGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x6f\x73\xdb\x36\xd2\x7f\x6d\x7d\x8a\x2d\x9f\x36\x43\xe6\x91\xa9\x3c\x7d\x3a\x7d\xa1\xab\x3b\xd3\x24\x4e\xe3\x6b\x9b\xe4\xe2\x26\x7d\x91\xc9\xb4\xb0\x08\x49\x68\x48\x80\x06\x20\xdb\x3a\x0e\xbf\xfb\x0d\xfe\x10\x04\x28\x92\x96\x1c\xf7\xda\xeb\x35\x6f\x22\x62\x81\xc5\xee\x0f\x8b\xc5\xee\x02\xae\x2a\xc8\xf0\x92\x50\x0c\x91\xc8\xc9\x02\x97\x88\xa3\xe2\x0a\xe5\x24\x43\x92\xf1\xa8\xae\x27\x55\x05\x64\x09\x8c\x43\xfa\x03\xa1\x67\x12\x17\x02\xd2\x1f\xd0\x8d\xf9\x65\xe8\x0b\x54\xe0\x9c\xfc\x13\x43\xfa\x02\x15\x18\xea\xfa\x5c\x7d\xcc\x4f\x80\x50\xf9\xe5\x17\x71\x8e\x69\x6c\xb8\x20\x9a\x41\x4c\x99\x84\xf4\x4c\x7c\xc3\x39\xda\x26\xf6\xf3\x39\x12\x4f\x89\x58\x70\x52\x10\xaa\x26\x4e\x5c\xb7\x33\x2a\x31\x5f\xa2\x05\x6e\x9b\xce\x25\xc7\xa8\x48\xd4\xcf\x17\x9b\x3c\x47\x17\xb9\x9a\xf3\x61\x55\x01\xa6\x19\xd4\x75\x55\x41\xfa\x16\xe5\x1b\x7c\x7a\x53\x72\x2c\x04\x61\x14\xea\x3a\x49\x26\xae\x87\x55\xaa\xd5\xa8\xae\x27\x64\x09\x98\x73\x25\xb5\x55\x1f\x3b\xb2\x92\x3e\x7d\x85\xe4\x1a\xea\x7a\x0a\x55\x05\x25\x27\x54\x2e\x21\xfa\xec\x32\x82\xf4\x7b\xb6\x40\xd2\xcc\xa1\x89\xbd\x68\x68\x8a\x3f\x5f\xf2\x37\x3d\xdd\x27\x27\x40\x49\x0e\xd5\x04\x80\x63\xb9\xe1\x54\xb5\x4e\xea\x1e\x51\x3d\xc8\xfb\x44\xb5\xe4\x7b\x12\xd5\xf1\x3b\x5c\xd0\x37\x94\x5c\x6e\xf0\x98\xac\x5e\x8f\xc3\xc4\xfd\xbd\x2d\xe8\x40\x24\x4e\xe9\xa6\x18\x80\x40\x91\xfe\xa3\x74\x37\xf6\x6b\x35\x3a\x04\x88\xf6\x57\xe3\x67\x2c\x08\x84\xd1\x57\x4a\x62\x49\xae\x70\x64\xba\xd8\x2d\xf9\x3d\xa6\x2b\xb9\x1e\xdc\x92\x86\x6c\x3d\x4a\x03\xa0\x87\x65\x55\xe1\x5c\xe0\xba\x8e\xa2\xaa\xc2\x34\xdb\x13\xdc\x10\x8e\xd8\xc7\xa3\x0b\x87\x9e\x6d\x67\x48\xe2\x23\x68\xa8\x4f\x36\x42\xb2\xe2\x19\xe3\x05\x92\x12\x73\xa8\xeb\xf4\x5c\x72\x42\x57\x71\xdb\xd9\xec\xb6\x56\xe9\x5b\xb1\xd5\x2a\x35\x60\xa1\x9b\x51\xb0\x1a\xf2\x9f\x0b\xac\x56\xe9\x83\xc0\x7a\xa5\xf9\xd2\x7e\xa8\x2c\xf1\xcf\x03\xd4\x2f\x55\xd5\x6a\xfc\xcb\x61\x56\x45\x28\x29\x36\xc5\xe0\x06\x54\x44\x23\x0d\xbe\x84\xf4\xfc\x1a\xad\x56\x98\xff\xb8\x2d\x31\x44\x84\x4a\xbc\xc2\x5c\xed\xe7\x33\x2a\x9d\x38\xf7\x0d\xeb\xd8\xbc\xc4\xcc\x9b\x0b\x05\xdf\x32\x67\xa8\x15\xe3\xcb\x2f\xe2\x3e\x8c\xc7\x57\x25\x69\x76\xa8\xc1\x44\x7f\x9d\xde\x2c\xf2\x8d\x20\x57\xd8\x35\x1f\xba\x6d\x47\x00\x36\xc4\xff\x3a\x80\x1b\x4c\x3a\x00\x37\xcd\x87\x01\xbc\xc9\x25\x29\x73\xfc\x72\x39\x80\xb1\xa3\xdf\x1f\x70\x1a\x89\x8f\x01\xc0\x93\xf9\x20\x65\xd5\xa1\x3c\x1e\x65\xdc\xa7\x65\x7c\x5c\x10\x72\x08\x2c\x07\x9f\x0c\xad\xec\xff\x73\x15\x35\xb8\xec\x0d\x65\xf3\xbf\x8b\x53\x4a\xce\x4a\xcc\xe5\xb6\x93\x12\x79\x01\xde\x99\x70\x11\x8c\x81\x56\xe2\xa2\xcc\x91\x1c\x0a\x72\x52\xd3\x2b\x8c\x12\xfb\xf4\xeb\x5b\x4c\x43\xd7\x16\xab\x57\x50\xe9\x1c\x55\x95\x5b\xa3\xba\x8e\x4c\x83\xdd\xb7\xa6\xbf\x6e\xed\x43\xd7\x41\x38\x85\xa5\xee\x29\xc6\x91\xea\x91\x5b\x2f\x7f\x57\xf1\xbe\x2c\xd2\x57\xfc\xb8\x13\x16\x36\x68\x5f\x10\x9a\x95\x0d\x54\x7a\x7c\x34\x18\x41\xb6\x73\xa8\x51\xc6\x2b\x4d\xae\x10\x57\x36\x70\x85\x38\x55\xb9\x4c\xfa\x64\x4d\xf2\xac\x27\x92\x7d\xad\x23\xd9\x6f\x99\xf6\x6b\x75\x3d\x59\x32\x0e\x6d\x8e\x6b\x46\x3d\x47\xe2\xad\x5b\x40\xd1\xb4\x3e\x61\xf4\x0a\x73\xb5\x42\xb6\xa1\x6f\xe9\x14\xf3\x33\x9a\xe1\x9b\xb7\xc8\x7e\x5a\x4f\xf9\x73\x68\xab\xb7\xca\xf9\x56\xad\x3e\x47\x74\x85\xf7\xea\xfe\x44\xaf\x58\x57\x91\x66\x91\x14\xea\x9a\xda\xa8\x32\xc8\x65\x7e\x02\xe2\x1a\xad\xd2\xf3\x32\x27\xf2\xf1\xd6\xa8\x16\xef\x25\xf0\xae\xff\x68\x70\xcb\x73\xbc\x50\x50\x1a\x6e\x2a\xac\x31\xd2\xf4\x99\x4d\xb3\xa4\x66\x30\x58\xc1\x8f\x0d\x8c\x4e\x0f\x65\x81\xdd\x55\x71\xc4\x5b\x45\x9d\x36\xbb\xcb\x43\xc4\xb1\x31\xa7\xe9\xed\xea\x2a\x1d\xec\x3e\xf5\xf6\x8c\xbf\x6b\x18\x17\xe9\x19\xd5\xfb\x40\x59\x5b\xdc\xce\x36\x98\xf6\x19\x72\xe0\x77\xa3\x76\x98\xb3\xda\x68\x3f\x1b\x52\x22\x06\xf8\xb5\xb0\xed\xda\xee\x1d\xe0\xb3\x9e\x23\x7d\x85\xb8\xc0\x7f\x5e\xd4\xf6\x47\xc6\xda\xd4\xed\x30\x18\x7e\xd6\xb5\xd9\xa9\xcc\x47\x77\x67\x0c\x1d\x43\xe1\xfe\xd8\xc7\xef\x9d\x00\x2a\x4b\x4c\xb3\xbd\x16\xea\xf5\xbe\x58\xf9\x3e\xba\x44\x8b\x0f\xc8\x38\xac\xf4\x95\xfd\xad\x54\x9a\xcd\xe0\xc7\x35\x11\xb0\x24\x39\x86\x6b\x24\x60\x85\x29\xe6\x48\xe2\x0c\x2e\xb6\x20\xd7\x58\xfb\x9c\x15\xe6\x20\x19\xcb\x53\xd5\xff\x34\x23\x92\xd0\x15\x48\x37\xae\x20\xab\xb5\x84\x92\xb3\x2b\x0c\xcb\x8d\xd4\xac\xd6\x98\xc2\x96\x6d\x80\xe3\x63\xbe\xa1\x01\xa7\x66\x0a\x58\xb0\xa2\x40\x34\x9b\x4c\x48\x51\x32\x2e\x21\x9e\x00\x44\x14\xcb\xd9\x5a\xca\x32\x52\x78\x47\x2b\x22\xd7\x9b\x8b\x74\xc1\x8a\xd9\x8a\xd0\xe3\x15\xa3\x64\xa1\x7e\x45\x5d\x22\x3b\x66\x25\xa6\xa8\x24\x33\x63\xa8\x23\x1d\x9a\x83\x7b\xa4\x0b\xdf\x50\x49\x8a\x3d\x7a\xcc\x0a\x92\x65\x39\xbe\x46\x7c\xac\xb3\xd2\xbc\x4b\xe6\x58\x8a\xcd\x23\xad\x95\x05\x66\x86\x4a\xa2\xb5\x16\x92\x2f\x0b\x39\xc8\x4c\x53\x23\x6b\x8e\xe6\x1c\x4a\x9f\xe2\x25\xda\xe4\xf2\x4c\x03\x29\xcc\xf1\x16\x6c\xc4\xc6\x1a\x3d\xc3\xb6\x63\x3f\xfd\x80\xb7\x53\xf8\xf4\x4a\x59\x90\xda\x25\x69\xc0\x44\x51\xd5\xf6\xe9\xf0\xb3\xdd\x3b\x5c\x13\x6d\x50\xa7\x34\x2b\x99\xc9\x45\x4a\x24\x16\x28\xa8\x63\x02\xbe\xc1\x8b\x8d\xc4\x42\x9b\xc4\x82\x71\x0c\x39\x5b\x91\x05\xb0\xa5\x6e\xe1\x58\x6d\xae\x4c\xf1\xe1\x6c\x23\xb1\x62\xad\xb9\xa5\x93\xe5\x86\x2e\xc6\x99\xc7\x6b\x44\xb3\x1c\x73\x50\x5d\xe3\x85\xbc\x81\x87\x2b\x42\xd5\x09\x22\xf1\x8d\x74\xf1\x36\x47\xba\x02\x3a\x85\xd2\xfc\x7c\xd8\xcb\xcc\xf4\x73\xca\x25\xf0\x10\x95\x24\x7d\x8d\x45\xc9\xa8\xc0\x09\x28\xce\xcf\xcd\x7c\xcf\x94\x64\xd5\xe4\xc8\xba\x4a\x2d\xe8\xce\xf4\x89\xea\x71\xd4\x15\x62\x36\x6b\xf7\x83\x15\x67\xc9\x59\x01\x1c\x5f\x6e\xb0\x90\x93\xa3\x23\xdb\x3a\x3f\x81\x07\x63\x72\xd6\x93\xa3\x23\xeb\xfc\xcd\x88\x94\x63\x94\xbd\x36\x6c\x94\x30\xc9\xe4\xe8\xa8\xeb\xdf\x8f\xf4\x98\x97\x17\xbf\xaa\x61\x98\xf3\x34\x7e\x68\xfd\xfc\x13\x56\x94\x4c\x10\x89\x4f\xd5\xb7\x1a\x7b\x54\x72\x76\x91\xe3\x42\x75\x55\x48\xbc\x32\x9f\x9a\xc9\xd1\x8f\x44\xe6\x78\x0e\x10\xbd\xa1\x25\x67\x0b\x2c\x84\xce\x24\x4e\xa9\x24\x72\x9b\x46\x53\xdd\xe9\x5c\x22\xb9\x11\x73\x20\x54\xc6\x66\xda\xf4\x09\xcb\x70\x9c\x24\x86\xfe\x14\x4b\x44\xf2\x39\x58\x9a\x9e\x39\x36\x34\xa5\xdd\xd1\x42\xde\xa4\x3f\x71\x22\x31\x4f\x9f\x63\x94\x61\x1e\x27\xe9\x39\x96\x71\xa4\x11\xa6\xf2\x58\x9d\x26\x2a\xd6\x46\x65\x99\x13\x73\xe4\xcc\xac\xd4\xff\xfb\xab\x60\x34\x4a\x1a\x36\x7f\x3f\x7f\xf9\x22\xb6\xa4\xd4\xc8\x35\x05\xfb\xad\x3b\x99\xa5\x9c\x98\x99\x5b\x0b\x57\x04\x51\x2a\x08\xac\xa9\x29\x64\x87\x0c\xcb\xa6\x2e\x8a\x9f\xb8\x26\x72\xb1\x06\x35\x58\xeb\xac\xb1\x5f\x20\x81\x41\xf9\x3a\x2b\xc1\x0b\x66\x15\x99\x37\x62\x7e\x73\xc1\xb8\xfc\x89\xc8\xb5\xe9\x10\xbb\xf1\x8a\x67\x66\xf6\xfc\x3c\xd0\xc9\xf5\x98\x9a\xc9\x1e\xb3\x6c\x9b\x18\x35\xea\x89\x71\xf8\x23\x56\x04\x0b\x46\x25\x22\x54\x00\xca\x73\xbd\x1f\x2f\xd8\x86\x66\xce\x32\x19\xd7\x8d\x55\x05\xeb\x4d\x81\x68\xb0\xb1\xd5\x61\xa8\x21\x57\x73\xc8\x6d\x49\x16\x28\xcf\xf5\x19\x22\x30\x20\x8e\x81\x5d\x28\xd6\x38\x33\x06\x8e\x8c\xe6\xd6\x40\x27\xb3\x99\x1a\x66\x3d\xe1\x5c\xcf\x87\x25\xe6\x42\x9f\x58\x76\x8a\x89\x54\xd1\xc2\x98\xf8\x42\xf2\xcd\x42\x42\x15\x3a\xc7\x76\xb3\x3d\xb4\x4b\xf5\x14\xab\x7c\xb8\xb4\x31\x89\x9a\x62\xa7\x25\xa8\xed\x29\x29\x09\xc7\xd6\x75\x36\x5f\x73\x90\x7c\x83\xbb\x7d\x6d\x61\xc4\x74\xb5\x1f\xf3\x26\x32\xee\x96\x4f\xa0\xae\xbf\x82\xa0\xc2\xdf\x12\x76\x18\x9b\x92\x96\x65\x6c\x3e\x7a\x18\xbb\x5e\x5f\x77\x18\x3b\xc2\x0e\x63\x57\xe0\xb0\xbc\xed\x37\xbc\x5c\xce\xcd\xed\x98\xdf\xa1\x47\x5f\x53\xf5\x75\x1a\x83\xf9\xb6\x63\x3d\x72\x8f\x46\xc1\x50\x42\xc3\xa1\x1e\xb9\x3b\xd4\x56\x50\xcd\x40\xfb\x31\xb7\xe1\x4d\x43\xe9\x91\xd4\xdd\x7e\x19\x41\xf5\xa7\x93\xb3\x21\xf6\x88\xe9\x8f\x23\x34\x18\xd7\x12\xbb\xe3\x3a\x17\x6e\x00\xa6\xa1\xdf\x6c\xbc\x18\x79\x02\x70\x66\x95\xf1\x5a\xbb\x03\x7a\x92\xb7\x09\x40\xdb\x0a\xa6\xd9\xf0\xe9\xe9\xdc\xe5\xf7\x1c\x09\x1b\x45\x18\x4e\xf6\x63\xbe\x5b\xbf\x69\xbb\xf9\x01\xc5\xc3\x99\x4b\x00\x75\x79\xe9\x7c\xb1\xc6\x05\xb2\xa1\x48\xbb\x61\xcf\x9e\xda\x70\xe2\xde\xee\xcd\x9e\x91\x1c\xeb\x2d\x7e\xc8\x55\x9a\xcd\x3d\xfc\x3a\x40\x7f\xbc\xd2\x95\xd4\xa8\x95\x9e\x89\xc7\x48\x60\xc5\x22\x9c\xa5\xd3\xa9\x11\x64\x64\x72\x2f\x20\x33\x95\x2e\xe3\xa5\xbd\xd3\x1b\x2e\x98\x5c\x83\x4a\xbf\x85\x16\xa4\x09\x60\x05\xa0\x26\x4e\x98\x02\x91\x80\x84\xd8\x14\x3a\xb2\x42\x52\x05\xd7\x65\x8e\x6f\x54\x98\x4e\x57\x02\x88\xfa\x2a\x30\x95\x80\xc0\x56\x50\x94\xbc\xb1\x89\x26\xd3\xd7\x78\x45\x84\xe4\xdb\xc4\xe4\x79\xea\x88\x36\x30\x2b\x51\x94\xdb\x17\x9a\x01\xe8\xb0\x4f\xa8\xc9\xae\x49\x9e\xc3\x46\x60\xe5\x73\x91\x4e\x00\x0a\x2c\xd7\x2c\x03\xe5\xf6\x85\x09\xd6\x74\x86\xf8\x1a\x2f\x30\xb9\xc2\xbc\x01\x74\x2c\xda\x4a\xa0\x13\xb4\x74\x22\x28\x23\x9c\xca\x4b\xaf\x10\x57\x07\x1c\xbc\x7b\xaf\xdb\xda\xda\x83\x73\xf6\x36\x05\xd6\x75\x13\xa3\xe4\x0b\x7c\x6d\xac\x5f\x78\xf5\xc9\x49\x3b\xf4\x39\x12\xff\xd8\x60\xbe\x75\x2c\x2e\xf5\x68\x1b\xea\x9b\x1c\x4b\x28\xa9\x9a\x53\x2b\x7d\xf3\xfa\xfb\x54\x0f\x89\x93\x21\x8e\x6a\x46\xc7\xb0\xad\x1e\xfa\x5c\x74\x92\x6e\x7c\x2c\xe2\x52\x0d\x88\xff\xff\x73\xf8\xea\x2b\xf8\xfc\x51\xb7\xf4\xe7\x87\x70\xfa\xfc\x3c\xe5\xfc\x05\x93\x6e\xb0\xcd\xd9\x9b\x7f\x5e\x9d\xb0\x69\xaa\x5d\xbd\x61\x48\x12\x2d\xc0\x6e\xcd\x71\x9c\xeb\xe4\xa8\x0e\x75\xd6\x68\x39\xc5\x27\x00\xcb\xec\x36\x34\xd5\xb0\xc4\x4f\x28\xba\x90\x76\xcf\xf3\xd0\xe1\x04\x75\x4f\x53\x08\x6d\x97\x53\xad\x66\xaf\xe1\x4d\xe1\x72\xfd\x61\x80\xf2\xb3\x12\xf8\x52\xa4\xdf\x62\xf9\xf2\x3b\xff\x55\x80\x57\x2b\xb1\xb5\xaa\x8e\x95\xa7\x6a\xb7\xf6\x38\xbe\xf8\x70\x21\x06\x8b\xc0\xa0\x37\x80\xab\x1d\x70\x2c\x74\xf9\xa7\x2d\x92\xb4\x95\xa5\x33\xa1\x04\x6f\x80\xe0\xfd\x9e\x6e\x7e\x02\xef\xde\x0b\x5d\x7c\xae\xd4\xb2\xe8\xee\x81\xd6\xf5\x9d\xd5\xee\x9f\x72\xaa\x4f\xc2\xfb\x52\xd1\xe4\x04\x8d\x92\x77\x93\xd3\x37\x47\xc3\xef\x9d\xde\x66\x4f\x10\x65\x54\x45\xb4\xa6\xf1\x3b\xbc\x0d\x80\x79\x7f\xbf\x9a\x38\x9f\xa1\x6d\xd9\xb6\x35\xa7\x9c\xb1\xef\x9d\x07\x42\xfd\xcf\x86\x8c\xb8\xd3\xbe\xbd\xae\x26\x51\x4c\x07\xec\xfa\x76\xd9\x55\x9a\xf8\x02\x5f\xc7\x5f\x3c\x7a\x34\x85\x48\x39\x6e\x42\x57\xa6\x22\xf4\xd9\x25\x2c\x11\xc9\x55\x88\xfc\xd9\x55\xb4\x53\x16\x8c\x43\x39\x93\xa6\x72\x99\x68\x38\x0c\x12\x66\xe2\xbe\xb5\xeb\xb7\xde\xd6\xb3\x28\xa5\xaa\xa7\x48\xa2\x79\x2f\x24\x53\x30\xa0\xf4\x53\x0d\xad\xee\x2c\x4b\x5d\x2f\xb3\xa1\xed\x99\x8d\xbb\x8f\x65\x76\xaf\xfe\xe3\x2e\x72\x7c\x8c\x51\x76\xfc\x70\xd7\x52\xff\xf2\xb8\xfd\x9b\x58\x85\x69\x9d\x8d\xfc\x97\x05\x85\xc1\x91\x85\xe8\x31\xcb\xac\xbd\xb4\x19\x03\x59\xba\xcd\xfc\x1c\xe9\x1e\xbe\x63\x4e\xbc\x0b\xb5\x6e\xa4\x6d\x73\xdd\x7d\x5d\x86\xef\x0f\xd5\x34\xc1\x96\xf7\xc4\xdc\x09\xfa\xbd\xa6\xd3\x9b\x92\x71\xa9\x6b\x03\x17\x2c\xdb\x06\x37\x58\x3f\xb0\x0c\xe7\xa2\x2d\xb8\xa7\x6f\x68\x81\xb8\x58\xa3\xbc\xaa\x54\x54\x4a\xca\x86\xd6\xdc\x85\xec\x0c\xd9\xb9\xd4\x3d\xcf\xc9\xa2\xcd\x1e\xe3\xae\x0a\x53\x53\x37\x55\x31\xb3\xca\x06\x78\x8f\x4b\x87\xde\xaa\x86\xeb\x76\x72\x02\x84\xa5\xa7\x2f\x9f\xb9\xb0\x4f\xb7\x36\x2e\xbf\x19\xb5\xf7\x2b\xc9\x64\x62\x22\x44\xeb\xcf\x8d\xdc\x83\x66\xd3\xe2\xaf\x42\x7c\x85\x68\xe7\x5a\x18\x3a\xb1\xeb\x63\x42\x33\x5d\xfa\x7a\xa0\x3a\x77\x4c\xf4\x60\x55\x07\x4f\x3a\x5f\xed\x5b\xcf\xb2\x31\x34\x2c\x1c\xf6\x94\x0b\x6a\xf2\xe3\x07\xad\x0e\xce\x4d\x51\xf4\x23\x65\x98\x42\x14\xd9\x03\x77\x00\x9f\xce\x6a\xf9\x5b\xb9\x7b\x3e\xf7\x1e\x04\xcd\x15\x98\xf9\x8c\x7b\xd2\x67\x3f\x91\x0f\xae\xc5\x73\x82\x04\xce\xbc\xdb\x50\x93\xc8\xbe\xbc\xf8\x15\x2f\x64\x92\x98\xcc\x0e\x7e\x36\xef\x5e\xc3\x0b\xfd\x9d\x5c\xd3\xbf\xa8\xdf\xd3\x29\x34\x86\x10\x3a\xd7\xf1\x79\x52\x9b\x51\xe3\x78\xc4\x51\x8e\x3a\x4b\xf3\xef\x82\x63\xf4\x61\xd2\x64\x54\x3d\xeb\x10\xfc\xb0\x47\xcd\x3e\xe0\x3a\x82\x43\xd7\xb5\xec\xc2\xdb\x6a\xae\x36\xd4\x9e\xba\x8d\x68\xb6\x6b\x4b\x1a\xdd\x1c\x53\xd5\x31\x51\x9b\xf0\x91\xe3\x73\x88\xf7\x3e\xb0\x34\x53\xd7\x0f\x9c\x10\x17\xc6\xd9\x1b\xe1\xba\x57\x63\x3d\x55\x5f\xdf\xe6\xff\x3d\x2e\xa2\xf6\x65\xda\xbd\xbb\x73\xbf\x7d\x24\xbf\x76\x40\x86\x97\xf6\xee\x32\xa7\x7d\x3b\x63\xfc\x08\xc7\x22\x4d\xd3\xe6\xb4\xb6\x83\x28\xc9\x27\xf5\x64\x52\x55\xf0\xe9\x22\x47\x42\x68\xc0\xe7\x27\x10\x77\x16\x21\xb1\x8f\x7f\x76\xb2\xf2\x36\x27\xd7\xc6\x17\x9c\xf1\x41\xb9\x2e\xf8\x23\x16\xff\xcd\x56\xff\xdb\xab\xf1\x9a\x92\x27\x6c\x5b\x4e\x1a\xcc\x47\xd1\xb5\xca\x10\x5c\xc2\x3b\x85\x35\x12\xdf\xe1\x2d\x5c\x30\x96\xbb\x78\x07\x06\xaa\x63\x55\x10\xc4\x34\x45\x47\x97\x62\x27\x81\xe9\x90\x25\x7c\x62\x99\xf7\xad\xcd\x9d\x8e\xd3\xc0\x08\x74\x29\x0c\x5d\x83\xd1\xc4\x33\x09\xa3\x63\x60\x16\xe8\x5a\x25\x4a\x86\xf0\xce\xef\x74\xfc\x7f\xef\x5b\xbe\xfb\x28\x66\x88\xdf\xe4\x39\xbb\x3e\x2d\x4a\xb9\xd5\xf5\x9c\xd0\x7d\xb8\x17\x72\xcd\x20\xfb\xb2\x6d\xff\x3f\xaf\xe0\xe8\xba\x3f\xe6\xf4\x0a\x50\x3d\x91\x77\x0c\x5d\xc9\xc1\x38\x42\x23\x74\x23\x4e\x32\x24\xbf\x86\xe9\x04\xa2\x08\x2a\x98\xcd\x00\x2b\x7a\x53\xfa\x2c\x91\x30\xb7\x63\x4c\xae\x31\x87\xf6\x39\xa1\xf0\x0f\xc4\xa0\x98\x6e\x1f\xc0\x85\x5e\xa0\xae\x9b\x0e\x41\x09\xab\x93\xe6\xb7\x71\x4f\x1b\x17\x31\x61\x52\x5c\x73\xbb\xe8\xfc\xa0\xbf\x81\x82\x22\x73\xdc\x54\x96\xc7\x0b\xf9\xbb\xef\x26\x93\xd0\x67\xf7\xff\xfd\x49\x8f\x1f\xbe\xbd\x98\x6f\xf6\xbb\xf3\xcc\xb0\x5b\xb9\xf7\x9d\x75\x5f\x9e\x6e\x45\xef\x06\x95\xce\x81\xed\x7a\x75\x77\x6d\xd2\x3e\x27\xd3\x4b\x1a\x3e\x39\xf3\x1f\x9b\x29\xeb\xbb\xd3\xa3\xa8\xbd\xff\x76\x28\x20\xba\xa5\x36\x76\xdf\xaa\x30\x86\xfa\xd0\x01\xa7\x55\xdb\x49\xd7\x77\x9d\x6a\x08\xc1\xce\xb3\xb1\x40\xc0\xe0\xfd\xaa\x2f\xe7\x1f\x1a\xa1\x3b\x5c\x32\x8d\xdc\x28\x35\xdf\x0d\xe8\xe1\xd5\x4e\xac\xe1\x4c\xf5\x9f\x93\xb4\xd2\x26\x26\xaf\x31\x8f\xba\xef\xb8\x9e\x1c\x5d\xef\xd8\xb3\x75\x34\xfe\x83\xd8\xdb\xea\x9e\x8d\x4b\xee\x2d\x0c\x8c\x25\xf9\x3d\x0e\xb7\x11\x24\x08\x1b\x8c\x9a\x7e\x52\xf0\xc7\x3b\xb8\x3b\xb1\xdd\x6f\x7e\x40\x3b\xe7\x83\x2f\x7b\x6e\x62\xa3\x62\x93\x4b\x12\x19\x3f\xb7\xc7\x6b\xe2\xb9\x3b\xc0\xc3\x54\xf9\xf2\xaa\x3f\x4e\xde\x23\x2c\x18\x1a\xda\x1f\x2a\xc0\x31\xd8\x60\x61\xcf\x07\x8e\x43\x6f\x97\x07\xa6\xed\x79\x45\xda\xf3\x4e\x79\x67\x2b\xe8\x42\xcd\xed\xe1\x49\x0b\xc4\x5e\xa2\x07\xe9\xc9\x6f\x66\x18\x41\x58\x32\x70\x25\x9f\xe1\xe5\xdb\xe6\x8d\x5e\xff\x43\x70\xef\x3c\xdf\x0f\xc3\xbb\x61\xf1\xe0\x81\x1e\xd2\xc8\xe3\x1b\xd2\xa0\x73\x6b\x3a\x07\x05\x9e\x8f\x5d\x07\x4a\xf2\x6e\xda\xe6\xe3\x3a\xfa\x86\xdd\xf5\x1a\x74\xc6\xb7\x3f\xc4\x0d\xee\x79\x75\x59\xee\xf7\x75\xc5\xbb\xbe\x38\xfc\xcb\x10\x15\x75\x75\xff\x80\xa2\x5f\xf2\xbb\x78\xec\x3d\xf4\xb9\x25\x9f\xda\xe3\x69\x75\xef\x89\xd3\xfb\x47\xcd\xf6\xd7\x6c\x06\x57\xa4\x98\xc3\x52\x9e\xac\xd8\xe4\x5f\x01\x00\x00\xff\xff\xa4\xd7\x9b\x7a\x55\x41\x00\x00")

func templatesParameterGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesParameterGotmpl,
		"templates/parameter.gotmpl",
	)
}

func templatesParameterGotmpl() (*asset, error) {
	bytes, err := templatesParameterGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/parameter.gotmpl", size: 16725, mode: os.FileMode(420), modTime: time.Unix(1517451889, 0)}
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
	"config.yaml": configYaml,
	"templates/api.gotmpl": templatesApiGotmpl,
	"templates/config.gotmpl": templatesConfigGotmpl,
	"templates/parameter.gotmpl": templatesParameterGotmpl,
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
	"config.yaml": &bintree{configYaml, map[string]*bintree{}},
	"templates": &bintree{nil, map[string]*bintree{
		"api.gotmpl": &bintree{templatesApiGotmpl, map[string]*bintree{}},
		"config.gotmpl": &bintree{templatesConfigGotmpl, map[string]*bintree{}},
		"parameter.gotmpl": &bintree{templatesParameterGotmpl, map[string]*bintree{}},
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

