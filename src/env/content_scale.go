package env

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/slicex"
	"strconv"
	"strings"
)

func NewScaleContext(env string, source, target string, scale string, oneByOne bool, format string, ratio int) *ScaleContext {
	return &ScaleContext{envPath: env, source: source, target: target, scale: scale, oneByOne: oneByOne,
		format: format, ratio: ratio}
}

type ScaleContext struct {
	envPath  string
	source   string
	target   string
	scale    string
	oneByOne bool

	format string
	ratio  int

	sourceList []string
	scaleList  []float64
}

func (c *ScaleContext) Mode() ResizeMode {
	return ModeScale
}

func (c *ScaleContext) EnvPath() string {
	return c.envPath
}

func (c *ScaleContext) OneByOne() bool {
	return c.oneByOne
}

func (c *ScaleContext) FirstSource() string {
	return c.sourceList[0]
}

func (c *ScaleContext) SourceList() []string {
	return slicex.CopyString(c.sourceList)
}

func (c *ScaleContext) ScaleList() []float64 {
	rs := make([]float64, len(c.scaleList))
	copy(rs, c.scaleList)
	return rs
}

func (c *ScaleContext) GetOutPath(source string, targetDir string, scale float64) string {
	fileName, _, ext := filex.SplitFileName(source)
	newFileName := fmt.Sprintf("%s_x{%f}.{%s}", fileName, scale, ext)
	return filex.Combine(targetDir, newFileName)
}

func (c *ScaleContext) InitContext() error {
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
	if err := c.initScale(); nil != err {
		return err
	}
	return nil
}

func (c *ScaleContext) initSource() error {
	if c.source == "" {
		return errors.New(fmt.Sprintf("Mode[scale] src lack! "))
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

func (c *ScaleContext) initScale() error {
	if c.scale == "" {
		return errors.New(fmt.Sprintf("Mode[scale] scale lack! "))
	}
	scaleList := strings.Split(c.scale, ParamsSep)
	c.scaleList = make([]float64, 0, len(scaleList))
	for _, v := range scaleList {
		f, err := strconv.ParseFloat(v, 64)
		if nil != err {
			return errors.New(fmt.Sprintf("Mode[scale] scale error at [%s]! ", err))
		}
		c.scaleList = append(c.scaleList, f)
	}
	return nil
}

func (c *ScaleContext) getRatio() int {
	return GetRatio(c.ratio)
}
