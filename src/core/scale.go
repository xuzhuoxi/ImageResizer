// Create on 2022/6/24
// @author xuzhuoxi
package core

import (
	"errors"
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
		globalLogger.Warnln(fmt.Sprintf("\t[OneByOne.LoadImage] Erryr by[%s]", imgInfo.Err))
		return
	}
	scale := ctx.FirstScale()
	width, height := h.getImageSize(imgInfo.Image, scale)
	sizeImg, err := lib.ResizeImage(imgInfo.Image, width, height)
	if err != nil {
		globalLogger.Warnln(fmt.Sprintf("\t[OneByOne.ResizeImage] Erryr by[%s]", err))
		return
	}
	imgFormat := GetFormat(ctx.Format(), imgInfo.Format)
	err = lib.SaveImage(sizeImg, ctx.Target(), imgFormat, &jpeg.Options{Quality: 75})
	if nil != err {
		globalLogger.Warnln(fmt.Sprintf("\t[OneByOne.SaveImage] Erryr by[%s]", err))
		return
	}
	globalLogger.Infoln("\t", ctx.FirstSource()[len(ctx.EnvPath())+1:], "=>", ctx.Target()[len(ctx.EnvPath())+1:])
}

func (h *scaleHandler) handleImages(ctx *env.ScaleContext) {
	//fmt.Println("scaleHandler.handleImages")
	for _, source := range ctx.SourceList() {
		if filex.IsDir(source) {
			filex.WalkInDir(source, func(path string, info fs.FileInfo, err error) error {
				if nil != err {
					return err
				}
				if info.IsDir() {
					return nil
				}
				if ctx.CheckIncludeFile(path) {
					if err := h.handleImage(ctx, path); nil != err {
						globalLogger.Warnln(err)
						return err
					}
				}
				return nil
			})

		} else {
			if err := h.handleImage(ctx, source); nil != err {
				globalLogger.Warnln(err)
			}
		}
	}
}

func (h *scaleHandler) handleImage(ctx *env.ScaleContext, imgPath string) error {
	imgInfo := lib.LoadImage(imgPath)
	if imgInfo.Err != nil {
		return errors.New(fmt.Sprintf("\t[HandleImage.LoadImage] Erryr by[%s][%s]", imgInfo.Err, imgPath))
	}
	//fmt.Println("sizeHandler.handleImage:", imgPath, imgInfo.Format)
	for _, scale := range ctx.ScaleList() {
		//fmt.Println("scaleHandler.handleImage:", imgPath, scale)
		width, height := h.getImageSize(imgInfo.Image, scale)
		sizeImg, err := lib.ResizeImage(imgInfo.Image, width, height)
		if err != nil {
			return errors.New(fmt.Sprintf("\t[HandleImage.ResizeImage] Erryr by[%s]", err))
		}
		imgFormat := GetFormat(ctx.Format(), imgInfo.Format)
		path := ctx.GetOutPath(imgInfo.Path, ctx.Target(), scale, string(imgFormat))
		options := GetOptions(imgFormat, ctx.Ratio())
		//fmt.Println("scaleHandler.handleImage2:", imgPath, width, height, path, imgFormat, options)
		err = lib.SaveImage(sizeImg, path, imgFormat, options)
		if nil != err {
			return errors.New(fmt.Sprintf("\t[HandleImage.SaveImage] Erryr by[%s]", err))
		}
		globalLogger.Infoln("\t", imgPath[len(ctx.EnvPath())+1:], "=>", path[len(ctx.EnvPath())+1:])
	}
	return nil
}

func (h *scaleHandler) getImageSize(image image.Image, scale float64) (width uint, height uint) {
	imgSize := image.Bounds().Size()
	width = uint(math.Floor(float64(imgSize.X) * scale))
	height = uint(math.Floor(float64(imgSize.Y) * scale))
	return
}
