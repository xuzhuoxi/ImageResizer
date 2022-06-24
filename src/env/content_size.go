package env

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/slicex"
	"strconv"
	"strings"
)

type ImageSize struct {
	Width  uint
	Height uint
}

func NewSizeContext(env string, include string, source, target string, size string, oneByOne bool, format string, ratio int) *SizeContext {
	return &SizeContext{envPath: env, include: include, source: source, target: target, size: size, oneByOne: oneByOne,
		format: format, ratio: ratio}
}

type SizeContext struct {
	envPath  string
	include  string
	source   string
	target   string
	size     string
	oneByOne bool

	format string
	ratio  int

	subIncludes []string
	sourceList  []string
	sizeList    []ImageSize
}

func (c *SizeContext) String() string {
	return fmt.Sprintf("{Env=%s, One=%t, SrcLen=%d, SizeLen=%d}", c.envPath, c.oneByOne, len(c.sourceList), len(c.sizeList))
}

func (c *SizeContext) Mode() ResizeMode {
	return ModeSize
}

func (c *SizeContext) EnvPath() string {
	return c.envPath
}

func (c *SizeContext) OneByOne() bool {
	return c.oneByOne
}

func (c *SizeContext) Format() string {
	return c.format
}

func (c *SizeContext) Ratio() int {
	return c.ratio
}

func (c *SizeContext) Target() string {
	return c.target
}

func (c *SizeContext) FirstSource() string {
	return c.sourceList[0]
}

func (c *SizeContext) SourceList() []string {
	return slicex.CopyString(c.sourceList)
}

func (c *SizeContext) FirstSize() ImageSize {
	return c.sizeList[0]
}

func (c *SizeContext) SizeList() []ImageSize {
	rs := make([]ImageSize, len(c.sizeList))
	copy(rs, c.sizeList)
	return rs
}

func (c *SizeContext) CheckIncludeFile(filePath string) bool {
	return checkFileExt(filePath, c.subIncludes)
}

func (c *SizeContext) GetOutPath(source string, targetDir string, size ImageSize, format string) string {
	fileName, _, _ := filex.SplitFileName(source)
	newFileName := fmt.Sprintf("%s_%dx%d.%s", fileName, size.Width, size.Height, format)
	return filex.Combine(targetDir, newFileName)
}

func (c *SizeContext) InitContext() error {
	c.initInclude()
	if err := c.initSource(); nil != err {
		return err
	}
	if err := c.initTarget(); nil != err {
		return err
	}
	c.ratio = c.getRatio()
	if err := c.initSize(); nil != err {
		return err
	}
	return nil
}

func (c *SizeContext) initInclude() {
	c.subIncludes = strings.Split(c.include, ParamsSep)
}

func (c *SizeContext) initSource() error {
	if c.source == "" {
		return errors.New(fmt.Sprintf("Mode[size] src lack! "))
	}
	if c.oneByOne {
		source, err := handleSourcePath(c.source, c.envPath)
		if nil != err {
			return err
		}
		c.sourceList = []string{source}
		return nil
	}
	list, err := handleSourceList(c.source, c.envPath)
	if nil != err {
		return err
	}
	c.sourceList = list
	return nil
}

func (c *SizeContext) initTarget() error {
	if c.target == "" {
		return errors.New(fmt.Sprintf("Mode[size] tar lack! "))
	}
	c.target = filex.Combine(c.envPath, c.target)
	return nil
}

func (c *SizeContext) initSize() error {
	if c.size == "" {
		return errors.New(fmt.Sprintf("Mode[size] size lack! "))
	}
	scaleList := strings.Split(c.size, ParamsSep)
	c.sizeList = make([]ImageSize, 0, len(scaleList))
	for _, v := range scaleList {
		size, err := c.parseSize(v)
		if nil != err {
			return err
		}
		c.sizeList = append(c.sizeList, size)
	}
	return nil
}

func (c *SizeContext) parseSize(sizeStr string) (size ImageSize, err error) {
	sizeStr = strings.ToLower(sizeStr)
	if strings.Index(sizeStr, "x") != -1 {
		vArr := strings.Split(sizeStr, "x")
		if len(vArr) != 2 {
			return ImageSize{}, errors.New(fmt.Sprintf("Mode[size] size format error at [%s]! ", sizeStr))
		}
		w, err := strconv.ParseUint(vArr[0], 10, 32)
		if nil != err {
			return ImageSize{}, errors.New(fmt.Sprintf("Mode[size] size format error at Width[%s]! ", err))
		}
		h, err := strconv.ParseUint(vArr[1], 10, 32)
		if nil != err {
			return ImageSize{}, errors.New(fmt.Sprintf("Mode[size] size format error at Height[%s]! ", err))
		}
		return ImageSize{Width: uint(w), Height: uint(h)}, nil
	}
	l, err := strconv.ParseUint(sizeStr, 10, 32)
	if nil != err {
		return ImageSize{}, errors.New(fmt.Sprintf("Mode[size] size format error at Height[%s]! ", err))
	}
	return ImageSize{Width: uint(l), Height: uint(l)}, nil
}

func (c *SizeContext) getRatio() int {
	return GetRatio(c.ratio)
}
