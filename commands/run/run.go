package run

import (
	"fmt"
	"github.com/gogf/gf-cli/commands/swagger"
	"github.com/gogf/gf-cli/library/mlog"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/container/gtype"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gfsnotify"
	"github.com/gogf/gf/os/gproc"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/text/gstr"
	"os"
	"runtime"
	"strings"
	"time"
)

type App struct {
	File    string // Go run file name/path.
	Options string // Extra "go run" options.
	Args    string // Auto parse and pack swagger files.
	Swagger bool   // Auto parse and pack swagger files.
}

const (
	gPROXY_CHECK_TIMEOUT = time.Second
)

var (
	process    *gproc.Process
	httpClient = ghttp.NewClient()
)

func init() {
	httpClient.SetTimeout(gPROXY_CHECK_TIMEOUT)
}

func Help() {
	mlog.Print(gstr.TrimLeft(`
用法
    gf run 文件 [选项]

主题
    文件 构建文件路径。
    选项 与"go run"/"go build"有相同的选项，除了以下定义的一些选项。

选项
    -/--args     自定义流程参数。
    -/--swagger  在运行之前，自动解析并打包swagger到packed/data-swagger.go。

示例
    gf run main.go
    gf run main.go --swagger
    gf run main.go --args "server -p 8080"
    gf run main.go -mod=vendor

说明
    "run"命令用于运行具有类似热编译功能的go代码，当代码更改时，它将异步编译并运行代码。
`))
}

func Run() {
	parser, err := gcmd.Parse(g.MapStrBool{
		"args": true,
	})
	if err != nil {
		mlog.Fatal(err)
	}
	mlog.SetHeaderPrint(true)
	file := gcmd.GetArg(2)
	if len(file) < 1 {
		mlog.Fatal("文件路径不能为空")
	}
	app := &App{
		File: file,
	}
	// ================================================================================
	// This command is very special that it supports options of "go run" and "go build"
	// from the third parameter of os.Args. That means, we should filter any parameter
	// that "go run" and "go build" do not allow.
	// ================================================================================
	// Swagger checks.
	array := garray.NewStrArrayFrom(os.Args)
	index := array.Search("--swagger")
	if index < 0 {
		index = array.Search("-swagger")
	}
	if index != -1 {
		app.Swagger = true
		array.Remove(index)
	}
	// args checks.
	args := parser.GetOpt("args")
	if args != "" {
		app.Args = args
		index := -1
		array.Iterator(func(k int, v string) bool {
			if gstr.Contains(v, "-args") {
				index = k
				return false
			}
			return true
		})
		if index != -1 {
			v, _ := array.Get(index)
			if gstr.Contains(v, "=") {
				array.Remove(index)
			} else {
				array.Remove(index)
				array.Remove(index)
			}
		}
	}
	// -y checks
	array.RemoveValue("-y")
	array.RemoveValue("--y")
	if array.Len() > 3 {
		app.Options = strings.Join(array.SubSlice(3), " ")
	}
	dirty := gtype.NewBool()
	_, err = gfsnotify.Add(gfile.RealPath("."), func(event *gfsnotify.Event) {
		if gfile.ExtName(event.Path) != "go" {
			return
		}
		// Ignore swagger file.
		if gfile.Basename(event.Path) == "data-swagger.go" {
			return
		}
		// Variable <dirty> is used for running the changes only one in one second.
		if !dirty.Cas(false, true) {
			return
		}
		// With some delay in case of multiple code changes in very short interval.
		gtimer.SetTimeout(1500*gtime.MS, func() {
			defer dirty.Set(false)
			mlog.Printf(`go文件变动：%s`, event.String())
			app.Run()
		})
	})
	if err != nil {
		mlog.Fatal(err)
	}
	go app.Run()
	select {}
}

func (app *App) Run() {
	// Rebuild and run the codes.
	renamePath := ""
	mlog.Printf("构建：%s", app.File)
	outputPath := gfile.Join("bin", gfile.Name(app.File))
	if runtime.GOOS == "windows" {
		outputPath += ".exe"
		if gfile.Exists(outputPath) {
			renamePath = outputPath + "~"
			if err := gfile.Rename(outputPath, renamePath); err != nil {
				mlog.Print(err)
			}
		}
	}
	// Auto swagger.
	if app.Swagger {
		if err := gproc.ShellRun(`gfctl swagger`); err != nil {
			return
		}
		if gfile.Exists("swagger") {
			packCmd := fmt.Sprintf(`gfctl pack %s packed/%s -n packed -y`, "swagger", swagger.PackedGoFileName)
			mlog.Print(packCmd)
			if err := gproc.ShellRun(packCmd); err != nil {
				return
			}
		}
	}
	// In case of `pipe: too many open files` error.
	// Build the app.
	buildCommand := fmt.Sprintf(`go build -o %s %s %s`, outputPath, app.Options, app.File)
	mlog.Print(buildCommand)
	result, err := gproc.ShellExec(buildCommand)
	if err != nil {
		mlog.Printf("构建错误：\n%s%s", result, err.Error())
		return
	}
	// Kill the old process if build successfully.
	if process != nil {
		if err := process.Kill(); err != nil {
			mlog.Debugf("终止进程错误：%s", err.Error())
			//return
		}
	}
	// Run the binary file.
	runCommand := fmt.Sprintf(`%s %s`, outputPath, app.Args)
	mlog.Print(runCommand)
	if runtime.GOOS == "windows" {
		// Special handling for windows platform.
		// DO NOT USE "cmd /c" command.
		process = gproc.NewProcess(runCommand, nil)
	} else {
		process = gproc.NewProcessCmd(runCommand, nil)
	}
	if pid, err := process.Start(); err != nil {
		mlog.Printf("即时编译运行错误：%s", err.Error())
	} else {
		mlog.Printf("即时编译运行进程编号pid：%d", pid)
	}
}
