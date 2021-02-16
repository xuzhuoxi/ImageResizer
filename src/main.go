//
//Created by xuzhuoxi
//on 2019-04-26.
//@author xuzhuoxi
//
package main

import (
	"fmt"
	"github.com/viphxin/xingo/logger"
	"github.com/xuzhuoxi/IconGen/src/lib"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/jpegx"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/pngx"
	"github.com/xuzhuoxi/infra-go/logx"
	"image/jpeg"
	"os"
)

var (
	globalLogger   logx.ILogger
	globalOutRatio int
)

func file2Size(inFilePath string, outFilePath string, size lib.Size, format string) {
	img, _, err := lib.LoadImage(inFilePath)
	if nil != err {
		return
	}
	newImg, _ := lib.ResizeImage(img, uint(size.Width), uint(size.Height))
	lib.SaveImage(newImg, outFilePath, formatx.ImageFormat(format), &jpeg.Options{Quality: globalOutRatio})
	globalLogger.Infoln("IconGen Gen Image:", outFilePath)
}

func file2Sizes(inFilePath string, outFolder string, sizes []lib.Size, format string) {
	_, fileName := filex.Split(inFilePath)
	baseName, _, _ := filex.SplitFileName(fileName)
	for _, size := range sizes {
		outFileName := fmt.Sprintf("%s_%dx%d.%s", baseName, size.Width, size.Height, format)
		outFilePath := filex.Combine(outFolder, outFileName)
		file2Size(inFilePath, outFilePath, size, format)
	}
}

func CheckInPath(cfg *lib.Config) bool {
	if !filex.IsExist(cfg.InPath) {
		logger.Error("InPath does not Exist:", cfg.InPath)
		return false
	}
	if cfg.InFolder != filex.IsFolder(cfg.InPath) {
		logger.Error(
			fmt.Sprintf("Contradiction between InPath(%s) and InFolder(%v)", cfg.InPath, cfg.InFolder))
		return false
	}
	return true
}

func CheckOutPath(cfg *lib.Config) bool {
	if filex.IsExist(cfg.OutPath) {
		if cfg.OutFolder != filex.IsFolder(cfg.OutPath) {
			logger.Error(
				fmt.Sprintf("Contradiction between OutPath(%s) and OutFolder(%v)", cfg.OutPath, cfg.OutFolder))
			return false
		}
		return true
	}
	if cfg.OutFolder {
		err := os.MkdirAll(cfg.OutPath, os.ModePerm)
		if nil != err {
			logger.Error("Init OutPath Folder Error! ", err.Error())
			return false
		}
	} else {
		dir, _ := filex.Split(cfg.OutPath)
		err := os.MkdirAll(dir, os.ModePerm)
		if nil != err {
			logger.Error("Init OutPath Folder Error! ", err.Error())
			return false
		}
	}
	return true
}

func CheckPath(cfg *lib.Config) bool {
	if cfg.InFolder && !cfg.OutFolder {
		logger.Error(
			fmt.Sprintf("Contradiction between InFolder(%v) and OutFolder(%v)", cfg.InFolder, cfg.OutFolder))
		return false
	}
	return true
}

func CheckSize(cfg *lib.Config) bool {
	if !cfg.OutFolder && len(cfg.OutSizes) != 1 {
		logger.Error(
			fmt.Sprintf("Contradiction between OutFolder(%v) and Size's Len(%d)", cfg.OutFolder, len(cfg.OutSizes)))
		return false
	}
	return true
}

func main() {
	globalLogger = logx.NewLogger()
	globalLogger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})

	cfg, err := lib.ParseFlag()
	if err != nil {
		logger.Error(err)
		return
	}
	globalOutRatio = cfg.OutRatio

	if !CheckInPath(cfg) || !CheckOutPath(cfg) || !CheckPath(cfg) || !CheckSize(cfg) {
		return
	}

	globalLogger.Infoln("IconGen Start...")
	globalLogger.Infoln("Info:", cfg)

	// 文件 => 文件， 1个Size
	if !cfg.InFolder && !cfg.OutFolder {
		format := cfg.OutFormat
		if "" == format {
			_, fileName := filex.Split(cfg.InPath)
			_, _, ext := filex.SplitFileName(fileName)
			format = ext
		}
		file2Size(cfg.InPath, cfg.OutPath, cfg.OutSizes[0], format)
		return
	}

	// 文件 => 文件夹， 不限制Size
	if !cfg.InFolder && cfg.OutFolder {
		format := cfg.OutFormat
		if "" == format {
			_, fileName := filex.Split(cfg.InPath)
			_, _, ext := filex.SplitFileName(fileName)
			format = ext
		}
		file2Sizes(cfg.InPath, cfg.OutPath, cfg.OutSizes, format)
		return
	}

	// 文件夹 => 文件夹，不限制Size
	var list []string
	err = filex.WalkInDir(cfg.InPath, func(path string, info os.FileInfo, err error) error {
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
		format := cfg.OutFormat
		if "" == format {
			_, fileName := filex.Split(file)
			_, _, ext := filex.SplitFileName(fileName)
			format = ext
		}
		file2Sizes(file, cfg.OutPath, cfg.OutSizes, format)
	}
	globalLogger.Infoln("IconGen Finish.")
}
