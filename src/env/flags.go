package env

import (
	"flag"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/osxu"
	"strings"
)

type CmdFlags struct {
	EnvPath    string // [可选]环境目录,其它目录如果使用相对路径，则以这里为基准
	Mode       string // [*]运行模式
	Include    string // 要处理的文件扩展名
	CfgPath    string // [icon*]icon模式下配置文件路径
	Source     string // [scale*|icon*]字符串路径，文件夹或文件,"./"开头视为相对路径
	TargetDir  string // [scale*|icon*]字符串目录路径
	TargetFile string // [scale*|icon*]字符串文件路径
	Size       string // [size*]输出固定大小,格式：[整数/宽x高],...
	Scale      string // [scale*]缩放比例,格式：[整数/宽x高],...
	Format     string // [可选]输出文件格式,图像格式[pngx,jpeg,gifx,jpg]
	Ratio      int    // [可选]输出文件压缩比，整数(0,100]
}

func (f *CmdFlags) GetContexts() (iconCtx *IconContext, sizeCtc *SizeContext, scaleCtx *ScaleContext, err error) {
	if ModeIcon == f.Mode {
		ctx := NewIconContext(f.EnvPath, f.CfgPath, f.Source, f.TargetDir, f.Format, f.Ratio)
		err = ctx.InitContext()
		if nil != err {
			return
		}
		iconCtx = ctx
		return
	}
	target, oneByOne := f.getTarget()
	if ModeSize == f.Mode {
		ctx := NewSizeContext(f.EnvPath, f.Include, f.Source, target, f.Size, oneByOne, f.Format, f.Ratio)
		err = ctx.InitContext()
		if nil != err {
			return
		}
		sizeCtc = ctx
		return
	}
	if ModeScale == f.Mode {
		ctx := NewScaleContext(f.EnvPath, f.Include, f.Source, target, f.Scale, oneByOne, f.Format, f.Ratio)
		err = ctx.InitContext()
		if nil != err {
			return
		}
		scaleCtx = ctx
		return
	}
	return
}

func (f *CmdFlags) getTarget() (target string, oneByOne bool) {
	oneByOne = f.TargetFile != ""
	if oneByOne {
		target = f.TargetFile
	} else {
		target = f.TargetDir
	}
	return
}

func ParseFlags() *CmdFlags {
	// 空[执行文件目录]
	// 绝对路径
	// 相对于执行目录的相对路径
	envPath := flag.String("env", "", "Running Environment Path! ")
	mode := flag.String("mode", "", "Running Mode! ")
	include := flag.String("include", "", "Included Formats! ")

	cfgPath := flag.String("cfg", "", "Mode[icon] config path! ")
	source := flag.String("src", "", "Mode[scale|size] Source paths! ")
	targetDir := flag.String("tar_dir", "", "Mode[scale|size] Target paths! ")
	targetFile := flag.String("tar_file", "", "Mode[scale|size] Target paths! ")

	size := flag.String("size", "", "Mode[size] Size! ")
	scale := flag.String("scale", "", "Mode[scale] Scale! ")

	format := flag.String("format", "", "Mode[icon|scale|size] Format! ")
	ratio := flag.Int("ratio", 0, "Mode[icon|scale|size] Ratio! ")

	flag.Parse()

	env := getEnvPath(*envPath)

	return &CmdFlags{
		EnvPath: env, Mode: strings.ToLower(*mode), Include: strings.ToLower(*include),
		CfgPath: *cfgPath, Source: *source, TargetDir: *targetDir, TargetFile: *targetFile,
		Size: strings.ToLower(*size), Scale: *scale, Format: strings.ToLower(*format), Ratio: *ratio}
}

func getEnvPath(env string) string {
	runningRoot := osxu.GetRunningDir()
	if "" == env {
		env = runningRoot
	} else if !filex.IsDir(env) {
		env = filex.Combine(runningRoot, env)
	}
	return env
}
