//
//Created on 2019-04-27.
//@author xuzhuoxi
//
package lib

import (
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/jpegx"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/pngx"
	"github.com/xuzhuoxi/infra-go/imagex/resizex"
	"image"
	"os"
)

type ImageInfo struct {
	Path   string
	Image  image.Image
	Format string
	Err    error
}

var defaultHandler = ImageResizeHandler{}

func LoadImage(fullPath string) *ImageInfo {
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

func (h *ImageResizeHandler) LoadImage(fullPath string) *ImageInfo {
	img, format, err := imagex.LoadImage(fullPath, "")
	return &ImageInfo{Path: fullPath, Image: img, Format: format, Err: err}
}

func (h *ImageResizeHandler) ResizeImage(source image.Image, width, height uint) (img image.Image, err error) {
	return resizex.ResizeImage(source, width, height)
}

func (h *ImageResizeHandler) SaveImage(img image.Image, fullPath string, format formatx.ImageFormat, options interface{}) error {
	dir, err := filex.GetUpDir(fullPath)
	if nil != err {
		return err
	}
	if !filex.IsExist(dir) {
		os.MkdirAll(dir, os.ModePerm)
	}
	return imagex.SaveImage(img, fullPath, format, options)
}
