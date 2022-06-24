// Create on 2022/6/24
// @author xuzhuoxi
package core

import (
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"github.com/xuzhuoxi/infra-go/logx"
	"image/jpeg"
)

var (
	globalLogger logx.ILogger
)

func RegisterLogger(logger logx.ILogger) {
	globalLogger = logger
}

func GetFormat(first string, other ...string) formatx.ImageFormat {
	if "" != first {
		return formatx.ImageFormat(first)
	}
	for _, v := range other {
		if "" != v {
			return formatx.ImageFormat(v)
		}
	}
	return formatx.Auto
}

func GetOptions(format formatx.ImageFormat, ratio int) interface{} {
	if formatx.PNG == format {
		return nil
	}
	if formatx.Auto == format {
		return nil
	}
	return &jpeg.Options{Quality: ratio}
}
