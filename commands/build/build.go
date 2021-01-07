package build

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf-cli/library/mlog"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/genv"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gproc"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gutil"
	"regexp"
	"runtime"
	"strings"
)

// https://golang.google.cn/doc/install/source
// Here're the most commonly used platforms and arches,
// but some are removed:
//    android    arm
//    dragonfly amd64
//    plan9     386
//    plan9     amd64
//    solaris   amd64
const platforms = `
    darwin    amd64
    freebsd   386
    freebsd   amd64
    freebsd   arm
    linux     386
    linux     amd64
    linux     arm
    linux     arm64
    linux     ppc64
    linux     ppc64le
    linux     mips
    linux     mipsle
    linux     mips64
    linux     mips64le
    netbsd    386
    netbsd    amd64
    netbsd    arm
    openbsd   386
    openbsd   amd64
    openbsd   arm
    windows   386
    windows   amd64
`

const (
	nodeNameInConfigFile = "gfcli.build" // nodeNameInConfigFile is the node name for compiler configurations in configuration file.
	packedGoFileName     = "data.go"     // packedGoFileName specifies the file name for packing common folders into one single go file.
)

func Help() {
	mlog.Print(gstr.TrimLeft(`
用法 
    gf build 文件 [选项]

主题
    文件  构建文件路径。

选项
    -n, --name       生成的可执行文件名称。如果是windows平台，那么默认会加上.exe后缀。
    -v, --version    程序版本，如果指定版本信息，那么程序生成的路径中会多一层以版本名称的目录。
    -a, --arch       编译架构，多个以,号分隔，如果是all表示编译所有支持架构。
    -s, --system     编译平台，多个以,号分隔，如果是all表示编译所有支持平台。
    -o, --output     输出的可执行文件路径，当该参数指定时，name和path参数失效，常用于编译单个可执行文件。
    -p, --path       编译可执行文件存储的目录地址，默认是'./bin'（同时指定 -a -s 参数时生效）。
    -e, --extra      额外自定义的编译参数，会直接传递给go build命令。
    -m, --mod        同go build -mod编译选项，使用"-m none"禁用go module。
    -c, --cgo        是否开启cgo特性，默认是关闭的。如果开启，那么交叉编译可能会有问题。
    --CC             开启交叉编译，参数CC（linux推荐CC=x86_64-linux-musl-gcc，windows推荐CC=x86_64-w64-mingw32-gcc）。
    --CGO_LDFLAGS    开启交叉编译，参数CGO_LDFLAGS（linux推荐CGO_LDFLAGS="-static"，windows无此参数）。
    --pack           go打包前，将需要打包的目录，多个以,号分隔，生成到packed/data.go。
    --swagger        go打包前，自动解析并将swagger打包到pack/swagger中。

示例
    gf build main.go
    gf build main.go --swagger
    gf build main.go --pack public,template
    gf build main.go --cgo
    gf build main.go -m none 
    gf build main.go -n my-app -a all -s all
    gf build main.go -n my-app -a amd64,386 -s linux -p .
    gf build main.go -n my-app -v 1.0 -a amd64,386 -s linux,windows,darwin -p ./docker/bin

开启cgo特性，交叉编译示例（注意需要提前安装相关依赖，由于参数不同，需要分开编译）

    1、安装交叉编译工具（mac示例）
    brew install mingw-w64
    brew install FiloSottile/musl-cross/musl-cross

    2、linux_amd64环境
    gfctl build main.go --cgo --CC=x86_64-linux-musl-gcc --CGO_LDFLAGS="-static" --name gfctl --arch amd64 --system linux --version 1.15.0
    
    3、windows_amd64环境
    gfctl build main.go --cgo --CC=x86_64-w64-mingw32-gcc --name gfctl --arch amd64 --system windows --version 1.15.0
    
    4、mac_amd64环境
    gfctl build main.go --cgo --name gfctl --arch amd64 --system darwin --version 1.15.0

说明
    "build"命令是最常用的命令，它被设计成为一个强大的"go build"命令的包装器，方便交叉编译使用。
    它为构建二进制文件提供了更多功能：
    1. 适用于多种平台和跨平台的编译。
    2. 配置文件支持编译。
    3. 内置变量。

平台
    darwin    amd64
    freebsd   386,amd64,arm
    linux     386,amd64,arm,arm64,ppc64,ppc64le,mips,mipsle,mips64,mips64le
    netbsd    386,amd64,arm
    openbsd   386,amd64,arm
    windows   386,amd64
`))
}

func Run() {
	mlog.SetHeaderPrint(true)
	parser, err := gcmd.Parse(g.MapStrBool{
		"n,name":      true,
		"v,version":   true,
		"a,arch":      true,
		"s,system":    true,
		"o,output":    true,
		"p,path":      true,
		"e,extra":     true,
		"m,mod":       true,
		"pack":        true,
		"c,cgo":       false,
		"CC":          false, // cgo交叉编译环境变量CC
		"CGO_LDFLAGS": false, // cgo交叉编译环境变量CGO_LDFLAGS
		"swagger":     false,
	})
	if err != nil {
		mlog.Fatal(err)
	}
	file := parser.GetArg(2)
	if len(file) < 1 {
		// Check and use the main.go file.
		if gfile.Exists("main.go") {
			file = "main.go"
		} else {
			mlog.Fatal("构建文件路径不能为空")
		}
	}
	path := getOption(parser, "path", "./bin")
	name := getOption(parser, "name", gfile.Name(file))
	if len(name) < 1 || name == "*" {
		mlog.Fatal("name不能为空")
	}
	var (
		mod   = getOption(parser, "mod")
		extra = getOption(parser, "extra")
	)
	if mod != "" && mod != "none" {
		mlog.Debugf(`mod 是 %s`, mod)
		if extra == "" {
			extra = fmt.Sprintf(`-mod=%s`, mod)
		} else {
			extra = fmt.Sprintf(`-mod=%s %s`, mod, extra)
		}
	}
	var (
		cgoEnabled   = gconv.Bool(getOption(parser, "cgo"))
		CC           = getOption(parser, "CC")
		cgoLDFLAGS   = getOption(parser, "CGO_LDFLAGS")
		version      = getOption(parser, "version")
		outputPath   = getOption(parser, "output")
		archOption   = getOption(parser, "arch")
		systemOption = getOption(parser, "system")
		packStr      = getOption(parser, "pack")
		arches       = strings.Split(archOption, ",")
		systems      = strings.Split(systemOption, ",")
	)
	if !cgoEnabled {
		cgoEnabled = parser.ContainsOpt("cgo")
	}
	if len(version) > 0 {
		path += "/" + version
	}

	// Auto swagger.
	if containsOption(parser, "swagger") {
		if err := gproc.ShellRun(`gf swagger`); err != nil {
			return
		}
		if gfile.Exists("swagger") {
			packCmd := fmt.Sprintf(`gf pack %s packed/%s`, "swagger", packedGoFileName)
			mlog.Print(packCmd)
			if err := gproc.ShellRun(packCmd); err != nil {
				return
			}
		}
	}

	// Auto packing.
	if len(packStr) > 0 {
		packCmd := fmt.Sprintf(`gf pack %s packed/%s`, packStr, packedGoFileName)
		mlog.Print(packCmd)
		gproc.ShellRun(packCmd)
	}

	// Injected information by building flags.
	ldFlags := fmt.Sprintf(`-X 'github.com/gogf/gf/os/gbuild.builtInVarStr=%v'`, getBuildInVarStr())

	// start building
	mlog.Print("开始构建...")
	if cgoEnabled {
		// 开启交叉编译
		genv.Set("CGO_ENABLED", "1")

		// 交叉编译参数CC
		if CC != "" {
			genv.Set("CC", CC)
		}
		// 交叉编译参数CGO_LDFLAGS
		if cgoLDFLAGS != "" {
			genv.Set("CGO_LDFLAGS", cgoLDFLAGS)
		}

	} else {
		genv.Set("CGO_ENABLED", "0")
	}
	var (
		cmd   = ""
		ext   = ""
		reg   = regexp.MustCompile(`\s+`)
		lines = strings.Split(strings.TrimSpace(platforms), "\n")
	)
	for _, line := range lines {
		cmd = ""
		ext = ""
		line = strings.TrimSpace(line)
		line = reg.ReplaceAllString(line, " ")
		array := strings.Split(line, " ")
		array[0] = strings.TrimSpace(array[0])
		array[1] = strings.TrimSpace(array[1])
		if len(systems) > 0 && systems[0] != "" && systems[0] != "all" && !gstr.InArray(systems, array[0]) {
			continue
		}
		if len(arches) > 0 && arches[0] != "" && arches[0] != "all" && !gstr.InArray(arches, array[1]) {
			continue
		}
		if len(systemOption) == 0 && len(archOption) == 0 {
			if runtime.GOOS == "windows" {
				ext = ".exe"
			}
			// Single binary building, output the binary to current working folder.
			output := ""
			if len(outputPath) > 0 {
				output = "-o " + outputPath + ext
			} else {
				output = "-o " + name + ext
			}
			cmd = fmt.Sprintf(`go build %s -ldflags "%s" %s %s`, output, ldFlags, extra, file)
		} else {
			// Cross-building, output the compiled binary to specified path.
			if array[0] == "windows" {
				ext = ".exe"
			}
			genv.Set("GOOS", array[0])
			genv.Set("GOARCH", array[1])
			cmd = fmt.Sprintf(
				`go build -o %s/%s/%s%s -ldflags "%s" %s %s`,
				path, array[0]+"_"+array[1], name, ext, ldFlags, extra, file,
			)
		}
		// It's not necessary printing the complete command string.
		cmdShow, _ := gregex.ReplaceString(`\s+(-ldflags ".+?")\s+`, " ", cmd)
		mlog.Print(cmdShow)
		if result, err := gproc.ShellExec(cmd); err != nil {
			mlog.Fatalf("构建失败：%s%s", result, err.Error())
		}
		// single binary building.
		if len(systemOption) == 0 && len(archOption) == 0 {
			break
		}
	}
	mlog.Print("完成!")
}

// getOption retrieves option value from parser and configuration file.
// It returns the default value specified by parameter <value> is no value found.
func getOption(parser *gcmd.Parser, name string, value ...string) (result string) {
	result = parser.GetOpt(name)
	if result == "" && g.Config().Available() {
		result = g.Config().GetString(nodeNameInConfigFile + "." + name)
	}
	if result == "" && len(value) > 0 {
		result = value[0]
	}
	return
}

// containsOption checks whether the command option or the configuration file containing
// given option name.
func containsOption(parser *gcmd.Parser, name string) bool {
	result := parser.ContainsOpt(name)
	if !result && g.Config().Available() {
		result = g.Config().Contains(nodeNameInConfigFile + "." + name)
	}
	return result
}

// getBuildInVarMapJson retrieves and returns the custom build-in variables in configuration
// file as json.
func getBuildInVarStr() string {
	buildInVarMap := g.Map{}
	if g.Config().Available() {
		configMap := g.Config().GetMap(nodeNameInConfigFile)
		if len(configMap) > 0 {
			_, v := gutil.MapPossibleItemByKey(configMap, "VarMap")
			if v != nil {
				buildInVarMap = gconv.Map(v)
			}
		}
	}
	buildInVarMap["builtGit"] = getGitCommit()
	buildInVarMap["builtTime"] = gtime.Now().String()
	b, err := json.Marshal(buildInVarMap)
	if err != nil {
		mlog.Fatal(err)
	}
	return gbase64.EncodeToString(b)
}

// getGitCommit retrieves and returns the latest git commit hash string if present.
func getGitCommit() string {
	if gproc.SearchBinary("git") == "" {
		return ""
	}
	if s, _ := gproc.ShellExec("git rev-list -1 HEAD"); s != "" {
		if !gstr.Contains(s, " ") && !gstr.Contains(s, "fatal") {
			return gstr.Trim(s)
		}
	}
	return ""
}
