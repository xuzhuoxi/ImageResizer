package env

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/slicex"
	"strconv"
	"strings"
)

func NewScaleContext(env string, include string, source, target string, scale string, oneByOne bool, format string, ratio int) *ScaleContext {
	return &ScaleContext{envPath: env, include: include, source: source, target: target, scale: scale, oneByOne: oneByOne,
		format: format, ratio: ratio}
}

type ScaleContext struct {
	envPath  string
	include  string
	source   string
	target   string
	scale    string
	oneByOne bool

	format string
	ratio  int

	subIncludes []string
	sourceList  []string
	scaleList   []float64
}

func (c *ScaleContext) String() string {
	return fmt.Sprintf("{Env=%s, One=%t, SrcLen=%d, ScaleLen=%d}", c.envPath, c.oneByOne, len(c.sourceList), len(c.scaleList))
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

func (c *ScaleContext) Format() string {
	return c.format
}

func (c *ScaleContext) Ratio() int {
	return c.ratio
}

func (c *ScaleContext) Target() string {
	return c.target
}

func (c *ScaleContext) FirstSource() string {
	return c.sourceList[0]
}

func (c *ScaleContext) SourceList() []string {
	return slicex.CopyString(c.sourceList)
}

func (c *ScaleContext) FirstScale() float64 {
	return c.scaleList[0]
}

func (c *ScaleContext) ScaleList() []float64 {
	rs := make([]float64, len(c.scaleList))
	copy(rs, c.scaleList)
	return rs
}

func (c *ScaleContext) CheckIncludeFile(filePath string) bool {
	return checkFileExt(filePath, c.subIncludes)
}

func (c *ScaleContext) GetOutPath(source string, targetDir string, scale float64, format string) string {
	fileName, _, _ := filex.SplitFileName(source)
	newFileName := fmt.Sprintf("%s_x%.1f.%s", fileName, scale, format)
	return filex.Combine(targetDir, newFileName)
}

func (c *ScaleContext) InitContext() error {
	c.initInclude()
	if err := c.initSource(); nil != err {
		return err
	}
	if err := c.initTarget(); nil != err {
		return err
	}
	c.ratio = c.getRatio()
	if err := c.initScale(); nil != err {
		return err
	}
	return nil
}

func (c *ScaleContext) initInclude() {
	c.subIncludes = strings.Split(c.include, ParamsSep)
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

func (c *ScaleContext) initTarget() error {
	if c.target == "" {
		return errors.New(fmt.Sprintf("Mode[scale] tar lack! "))
	}
	c.target = filex.Combine(c.envPath, c.target)
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
