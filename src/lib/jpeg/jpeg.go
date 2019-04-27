//
//Created by xuzhuoxi
//on 2019-04-27.
//@author xuzhuoxi
//
package jpeg

import (
	"io"
	"image"
	"image/jpeg"
	"github.com/xuzhuoxi/IconGen/src/lib"
)

func init() {
	image.RegisterFormat("jpeg", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("jpg", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("jps", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)
	lib.RegisterEncode(string(lib.JPEG), EncodeJPEG)
	lib.RegisterEncode(string(lib.JPG), EncodeJPEG)
	lib.RegisterEncode(string(lib.JPS), EncodeJPEG)
}

func EncodeJPEG(w io.Writer, m image.Image, options interface{}) error {
	if nil == options {
		return jpeg.Encode(w, m, nil)
	} else {
		return jpeg.Encode(w, m, options.(*jpeg.Options))
	}
}
