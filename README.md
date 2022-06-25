# ImageResizer

ImageResizer 主要用于生成不同尺寸的图像。

中文 | [English](/README_EN.md)

## <span id="p1">兼容性
go1.16

## <span id="p2">如何开始

你可以选择[下载发行版本](#p2.1)或者[构造](#p2.2)获得执行文件。

### <span id="p2.1">下载发行版本

- 到以下地址下载: [https://github.com/xuzhuoxi/ImageResizer/releases](https://github.com/xuzhuoxi/ImageResizer/releases).

### <span id="p2.2">构造

- 下载仓库

	```sh
	go get -u github.com/xuzhuoxi/ImageResizer
	```

- 构造

  + 构造依赖到第三方库[goxc](https://github.com/laher/goxc)。
  + 如有必要，你可以修改相关构造脚本。
  + 建议先关闭gomod：`go env -w GO111MODULE=off`，由于goxc已经比较旧。
  + 执行构造脚本([goxc_build.sh](/build/goxc_build.sh)或([goxc_build.bat](/build/goxc_build.bat),执行文件将生成在[release](/build/release)目录中。

## <span id="p3">运行

工具仅支持命令行执行。

### <span id="p3.1">命令行参数说明

  - -env 
    + 【**可选**】运行时环境路径，支持绝对路径与相对于当前执行目录的相对路径，空表示使用执行文件所在目录
    + 例如: 
      -env=D:/workspaces
  - -mode
    + 【**必要**】执行模式[icon|size|scale]
    + 例如: 
      -mode=icon
  - -cfg
    + 【**icon必要**】配置文件路径
    + 例如: 
      -cfg=D:/workspaces/icon_ios.yaml
  - -src
    + 【**必要**】来源文件或目录，支持多个，可用英文逗号","分隔
    + 例如: 
      -src=D:/workspaces/IconDir,D:/workspaces/Icon.png
  - -include
    + 【**size,scale可选**】当来源包含目录时必要，扩展名过滤，支持多个，可用英文逗号","分隔
    + 例如:
      -include=jpg,png
  - -tar_dir
    + 【**size,scale:File与Dir二选一**】【**icon必要**】产出目录，不支持多个
    + 例如:
      -tar_dir=D:/workspaces/OutDir
  - -tar_file
    + 【**size,scale:File与Dir二选一**】产出文件
    + 例如:
      -tar_file=D:/workspaces/OutDir/IconNew.png
  - -size
    + 【**size必要**】产出文件尺寸，支持多个，可用英文逗号","分隔
    + 例如:
      -size=512,48x32
  - -scale
    + 【**scale必要**】产出文件比例，支持多个，可用英文逗号","分隔
    + 例如:
      -scale=0.8,1,1.5
  - -name
    + 【**icon可选**】模式下替换名称
    + 例如:
      -name=IconNew 
  - -format
    + 【**可选**】指定产出文件格式
    + 例如:
      -format=jpg
  - -ratio
    + 【**可选**】指定产出文件质量
    + 例如:
      -ratio=65

### <span id="p3.2">应用场景举例

**注意**：以下使用`$EnvPath`代替实现环境路径

- 模式"icon"应用场景：

  + icon配置文件说明：
  
	1. 配置文件使用ymal格式。
	2. 配置文件说明及相应结构：
		```golang
		type IconSize struct {
			Name string `yaml:"name"` // 文件路径
			Size string `yaml:"size"` // 尺寸,格式: 长x宽
		}
		type IconCfg struct {
			DefaultName string     `yaml:"default-name"`   // 默认名称，用于替换"{{name}}"中内容
			Format      string     `yaml:"default-format"` // 文件格式，空的时候读取源文件扩展名格式
			Ratio       int        `yaml:"default-ratio"`  // 品质压缩率
			List        []IconSize `yaml:"list"`           // 尺寸
		}
		```
	  - default-name: 在未使用-name参数时默认使用的名称
	  - default-format: 在未使用-format参数时默认使用的格式参数，填空字符串时表示使用源图像的格式
	  - default-ratio: 在未使用-ratio参数时默认使用的品质参数，填0时表示使用工具默认的参数85.
	  - name: 文件路径，**不用填扩展名**，支持“{{name}}”的替换参数。
	  - size: 图像尺寸，格式:长x宽
	3. 例子可参考[icon_ios.yaml](/demo/icon_ios.yaml)

  + 生成iOS应用图标,并指定文件名称前缀

	命令行：
    `ImageResizer -env=$EnvPath -mode=icon -cfg=icon_ios.yaml -src=src/SrcIcon.png -tar_dir=tar -name=AppIcon`

- 模式"size"应用场景：

  + 调整指定图像大小,并指定格式与质量

    命令行：
    `ImageResizer -env=$EnvPath -mode=size -src=src/SrcIcon.png -tar_file=tar/SrcIcon.png -size=512x512 -format=jpg -ratio=65`

  + 批量调整目录中的png图像到多种大小

    命令行：
    `ImageResizer -env=$EnvPath -mode=size -src=src -include=png -tar_dir=tar, -size=512,256x256`

- 模式"scale"应用场景：

  + 调整指定图像比例

    命令行：
    `ImageResizer -env=$EnvPath -mode=scale -src=src/SrcIcon.png -tar_file=tar/SrcIcon.png -scale=0.5`

  + 批量调整目录中的jpg图像到多种比例,并指定格式与质量

    命令行：
    `ImageResizer -env=$EnvPath -mode=scale -src=src -include=png -tar_dir=tar, -scale=0.6,1.2 -foramt=png -ratio=85`

### <span id="p3.3">例子

- 例子目录位于[demo](/demo).
- Win64平台可执行[run_icon.bat](/demo/run_icon.bat),[run_scale.bat](/demo/run_scale.bat),[run_size.bat](/demo/run_size.bat)进行测试。
- Mac平台可执行[run_icon.sh](/demo/run_icon.sh),[run_scale.sh](/demo/run_scale.sh),[run_size.sh](/demo/run_size.sh)进行测试。
- Linux平台修改Mac测试脚本中的执行文件路径进行测试。

  [命令行参数说明](#p3.1)

## <span id="p4">依赖库

- infra-go [https://github.com/xuzhuoxi/infra-go](https://github.com/xuzhuoxi/infra-go)

- goxc [https://github.com/laher/goxc](https://github.com/laher/goxc) 

## <span id="p5">联系作者

xuzhuoxi 

<xuzhuoxi@gmail.com> or <mailxuzhuoxi@163.com>

## <span id="p6">License
ImageResizer source code is available under the MIT [License](/LICENSE).


