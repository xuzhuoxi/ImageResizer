//
//Created by xuzhuoxi
//on 2019-04-27.
//@author xuzhuoxi
//
package png

import (
	"github.com/xuzhuoxi/IconGen/src/lib"
	"image/png"
	"io"
	"image"
)

func init() {
	lib.RegisterEncode(string(lib.PNG), EncodePNG)
}

func EncodePNG(w io.Writer, m image.Image, _ interface{}) error {
	return png.Encode(w, m)
}
