package env

import (
	"flag"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/osxu"
	"strings"
)

type CmdFlags struct {
	EnvPath string // 【可选】运行时环境路径，支持绝对路径与相对于当前执行目录的相对路径，空表示使用执行文件所在目录
	Mode    string // 【必要】执行模式[icon|size|scale]

	CfgPath       string // 【icon必要】配置文件路径
	Source        string // 【必要】来源文件或目录，支持多个，可用英文逗号","分隔
	SourceInclude string // 【size,scale可选】当来源包含目录时必要，扩展名过滤，支持多个，可用英文逗号","分隔
	TargetDir     string // 【size,scale:File与Dir二选一】【icon必要】产出目录，不支持多个
	TargetFile    string // 【size,scale:File与Dir二选一】产出文件
	TargetSize    string // 【size必要】产出文件尺寸，支持多个，可用英文逗号","分隔，格式：[整数/宽x高],...
	TargetScale   string // 【scale必要】产出文件比例，支持多个，可用英文逗号","分隔

	ReplaceName string // 【icon可选】模式下替换名称
	Format      string // 【可选】指定产出文件格式
	Ratio       int    // 【可选】指定产出文件质量,整数(0,100]
}

func (f *CmdFlags) GetContexts() (iconCtx *IconContext, sizeCtc *SizeContext, scaleCtx *ScaleContext, err error) {
	evnPath, _ := f.getEnvPath()
	if ModeIcon == f.Mode {
		ctx := NewIconContext(evnPath, f.CfgPath, f.Source, f.TargetDir,
			f.ReplaceName, f.Format, f.Ratio)
		err = ctx.InitContext()
		if nil != err {
			return
		}
		iconCtx = ctx
		return
	}
	target, oneByOne := f.getTarget()
	if ModeSize == f.Mode {
		ctx := NewSizeContext(evnPath, oneByOne,
			f.Source, f.SourceInclude, target, f.TargetSize,
			f.Format, f.Ratio)
		err = ctx.InitContext()
		if nil != err {
			return
		}
		sizeCtc = ctx
		return
	}
	if ModeScale == f.Mode {
		ctx := NewScaleContext(evnPath, oneByOne,
			f.Source, f.SourceInclude, target, f.TargetScale,
			f.Format, f.Ratio)
		err = ctx.InitContext()
		if nil != err {
			return
		}
		scaleCtx = ctx
		return
	}
	return
}

func (f *CmdFlags) getEnvPath() (evnPath string, isDefault bool) {
	runningRoot := osxu.GetRunningDir()
	if "" == f.EnvPath {
		return runningRoot, true
	}
	if filex.IsDir(f.EnvPath) {
		return f.EnvPath, false
	}
	return filex.Combine(runningRoot, f.EnvPath), false
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
	// 【可选】运行时环境路径，支持绝对路径与相对于当前执行目录的相对路径，空表示使用执行文件所在目录
	envPath := flag.String("env", "", "Running Environment Path! ")
	// 【必要】执行模式[icon|size|scale]
	mode := flag.String("mode", "", "Running Mode! ")

	// 【icon必要】配置文件路径
	cfgPath := flag.String("cfg", "", "Mode[icon] config path! ")
	// 【必要】来源文件或目录，支持多个，可用英文逗号","分隔
	source := flag.String("src", "", "Mode[scale|size] Source paths! ")
	// 【size,scale可选】当来源包含目录时必要，扩展名过滤，支持多个，可用英文逗号","分隔
	sourceInclude := flag.String("include", "", "Included Formats! ")
	// 【size,scale:File与Dir二选一】【icon必要】产出目录，不支持多个
	targetDir := flag.String("tar_dir", "", "Mode[scale|size] Target paths! ")
	// 【size,scale:File与Dir二选一】产出文件
	targetFile := flag.String("tar_file", "", "Mode[scale|size] Target paths! ")

	// 【size必要】产出文件尺寸，支持多个，可用英文逗号","分隔
	targetSize := flag.String("size", "", "Mode[size] Size! ")
	// 【scale必要】产出文件比例，支持多个，可用英文逗号","分隔
	targetScale := flag.String("scale", "", "Mode[scale] Scale! ")

	// 【icon可选】模式下替换名称
	replaceName := flag.String("name", "", "Mode[icon] Replace Name! ")
	// 【可选】指定产出文件格式
	format := flag.String("format", "", "Mode[icon|scale|size] Format! ")
	// 【可选】指定产出文件质量
	ratio := flag.Int("ratio", 0, "Mode[icon|scale|size] Ratio! ")

	flag.Parse()

	return &CmdFlags{
		EnvPath: *envPath, Mode: strings.ToLower(*mode),
		CfgPath: *cfgPath, Source: *source, SourceInclude: strings.ToLower(*sourceInclude),
		TargetDir: *targetDir, TargetFile: *targetFile, TargetSize: strings.ToLower(*targetSize), TargetScale: *targetScale,
		ReplaceName: *replaceName, Format: strings.ToLower(*format), Ratio: *ratio}
}
