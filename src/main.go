//
//Created by xuzhuoxi
//on 2019-04-26.
//@author xuzhuoxi
//
package main

import (
	"github.com/xuzhuoxi/ImageResizer/src/core"
	"github.com/xuzhuoxi/ImageResizer/src/env"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/mathx"
	"github.com/xuzhuoxi/infra-go/osxu"
)

var (
	globalLogger logx.ILogger
	logFileName  = "ImageResizer"
)

func main() {
	initLogger()
	cmdFlags := env.ParseFlags()
	iconCtx, sizeCtc, scaleCtx, err := cmdFlags.GetContexts()
	if nil != err {
		globalLogger.Warnln(err)
		return
	}
	core.HandleIcon(iconCtx)
	core.HandleSize(sizeCtc)
	core.HandleScale(scaleCtx)
}

func initLogger() {
	globalLogger = logx.NewLogger()
	globalLogger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})
	globalLogger.SetConfig(logx.LogConfig{Type: logx.TypeRollingFile, Level: logx.LevelAll,
		FileDir: osxu.GetRunningDir(), FileName: logFileName, FileExtName: ".log", MaxSize: 10 * mathx.MB})
	core.RegisterLogger(globalLogger)
}
