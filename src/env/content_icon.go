package env

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ImageResizer/src/lib"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"strconv"
	"strings"
)

type IconSize struct {
	Name string `yaml:"name"` // 文件路径
	Size string `yaml:"size"` // 尺寸,格式: 长x宽

	width  uint
	height uint
}

func (s *IconSize) Width() uint {
	s.initSize()
	return s.width
}

func (s *IconSize) Height() uint {
	s.initSize()
	return s.height
}

func (s *IconSize) initSize() {
	if s.width > 0 || s.height > 0 {
		return
	}
	arr := strings.Split(s.Size, "x")
	width, err := strconv.ParseUint(arr[0], 10, 32)
	if nil != err {
		panic(err)
	}
	height, err := strconv.ParseUint(arr[1], 10, 32)
	if nil != err {
		panic(err)
	}
	s.width, s.height = uint(width), uint(height)
}

type IconCfg struct {
	DefaultName string     `yaml:"default-name"`   // 默认名称，用于替换"{{name}}"中内容
	Format      string     `yaml:"default-format"` // 文件格式，空的时候读取源文件扩展名格式
	Ratio       int        `yaml:"default-ratio"`  // 品质压缩率
	List        []IconSize `yaml:"list"`           // 尺寸
}

func (i *IconCfg) String() string {
	return fmt.Sprintf("{Format=%s, Ratio=%d, ListLen=%d}", i.Format, i.Ratio, len(i.List))
}

func NewIconContext(env string, cfg string, source string, target string,
	replaceName string, format string, ratio int) *IconContext {
	return &IconContext{envPath: env, cfgPath: cfg, source: source, target: target,
		replaceName: replaceName, format: format, ratio: ratio}
}

type IconContext struct {
	envPath string
	cfgPath string

	source string
	target string

	replaceName string

	format string
	ratio  int

	cfg *IconCfg
}

func (c *IconContext) String() string {
	return fmt.Sprintf("{Env=%s, Cfg=%s, Src=%s, TarDir=%s, Format=%s, Ratio=%d}",
		c.envPath, c.cfgPath, c.source, c.target, c.format, c.ratio)
}

func (c *IconContext) Mode() ResizeMode {
	return ModeIcon
}

func (c *IconContext) EnvPath() string {
	return c.envPath
}

func (c *IconContext) CfgPath() string {
	return c.cfgPath
}

func (c *IconContext) Config() *IconCfg {
	return c.cfg
}

func (c *IconContext) Format() string {
	return c.format
}

func (c *IconContext) Ratio() int {
	return c.ratio
}

func (c *IconContext) Source() string {
	return c.source
}

func (c *IconContext) Target() string {
	return c.target
}

func (c *IconContext) GetOutPath(size IconSize, format string) string {
	extName := formatx.GetExtName(format)
	name := size.Name
	if "" != c.replaceName {
		name = strings.ReplaceAll(name, IconNameSubstitute, c.replaceName)
	} else {
		name = strings.ReplaceAll(name, IconNameSubstitute, c.cfg.DefaultName)
	}
	fileName := name + "." + extName
	return filex.Combine(c.target, fileName)
}

func (c *IconContext) InitContext() error {
	if err := c.initConfig(); nil != err {
		return err
	}
	if err := c.iniSource(); nil != err {
		return err
	}
	if err := c.iniTarget(); nil != err {
		return err
	}
	c.ratio = GetRatio(c.ratio, c.cfg.Ratio)
	return nil
}

func (c *IconContext) initConfig() error {
	path := c.cfgPath
	if !filex.IsFile(c.cfgPath) {
		path = filex.Combine(c.envPath, c.cfgPath)
	}
	cfg := &IconCfg{}
	err := lib.UnmarshalYamlData(path, cfg)
	c.cfg = cfg
	if nil != err {
		return err
	}
	return nil
}

func (c *IconContext) iniSource() error {
	if c.source == "" {
		return errors.New(fmt.Sprintf("Mode[icon] src lack! "))
	}
	source, err := handleSourcePath(c.source, c.envPath)
	if nil != err {
		return errors.New(fmt.Sprintf("Mode[icon] %s", err))
	}
	c.source = source
	return nil
}

func (c *IconContext) iniTarget() error {
	if c.source == "" {
		return errors.New(fmt.Sprintf("Mode[icon] tar lack! "))
	}
	if filex.IsDir(c.target) {
		return nil
	}
	c.target = filex.Combine(c.envPath, c.target)
	return nil
}

func (c *IconContext) getRatio() int {
	return GetRatio(c.ratio, c.cfg.Ratio)
}
