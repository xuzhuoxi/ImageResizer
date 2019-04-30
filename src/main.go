//
//Created by xuzhuoxi
//on 2019-04-26.
//@author xuzhuoxi
//
package main

import (
	"github.com/xuzhuoxi/IconGen/src/lib"
	_ "github.com/xuzhuoxi/IconGen/src/lib/png"
	_ "github.com/xuzhuoxi/IconGen/src/lib/jpeg"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"os"
	"fmt"
	"image/jpeg"
)

func main() {
	logger := logx.NewLogger()
	logger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})
	cfg, err := lib.ParseFlag()
	if err != nil {
		logger.Error(err)
		return
	}
	if !osxu.IsExist(cfg.InPath) {
		logger.Error("InPath does net Exist! ")
		return
	}
	if !osxu.IsExist(cfg.OutPath) { //输出路径不存在，创建
		err := os.MkdirAll(cfg.OutPath, os.ModePerm)
		if nil != err {
			logger.Error("OutPath Error! ")
			return
		}
	}
	handle := func(filePath string) {
		_, fileName := osxu.SplitFilePath(filePath)
		baseName, extName := osxu.SplitFileName(fileName)
		img, err := lib.LoadImage(filePath)
		if nil != err {
			return
		}
		fm := cfg.OutFormat
		if "" == fm {
			fm = extName
		}
		for _, size := range cfg.OutSizes {
			newImg, _ := lib.ResizeImage(img, uint(size.Width), uint(size.Height))
			fileName := fmt.Sprintf("%s_%dx%d.%s", baseName, size.Width, size.Height, fm)
			fileFullPath := cfg.OutPath + fileName
			lib.SaveImage(newImg, fileFullPath, lib.ImageFormat(fm), &jpeg.Options{Quality: cfg.OutRatio})
			logger.Infoln("IconGen Gen Image:", fileFullPath)
		}
	}
	logger.Infoln("IconGen Start...")
	if !osxu.IsFolder(cfg.InPath) {
		handle(cfg.InPath)
	} else {
		list, err := osxu.GetFolderFileList(cfg.InPath, false, func(fileInfo os.FileInfo) bool {
			extName := osxu.GetExtensionName(fileInfo.Name())
			if !lib.CheckFormat(extName) {
				return false
			}
			return true
		})
		if nil != err {
			logger.Error(err)
			return
		}
		for _, file := range list {
			handle(file.FullPath())
		}
	}
	logger.Infoln("IconGen Finish.")
}
