// Create on 2022/6/24
// @author xuzhuoxi
package core

import (
	"fmt"
	"github.com/xuzhuoxi/ImageResizer/src/env"
	"github.com/xuzhuoxi/ImageResizer/src/lib"
	"github.com/xuzhuoxi/infra-go/filex"
	"image/jpeg"
	"io/fs"
)

var defaultSizeHandler = &sizeHandler{}

func HandleSize(ctx *env.SizeContext) {
	if nil == ctx {
		return
	}
	globalLogger.Infoln(fmt.Sprintf("Resize Start in Mode[size]: %v", ctx))
	if ctx.OneByOne() {
		defaultSizeHandler.handlerOneByOne(ctx)
	} else {
		defaultSizeHandler.handleImages(ctx)
	}
	globalLogger.Infoln(fmt.Sprintf("Resize Finish in Mode[size]: %v", ctx))
}

type sizeHandler struct{}

func (h *sizeHandler) handlerOneByOne(ctx *env.SizeContext) {
	imgInfo := lib.LoadImage(ctx.FirstSource())
	if imgInfo.Err != nil {
		globalLogger.Warnln(imgInfo.Err)
		return
	}
	size := ctx.FirstSize()
	sizeImg, err := lib.ResizeImage(imgInfo.Image, size.Width, size.Height)
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

func (h *sizeHandler) handleImages(ctx *env.SizeContext) {
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

func (h *sizeHandler) handleImage(ctx *env.SizeContext, imgPath string) error {
	imgInfo := lib.LoadImage(imgPath)
	if imgInfo.Err != nil {
		return imgInfo.Err
	}
	for _, size := range ctx.SizeList() {
		sizeImg, err := lib.ResizeImage(imgInfo.Image, size.Width, size.Height)
		if err != nil {
			return err
		}
		imgFormat := GetFormat(ctx.Format(), imgInfo.Format)
		path := ctx.GetOutPath(imgInfo.Path, ctx.Target(), size, string(imgFormat))
		options := GetOptions(imgFormat, ctx.Ratio())
		err = lib.SaveImage(sizeImg, path, imgFormat, options)
		if nil != err {
			return err
		}
		continue
	}
	return nil
}
