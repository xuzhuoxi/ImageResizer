//
//Created by xuzhuoxi
//on 2019-04-26.
//@author xuzhuoxi
//
package lib2

import (
	"errors"
	"flag"
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"strconv"
	"strings"
)

type Size struct {
	//宽
	Width int
	//高
	Height int
}

type Config struct {
	// 基础目录，用于拼接相对路径
	// 如果当前为 空 或 "." 或 "./",则使用运行时路径
	BasePath string

	// 来源地址，文件夹或文件
	// 如果是文件夹，处理当层全部合法文件
	// 合法性：类型支持,大小支持(长宽)
	InPath string
	// 来源地址是否为文件夹
	InFolder bool

	// 输出地址，文件夹
	OutPath string
	// 输出地址是否为文件夹
	OutFolder bool

	// 输出大小
	OutSizes []Size
	// 输出的文件类型
	// 如果InPath为文件夹，此为必选项
	// 如果InPath为文件，可选
	// 		有：输出对应类型
	// 		无：输出原文件类型
	OutFormat string
	// 压缩比,非压缩文件类型忽略
	OutRatio int
}

func (cfg *Config) String() string {
	return fmt.Sprintf("Config{\nBasePath=%s,\nInPath=%s,\nOutPath=%s,\nOutSize=%s,\nOutFormat=%s,\nOutRatio=%d}",
		cfg.BasePath, cfg.InPath, cfg.OutPath, fmt.Sprint(cfg.OutSizes), cfg.OutFormat, cfg.OutRatio)
}

var RunningDir = osxu.GetRunningDir()

// -base 		可选	自定义基目录				字符串路径，文件夹或文件,"./"开头视为相对路径
// -size 		必选	输出大小					[整数/宽x高],...
// -in 			可选	来源地址					字符串路径，文件夹或文件,"./"开头视为相对路径
// -inFolder 	可选	来源地址是否为文件夹		字符串路径，文件夹或文件,"./"开头视为相对路径
// -out 		可选	输出地址					字符串路径，文件夹,"./"开头视为相对路径
// -inFolder 	可选	输出地址是否为文件夹		字符串路径，文件夹或文件,"./"开头视为相对路径
// -format 		可选	输出文件格式				图像格式[pngx,jpeg,gifx,jpg]
// -ratio 		可选	压缩比					整数(0,100]
func ParseFlag() (cfg *Config, err error) {
	base := flag.String("base", "", "Base Path! ")
	in := flag.String("in", "", "Input Path! ")
	inFolder := flag.Bool("inFolder", true, "Input is Folder? ")
	out := flag.String("out", "", "Output Path! ")
	outFolder := flag.Bool("outFolder", true, "Output is Folder? ")
	size := flag.String("size", "", "Size Config!")
	format := flag.String("format", "", "Format Config!")
	ratio := flag.Int("ratio", 75, "Ratio Config!")
	flag.Parse()

	if nil == size || "" == *size {
		return nil, errors.New("Size No Define! ")
	}
	BasePath := *base
	if "" == BasePath || "." == BasePath || "./" == BasePath {
		BasePath = RunningDir
	} else if strings.Index(BasePath, "./") == 0 {
		BasePath = filex.Combine(RunningDir, BasePath)
	}
	InPath := filex.FormatPath(*in)
	if "" == InPath || strings.Index(InPath, "./") == 0 {
		InPath = filex.Combine(BasePath, InPath)
	}
	OutPath := filex.FormatPath(*out)
	if "" == OutPath || strings.Index(OutPath, "./") == 0 {
		OutPath = filex.Combine(BasePath, OutPath)
	}
	if filex.IsExist(OutPath) && !filex.IsFolder(OutPath) {
		return nil, errors.New("Out Config Error! ")
	}
	sizes := strings.Split(*size, ",")
	if nil == sizes || len(sizes) == 0 {
		return nil, errors.New("Size Define Empty! ")
	}
	var OutSizes []Size
	for _, s := range sizes {
		ls := strings.ToLower(s)
		if strings.Index(ls, "x") != -1 {
			wh := strings.Split(ls, "x")
			if len(wh) != 2 {
				return nil, errors.New("Size Define Error: " + s)
			}
			width, err := strconv.Atoi(wh[0])
			if nil != err {
				return nil, errors.New("Size Define Error: " + s)
			}
			height, err := strconv.Atoi(wh[1])
			if nil != err {
				return nil, errors.New("Size Define Error: " + s)
			}
			OutSizes = append(OutSizes, Size{Width: width, Height: height})
		} else {
			sVal, err := strconv.Atoi(s)
			if nil != err {
				return nil, errors.New("Size Define Error: " + s)
			}
			OutSizes = append(OutSizes, Size{Width: sVal, Height: sVal})
		}
	}
	OutFormat := *format
	if "" != OutFormat && !formatx.CheckFormatRegistered(OutFormat) {
		return nil, errors.New("Format Define Error: " + OutFormat)
	}
	OutRatio := *ratio
	return &Config{
		InPath: InPath, InFolder: *inFolder,
		OutPath: OutPath, OutFolder: *outFolder,
		OutSizes: OutSizes, OutFormat: OutFormat, OutRatio: OutRatio}, nil
}
