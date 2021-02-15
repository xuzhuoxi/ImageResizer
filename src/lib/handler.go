//
//Created by xuzhuoxi
//on 2019-04-27.
//@author xuzhuoxi
//
package lib

import (
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"github.com/xuzhuoxi/infra-go/imagex/resizex"
	"image"
)

var defaultHandler = ImageResizeHandler{}

func LoadImage(fullPath string) (img image.Image, imgFormat string, err error) {
	return defaultHandler.LoadImage(fullPath)
}

func ResizeImage(source image.Image, width, height uint) (img image.Image, err error) {
	return defaultHandler.ResizeImage(source, width, height)
}

func SaveImage(img image.Image, fullPath string, format formatx.ImageFormat, options interface{}) error {
	return defaultHandler.SaveImage(img, fullPath, format, options)
}

type IImageResizeHandler interface {
	LoadImage(fullPath string) (img image.Image, err error)
	ResizeImage(source image.Image, width, height uint) (img image.Image, err error)
	SaveImage(img image.Image, fullPath string, format formatx.ImageFormat, options interface{}) error
}

type ImageResizeHandler struct{}

func (h *ImageResizeHandler) LoadImage(fullPath string) (img image.Image, imgFormat string, err error) {
	return imagex.LoadImage(fullPath, "")
}

func (h *ImageResizeHandler) ResizeImage(source image.Image, width, height uint) (img image.Image, err error) {
	return resizex.ResizeImage(source, width, height)
}

func (h *ImageResizeHandler) SaveImage(img image.Image, fullPath string, format formatx.ImageFormat, options interface{}) error {
	return imagex.SaveImage(img, fullPath, format, options)
}
