//
//Created by xuzhuoxi
//on 2019-04-26.
//@author xuzhuoxi
//
package lib

import (
	"strings"
	"io"
	"image"
	"errors"
)

type ImageFormat string

func (f ImageFormat) Encode(w io.Writer, m image.Image, options interface{}) error {
	if fm, ok := getFormat(string(f)); ok {
		return fm.encode(w, m, options)
	}
	return errors.New("No RegisterEncode:" + string(f))
}

type format struct {
	name   string
	encode func(w io.Writer, m image.Image, options interface{}) error
}

const (
	PNG  ImageFormat = "png"
	JPEG             = "jpeg"
	JPG              = "jpg"
	JPS              = "jps"
)

var formats []format

func getFormat(name string) (fm format, ok bool) {
	for _, fm := range formats {
		if name == fm.name {
			return fm, true
		}
	}
	return format{}, false
}

//--------------------

func CheckFormat(format string) bool {
	format = strings.ToLower(format)
	_, ok := getFormat(format)
	return ok
}

func RegisterEncode(name string, encodeFunc func(w io.Writer, m image.Image, options interface{}) error) {
	formats = append(formats, format{name, encodeFunc})
}
