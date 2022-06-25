# ImageResizer

ImageResizer is mainly used to generate images of different sizes.

[中文](/README.md) | English

## <span id="p1">Compatibility
go1.16

## <span id="p2">How To Get Started

You can choose [Download Release Version](#p2.1) or [Construction](#p2.2) to get the executable file.

### <span id="p2.1">Download Release Version

- Download from the following address: [https://github.com/xuzhuoxi/ImageResizer/releases](https://github.com/xuzhuoxi/ImageResizer/releases).

### <span id="p2.2">Build

- Download repository

	```sh
	go get -u github.com/xuzhuoxi/ImageResizer
	```

- Build

  + The construction depends on the third-party library [goxc](https://github.com/laher/goxc).

  + If necessary, you can modify the relevant construction scripts.

  + It is recommended to turn off gomod first: `go env -w GO111MODULE=off`, because goxc is old.

  + Execute the build script [goxc_build.sh](/build/goxc_build.sh) or [goxc_build.bat](/build/goxc_build.bat), the executable file will be generated in the "build/release" directory .

## <span id="p3">Run

The tool supports command line execution only.

### <span id="p3.1">Command Line Parameter Description

  - -env 
    + [**Optional**] Runtime environment path, supports absolute path and relative path relative to the current execution directory, empty means to use the directory where the execution file is located
    + For example: 
      -env=D:/workspaces
  - -mode
    +  [**Required**] Execution mode [icon|size|scale]
    + For example: 
      -mode=icon
  - -cfg
    +  [**icon necessary**] configuration file path
    + For example: 
      -cfg=D:/workspaces/icon_ios.yaml
  - -src
    + [**Required**] Source files or directories, multiple are supported, separated by commas ","
    + For example: 
      -src=D:/workspaces/IconDir,D:/workspaces/Icon.png
  - -include
    + [**size, scale optional**] Necessary when the source contains a directory, the extension is filtered, multiple, and can be separated by an English comma ","
    + For example:
      -include=jpg,png
  - -tar_dir
    + [**size, scale: File and Dir choose one**] **[icon necessary**] Output directory, does not support multiple
    + For example:
      -tar_dir=D:/workspaces/OutDir
  - -tar_file
    + [**size, scale: File and Dir choose one**] output file
    + For example:
      -tar_file=D:/workspaces/OutDir/IconNew.png
  - -size
    + [**size necessary**] output file size, support multiple, can be separated by English comma ","
    + For example:
      -size=512,48x32
  - -scale
    +  [**scale necessary**] output file scale, support multiple, can be separated by English comma ","
    + For example:
      -scale=0.8,1,1.5
  - -name
    +  [**icon optional**] replace name in mode
    + For example:
      -name=IconNew 
  - -format
    +  [**Optional**] Specify the output file format
    + For example:
      -format=jpg
  - -ratio
    + [**Optional**] Specify the output file quality
    + For example:
      -ratio=65

### <span id="p3.2">Example Of Application Scenarios

**Note**: The following uses `$EnvPath` instead of the actual environment path.

- Mode "icon" application scenarios:

  + icon profile description:
  
	1. The configuration file is in ymal format.

	2. Configuration file description and corresponding structure:

		```golang
		type IconSize struct {
			Name string `yaml:"name"` // file path
			Size string `yaml:"size"` // size, format: "length"x"width"
		}
		type IconCfg struct {
			DefaultName string     `yaml:"default-name"`    // default name, used to replace the content in "{{name}}"
			Format      string     `yaml:"default-format"`  // file format, read the source file extension format when empty
			Ratio       int        `yaml:"default-ratio"`   // quality compression ratio
			List        []IconSize `yaml:"list"`            // size
		}
		```

	  - default-name: the default name to use when the -name parameter is not used

	  - default-format: The format parameter used by default when the -format parameter is not used, and the format of the source image is used when the empty string is filled

	  - default-ratio: The quality parameter used by default when the -ratio parameter is not used. Filling in 0 means using the tool's default parameter of 85.

	  - name: file path, **do not need to fill in the extension**, supports the replacement parameter of "{{name}}".

	  - size: image size, format: "length"x"width"

	3. For example, please refer to [icon_ios.yaml](/demo/icon_ios.yaml)

  + Generate iOS app icon and specify file name prefix

	 Command Line：
    `ImageResizer -env=$EnvPath -mode=icon -cfg=icon_ios.yaml -src=src/SrcIcon.png -tar_dir=tar -name=AppIcon`

- Mode "size" application scenarios:

  + Resize the specified image, and specify the format and quality

     Command Line：
    `ImageResizer -env=$EnvPath -mode=size -src=src/SrcIcon.png -tar_file=tar/SrcIcon.png -size=512x512 -format=jpg -ratio=65`

  + Batch resize png images in a directory to multiple sizes

     Command Line：
    `ImageResizer -env=$EnvPath -mode=size -src=src -include=png -tar_dir=tar, -size=512,256x256`

- Mode "scale" application scenarios:

  + Adjust the specified image ratio

     Command Line：
    `ImageResizer -env=$EnvPath -mode=scale -src=src/SrcIcon.png -tar_file=tar/SrcIcon.png -scale=0.5`

  + Batch resize jpg images in catalogs to multiple scales, and specify format and quality

     Command Line：
    `ImageResizer -env=$EnvPath -mode=scale -src=src -include=png -tar_dir=tar, -scale=0.6,1.2 -foramt=png -ratio=85`

### <span id="p3.3">Example

- The example directory is located at [demo](/demo).
- The Win64 platform can execute [run_icon.bat](/demo/run_icon.bat), [run_scale.bat](/demo/run_scale.bat), [run_size.bat](/demo/run_size.bat) for testing.
- Mac platform can execute [run_icon.sh](/demo/run_icon.sh), [run_scale.sh](/demo/run_scale.sh), [run_size.sh](/demo/run_size.sh) for testing.
- Modify the execution file path in the Mac test script for testing on the Linux platform.

  [Command line parameter description](#p3.1)

## <span id="p4">Dependency Library

- infra-go [https://github.com/xuzhuoxi/infra-go](https://github.com/xuzhuoxi/infra-go)

- goxc [https://github.com/laher/goxc](https://github.com/laher/goxc) 

## <span id="p5">Contact

xuzhuoxi 

<xuzhuoxi@gmail.com> or <mailxuzhuoxi@163.com>

## <span id="p6">License
ImageResizer source code is available under the MIT [License](/LICENSE).


