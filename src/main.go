//
//Created by xuzhuoxi
//on 2019-04-26.
//@author xuzhuoxi
//
package main

import (
	"fmt"
	"github.com/xuzhuoxi/IconGen/src/lib"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/jpegx"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/pngx"
	"github.com/xuzhuoxi/infra-go/logx"
	"image/jpeg"
	"os"
)

func main() {
	logger := logx.NewLogger()
	logger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})
	cfg, err := lib.ParseFlag()
	if err != nil {
		logger.Error(err)
		return
	}
	if !filex.IsExist(cfg.InPath) {
		logger.Error("InPath does net Exist! ")
		return
	}
	if !filex.IsExist(cfg.OutPath) { //输出路径不存在，创建
		err := os.MkdirAll(cfg.OutPath, os.ModePerm)
		if nil != err {
			logger.Error("OutPath Error! ")
			return
		}
	}
	handle := func(filePath string) {
		_, fileName := filex.Split(filePath)
		baseName, _, extName := filex.SplitFileName(fileName)
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
			fileFullPath := filex.Combine(cfg.OutPath, fileName)
			lib.SaveImage(newImg, fileFullPath, formatx.ImageFormat(fm), &jpeg.Options{Quality: cfg.OutRatio})
			logger.Infoln("IconGen Gen Image:", fileFullPath)
		}
	}
	logger.Infoln("IconGen Start...")
	if !filex.IsFolder(cfg.InPath) {
		handle(cfg.InPath)
	} else {
		var list []string
		err := filex.WalkInDir(cfg.InPath, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if formatx.CheckFormatRegistered(filex.GetExtWithoutDot(info.Name())) {
				list = append(list, path)
			}
			return nil
		})
		if nil != err {
			logger.Error(err)
			return
		}
		for _, file := range list {
			handle(file)
		}
	}
	logger.Infoln("IconGen Finish.")
}
