package env

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/slicex"
	"strings"
)

const (
	ParamsSep          = ","
	DefaultRatio       = 75
	IconNameSubstitute = "{{name}}"
)

type ResizeMode = string

const (
	ModeIcon  = "icon"
	ModeScale = "scale"
	ModeSize  = "size"
)

func GetRatio(firstRatio int, ratio ...int) int {
	if firstRatio != 0 {
		return firstRatio
	}
	if len(ratio) != 0 {
		for _, r := range ratio {
			if r != 0 {
				return r
			}
		}
	}
	return DefaultRatio
}

func handleSourceList(source string, envPath string) (sourceList []string, err error) {
	srcList := strings.Split(source, ParamsSep)
	sourceList = make([]string, 0, len(srcList))
	for _, v := range srcList {
		src, err := handleSourcePath(v, envPath)
		if nil != err {
			return nil, err
		}
		sourceList = append(sourceList, src)
	}
	return
}

func handleSourcePath(source string, envPath string) (newSource string, err error) {
	if filex.IsExist(source) {
		return source, nil
	}
	newPath := filex.Combine(envPath, source)
	if filex.IsExist(newPath) {
		return newPath, nil
	}
	return "", errors.New(fmt.Sprintf("source is not exist at [%s]! ", source))
}

func checkFileExt(filePath string, includes []string) bool {
	if len(includes) == 0 {
		return false
	}
	ext := filex.GetExtWithoutDot(filePath)
	ext = strings.ToLower(ext)
	return slicex.ContainsString(includes, ext)
}
