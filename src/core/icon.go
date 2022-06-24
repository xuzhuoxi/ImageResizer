// Create on 2022/6/24
// @author xuzhuoxi
package core

import (
	"fmt"
	"github.com/xuzhuoxi/ImageResizer/src/env"
	"github.com/xuzhuoxi/ImageResizer/src/lib"
)

var defaultIconHandler = &iconHandler{}

func HandleIcon(ctx *env.IconContext) {
	if nil == ctx {
		return
	}
	globalLogger.Infoln(fmt.Sprintf("Resize Start in Mode[icon]: %v", ctx))
	defaultIconHandler.handleContext(ctx)
	globalLogger.Infoln(fmt.Sprintf("Resize Finish in Mode[icon]."))
}

type iconHandler struct{}

func (h *iconHandler) handleContext(ctx *env.IconContext) {
	imgInfo := lib.LoadImage(ctx.Source())
	if imgInfo.Err != nil {
		globalLogger.Warnln(fmt.Sprintf("\t[HandleContext.LoadImage] Error by[%s]", imgInfo.Err))
		return
	}
	cfg := ctx.Config()
	imgFormat := GetFormat(ctx.Format(), cfg.Format, imgInfo.Format)
	options := GetOptions(imgFormat, ctx.Ratio())
	for _, size := range cfg.List {
		sizeImg, err := lib.ResizeImage(imgInfo.Image, size.Width(), size.Height())
		if err != nil {
			globalLogger.Warnln(fmt.Sprintf("\t[HandleContext.ResizeImage] Error by[%s]", err))
			return
		}
		path := ctx.GetOutPath(size, string(imgFormat))
		err = lib.SaveImage(sizeImg, path, imgFormat, options)
		if nil != err {
			globalLogger.Warnln(fmt.Sprintf("\t[HandleContext.SaveImage] Error by[%s]", err))
			return
		}
		globalLogger.Infoln("\t", ctx.Source()[len(ctx.EnvPath())+1:], "=>", path[len(ctx.EnvPath())+1:])
	}
}
