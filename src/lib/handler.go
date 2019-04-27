//
//Created by xuzhuoxi
//on 2019-04-27.
//@author xuzhuoxi
//
package lib

import (
	"os"
	"image"
	"github.com/nfnt/resize"
)

var defaultHandler = ImageResizeHandler{}

func LoadImage(fullPath string) (img image.Image, err error) {
	return defaultHandler.LoadImage(fullPath)
}

func ResizeImage(source image.Image, width, height uint) (img image.Image, err error) {
	return defaultHandler.ResizeImage(source, width, height)
}

func SaveImage(img image.Image, fullPath string, format ImageFormat, options interface{}) error {
	return defaultHandler.SaveImage(img, fullPath, format, options)
}

type IImageResizeHandler interface {
	LoadImage(fullPath string) (img image.Image, err error)
	ResizeImage(source image.Image, width, height uint) (img image.Image, err error)
	SaveImage(img image.Image, fullPath string, format ImageFormat, options interface{}) error
}

type ImageResizeHandler struct{}

func (h *ImageResizeHandler) LoadImage(fullPath string) (img image.Image, err error) {
	file, _ := os.Open(fullPath)
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}

func (h *ImageResizeHandler) ResizeImage(source image.Image, width, height uint) (img image.Image, err error) {
	return resize.Resize(width, height, source, resize.Lanczos3), nil
}

func (h *ImageResizeHandler) SaveImage(img image.Image, fullPath string, format ImageFormat, options interface{}) error {
	os.Open(fullPath)
	file, _ := os.Create(fullPath)
	defer file.Close()
	return format.Encode(file, img, options)
}
