package main

import (
	"fmt"
	"github.com/gogf/gf-cli/commands/env"
	"github.com/gogf/gf-cli/commands/mod"
	"github.com/gogf/gf/errors/gerror"
	"strings"

	_ "github.com/gogf/gf-cli/boot"
	"github.com/gogf/gf-cli/commands/build"
	"github.com/gogf/gf-cli/commands/docker"
	"github.com/gogf/gf-cli/commands/fix"
	"github.com/gogf/gf-cli/commands/gen"
	"github.com/gogf/gf-cli/commands/get"
	"github.com/gogf/gf-cli/commands/initialize"
	"github.com/gogf/gf-cli/commands/install"
	"github.com/gogf/gf-cli/commands/pack"
	"github.com/gogf/gf-cli/commands/run"
	"github.com/gogf/gf-cli/commands/swagger"
	"github.com/gogf/gf-cli/commands/update"
	"github.com/gogf/gf-cli/library/allyes"
	"github.com/gogf/gf-cli/library/mlog"
	"github.com/gogf/gf-cli/library/proxy"
	"github.com/gogf/gf/os/gbuild"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
)

const (
	VERSION = "v1.15.0"
)

func init() {
	// Automatically sets the golang proxy for all commands.
	proxy.AutoSet()
}

var (
	helpContent = gstr.TrimLeft(`
用法
    gf 命令 [主题] [选项]

命令
    env        显示当前Golang环境变量
    get        默认情况下将GF安装或更新到系统...
    gen        自动为ORM模型生成go文件...
    mod        Go模块的额外功能...
    run        运行具有类似热编译功能的go代码...
    init       创建并初始化一个空的GF项目...
    help       显示有关指定命令的详细信息
    pack       将任何文件/目录打包到资源文件或go文件中...
    build      交叉编译跨平台的go项目...
    docker     为当前的GF项目创建一个docker镜像...
    swagger    当前项目的swagger功能...
    update     将当前gf二进制文件更新为最新的（可能需要root/admin权限）
    install    将gf二进制文件安装到系统（可能需要root/admin权限）
    version    显示当前二进制版本信息

选项
    -y         所有命令都是"yes"，没有提示询问
    -?,-h      显示指定命令的帮助或详细信息
    -v,-i      显示版本信息

附加
    使用"gf help 命令"或"gf 命令 -h"获取（注释末尾带有"..."）命令的详细信息，
`)
)

func main() {
	defer func() {
		if exception := recover(); exception != nil {
			if err, ok := exception.(error); ok {
				mlog.Print(gerror.Current(err).Error())
			} else {
				panic(exception)
			}
		}
	}()

	allyes.Init()

	command := gcmd.GetArg(1)
	// Help information
	if gcmd.ContainsOpt("h") && command != "" {
		help(command)
		return
	}
	switch command {
	case "help":
		help(gcmd.GetArg(2))
	case "version":
		version()
	case "env":
		env.Run()
	case "get":
		get.Run()
	case "gen":
		gen.Run()
	case "fix":
		fix.Run()
	case "mod":
		mod.Run()
	case "init":
		initialize.Run()
	case "pack":
		pack.Run()
	case "docker":
		docker.Run()
	case "swagger":
		swagger.Run()
	case "update":
		update.Run()
	case "install":
		install.Run()
	case "build":
		build.Run()
	case "run":
		run.Run()
	default:
		for k := range gcmd.GetOptAll() {
			switch k {
			case "?", "h":
				mlog.Print(helpContent)
				return
			case "i", "v":
				version()
				return
			}
		}
		// No argument or option, do installation checks.
		if !install.IsInstalled() {
			mlog.Print("嗨，这似乎是你第一次安装gf cli。")
			s := gcmd.Scanf("你想要安装gf二进制到你的系统吗？[y/n]: ")
			if strings.EqualFold(s, "y") {
				install.Run()
				gcmd.Scan("按<Enter>退出...")
				return
			}
		}
		mlog.Print(helpContent)
	}
}

// help shows more information for specified command.
func help(command string) {
	switch command {
	case "get":
		get.Help()
	case "gen":
		gen.Help()
	case "init":
		initialize.Help()
	case "docker":
		docker.Help()
	case "swagger":
		swagger.Help()
	case "build":
		build.Help()
	case "pack":
		pack.Help()
	case "run":
		run.Help()
	case "mod":
		mod.Help()
	default:
		mlog.Print(helpContent)
	}
}

// version prints the version information of the cli tool.
func version() {
	info := gbuild.Info()
	if info["git"] == "" {
		info["git"] = "none"
	}
	mlog.Printf(`GoFrame CLI工具 %s, https://goframe.org`, VERSION)
	mlog.Printf(`安装路径：%s`, gfile.SelfPath())
	if info["gf"] == "" {
		mlog.Print(`当前是自定义安装版本，没有安装信息。`)
		return
	}

	mlog.Print(gstr.Trim(fmt.Sprintf(`
Build Detail:
  Go Version:  %s
  GF Version:  %s
  Git Commit:  %s
  Build Time:  %s
`, info["go"], info["gf"], info["git"], info["time"])))
}
