package env

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/slicex"
	"strconv"
	"strings"
)

type ImageSize struct {
	Width  int
	Height int
}

func NewSizeContext(env string, source, target string, size string, oneByOne bool, format string, ratio int) *SizeContext {
	return &SizeContext{envPath: env, source: source, target: target, size: size, oneByOne: oneByOne,
		format: format, ratio: ratio}
}

type SizeContext struct {
	envPath  string
	source   string
	target   string
	size     string
	oneByOne bool

	format string
	ratio  int

	sourceList []string
	sizeList   []ImageSize
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

func (c *SizeContext) FirstSource() string {
	return c.sourceList[0]
}

func (c *SizeContext) SourceList() []string {
	return slicex.CopyString(c.sourceList)
}

func (c *SizeContext) SizeList() []ImageSize {
	rs := make([]ImageSize, len(c.sizeList))
	copy(rs, c.sizeList)
	return rs
}

func (c *SizeContext) InitContext() error {
	if err := c.initSource(); nil != err {
		return err
	}
	if c.target == "" {
		return errors.New(fmt.Sprintf("Mode[scale] tar lack! "))
	}
	if err := CheckFormat(c.format, c.ratio); nil != err {
		return errors.New(fmt.Sprintf("Mode[scale] %s", err))
	}
	c.ratio = c.getRatio()
	if err := c.initSize(); nil != err {
		return err
	}
	return nil
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
		w, err := strconv.Atoi(vArr[0])
		if nil != err {
			return ImageSize{}, errors.New(fmt.Sprintf("Mode[size] size format error at Width[%s]! ", err))
		}
		h, err := strconv.Atoi(vArr[1])
		if nil != err {
			return ImageSize{}, errors.New(fmt.Sprintf("Mode[size] size format error at Height[%s]! ", err))
		}
		return ImageSize{Width: w, Height: h}, nil
	}
	l, err := strconv.Atoi(sizeStr)
	if nil != err {
		return ImageSize{}, errors.New(fmt.Sprintf("Mode[size] size format error at Height[%s]! ", err))
	}
	return ImageSize{Width: l, Height: l}, nil
}

func (c *SizeContext) getRatio() int {
	return GetRatio(c.ratio)
}
