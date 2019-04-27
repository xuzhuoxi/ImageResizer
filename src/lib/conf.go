//
//Created by xuzhuoxi
//on 2019-04-26.
//@author xuzhuoxi
//
package lib

import (
	"flag"
	"github.com/xuzhuoxi/infra-go/osxu"
	"errors"
	"strings"
	"strconv"
)

type Config struct {
	// 来源地址，文件夹或文件
	// 如果是文件夹，处理当层全部合法文件
	// 合法性：类型支持,大小支持(长宽)
	InPath string
	// 来源地址，文件夹
	OutPath string

	// 输出大小
	OutSizes []int
	// 输出的文件类型
	// 如果InPath为文件夹，此为必选项
	// 如果InPath为文件，可选
	// 		有：输出对应类型
	// 		无：输出原文件类型
	OutFormat string
	// 压缩比,非压缩文件类型忽略
	OutRatio int
}

// -size 	必选	输出大小		整数[,整数]
// -in 		可选	输入			字符串路径，文件夹或文件,"./"开头视为相对路径
// -out 	可选	输出			字符串路径，文件夹,"./"开头视为相对路径
// -format 	可选	输出文件格式	图像格式[pngx,jpeg,gifx,jpg]
// -ratio 	可选	压缩比			整数(0,100]
func ParseFlag() (cfg *Config, err error) {
	basePath := osxu.RunningBaseDir()
	in := flag.String("in", basePath, "Input Path! ")
	out := flag.String("out", basePath, "Output Path! ")
	size := flag.String("size", "", "Size Config!")
	format := flag.String("format", "", "Format Config!")
	ratio := flag.Int("ratio", 75, "Ratio Config!")
	flag.Parse()

	if nil == size || "" == *size {
		return nil, errors.New("Size No Define! ")
	}
	InPath := osxu.FormatDirPath(*in)
	if strings.Index(InPath, "./") == 0 {
		InPath = basePath + InPath
	}
	OutPath := osxu.FormatDirPath(*out)
	if strings.Index(OutPath, "./") == 0 {
		OutPath = basePath + OutPath
	}
	if osxu.IsExist(OutPath) && !osxu.IsFolder(OutPath) {
		return nil, errors.New("Out Config Error! ")
	}
	sizes := strings.Split(*size, ",")
	if nil == sizes || len(sizes) == 0 {
		return nil, errors.New("Size Define Empty! ")
	}
	OutSizes := []int{}
	for _, s := range sizes {
		sVal, err := strconv.Atoi(s)
		if nil != err {
			return nil, errors.New("Size Define Error: " + s)
		}
		OutSizes = append(OutSizes, sVal)
	}
	OutFormat := *format
	if "" != OutFormat && !CheckFormat(OutFormat) {
		return nil, errors.New("Format Define Error: " + OutFormat)
	}
	OutRatio := *ratio
	return &Config{InPath: InPath, OutPath: OutPath, OutSizes: OutSizes, OutFormat: OutFormat, OutRatio: OutRatio}, nil
}
