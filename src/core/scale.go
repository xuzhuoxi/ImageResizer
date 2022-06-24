// Create on 2022/6/24
// @author xuzhuoxi
package core

import (
	"fmt"
	"github.com/xuzhuoxi/ImageResizer/src/env"
	"github.com/xuzhuoxi/ImageResizer/src/lib"
	"github.com/xuzhuoxi/infra-go/filex"
	"image"
	"image/jpeg"
	"io/fs"
	"math"
)

var defaultScaleHandler = &scaleHandler{}

func HandleScale(ctx *env.ScaleContext) {
	if nil == ctx {
		return
	}
	globalLogger.Infoln(fmt.Sprintf("Resize Start in Mode[scale]: %v", ctx))
	if ctx.OneByOne() {
		defaultScaleHandler.handlerOneByOne(ctx)
	} else {
		defaultScaleHandler.handleImages(ctx)
	}
	globalLogger.Infoln(fmt.Sprintf("Resize Finish in Mode[scale]: %v", ctx))
}

type scaleHandler struct{}

func (h *scaleHandler) handlerOneByOne(ctx *env.ScaleContext) {
	imgInfo := lib.LoadImage(ctx.FirstSource())
	if imgInfo.Err != nil {
		globalLogger.Warnln(imgInfo.Err)
		return
	}
	scale := ctx.FirstScale()
	width, height := h.getImageSize(imgInfo.Image, scale)
	sizeImg, err := lib.ResizeImage(imgInfo.Image, width, height)
	if err != nil {
		globalLogger.Warnln(err)
		return
	}
	imgFormat := GetFormat(ctx.Format(), imgInfo.Format)
	err = lib.SaveImage(sizeImg, ctx.Target(), imgFormat, &jpeg.Options{Quality: 75})
	if nil != err {
		globalLogger.Warnln(err)
		return
	}
}

func (h *scaleHandler) handleImages(ctx *env.ScaleContext) {
	//fmt.Println("scaleHandler.handleImages")
	for _, source := range ctx.SourceList() {
		if filex.IsDir(source) {
			filex.WalkInDir(source, func(path string, info fs.FileInfo, err error) error {
				if nil != err {
					return err
				}
				if ctx.CheckIncludeFile(path) {
					if err := h.handleImage(ctx, path); nil != err {
						globalLogger.Warnln(err)
						return err
					}
				}
				return nil
			})
		}
		if err := h.handleImage(ctx, source); nil != err {
			globalLogger.Warnln(err)
			return
		}
	}
}

func (h *scaleHandler) handleImage(ctx *env.ScaleContext, imgPath string) error {
	imgInfo := lib.LoadImage(imgPath)
	if imgInfo.Err != nil {
		return imgInfo.Err
	}
	for _, scale := range ctx.ScaleList() {
		//fmt.Println("scaleHandler.handleImage:", imgPath, scale)
		width, height := h.getImageSize(imgInfo.Image, scale)
		sizeImg, err := lib.ResizeImage(imgInfo.Image, width, height)
		if err != nil {
			return err
		}
		imgFormat := GetFormat(ctx.Format(), imgInfo.Format)
		path := ctx.GetOutPath(imgInfo.Path, ctx.Target(), scale, string(imgFormat))
		options := GetOptions(imgFormat, ctx.Ratio())
		//fmt.Println("scaleHandler.handleImage2:", imgPath, width, height, path, imgFormat, options)
		err = lib.SaveImage(sizeImg, path, imgFormat, options)
		if nil != err {
			return err
		}
		continue
	}
	return nil
}

func (h *scaleHandler) getImageSize(image image.Image, scale float64) (width uint, height uint) {
	imgSize := image.Bounds().Size()
	width = uint(math.Floor(float64(imgSize.X) * scale))
	height = uint(math.Floor(float64(imgSize.Y) * scale))
	return
}
