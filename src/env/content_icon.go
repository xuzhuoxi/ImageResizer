package env

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ImageResizer/src/util"
	"github.com/xuzhuoxi/infra-go/filex"
	"strings"
)

type IconSize struct {
	Name   string `yaml:"name"` // 文件路径
	Width  int    `yaml:"w"`    // 宽度
	Height int    `yaml:"h"`    // 高度
}

type IconCfg struct {
	Source string     `yaml:"src"`    // 源文件路径
	Target string     `yaml:"tar"`    // 目标文件目录
	Format string     `yaml:"format"` // 文件格式，空的时候读取源文件扩展名格式
	Ratio  int        `yaml:"ratio"`  // 品质压缩率
	Size   []IconSize `yaml:"size"`   // 尺寸
}

func (i *IconCfg) updateSource(newSrc string) {
	i.Source = newSrc
}

func (i *IconCfg) updateTarget(newTar string) {
	i.Target = newTar
}

func (i *IconCfg) updateFormat(format string, ratio int) {
	i.Format, i.Ratio = format, ratio
}

func (i *IconCfg) GetOutPath(size IconSize) string {
	return filex.Combine(i.Target, size.Name, ".", i.Format)
}

func NewIconContext(env, cfg string, source, target string, format string, ratio int) *IconContext {
	return &IconContext{envPath: env, cfgPath: cfg, source: source, target: target, format: format, ratio: ratio}
}

type IconContext struct {
	envPath string
	cfgPath string

	source string
	target string
	format string
	ratio  int

	cfg *IconCfg
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

func (c *IconContext) GetOutPath(size IconSize) string {
	return c.cfg.GetOutPath(size)
}

func (c *IconContext) InitContext() error {
	path := c.cfgPath
	if !filex.IsFile(c.cfgPath) {
		path = filex.Combine(c.envPath, c.cfgPath)
	}
	cfg := &IconCfg{}
	err := util.UnmarshalYamlData(path, cfg)
	c.cfg = cfg
	if nil != err {
		return err
	}
	if c.source != "" {
		if filex.IsFile(c.source) {
			cfg.updateSource(c.source)
		} else {
			newSource := filex.Combine(c.envPath, c.source)
			if !filex.IsFile(newSource) {
				return errors.New(fmt.Sprintf("Mode[icon] source[%s] is not exist! ", c.source))
			}
			cfg.updateSource(newSource)
		}
	}
	if c.target != "" {
		if filex.IsFile(c.target) {
			cfg.updateSource(c.target)
		} else {
			cfg.updateSource(filex.Combine(c.envPath, c.target))
		}
	}
	if c.format != "" {
		cfg.updateFormat(c.format, c.getRatio())
	} else {
		if cfg.Format == "" {
			ext := filex.GetExtWithoutDot(cfg.Source)
			if ext == "" {
				return errors.New(fmt.Sprintf("Mode[icon] format not exist anywhere! "))
			}
			cfg.updateFormat(strings.ToUpper(ext), c.getRatio())
		}
	}
	return nil
}

func (c *IconContext) getRatio() int {
	return GetRatio(c.ratio, c.cfg.Ratio)
}
