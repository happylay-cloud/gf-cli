package swagger

import (
	"errors"
	"fmt"
	"github.com/gogf/gf-cli/library/mlog"
	"github.com/gogf/gf/container/gtype"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gfsnotify"
	"github.com/gogf/gf/os/gproc"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/swagger"
)

const (
	defaultOutput    = "./swagger"
	swaggoRepoPath   = "github.com/swaggo/swag/cmd/swag"
	PackedGoFileName = "swagger.go"
)

func Help() {
	mlog.Print(gstr.TrimLeft(`
用法    
    gf swagger [选项]

选项
    -s, --server  在生成swagger文件之后，在指定的地址启动swagger服务器。
    -o, --output  存储已解析的swagger文件的输出目录，默认输出目录是"./swagger"。
    -/--pack      自动将swagger解析并将其打包到packed/swagger.go。 

示例
    gf swagger
    gf swagger --pack
    gf swagger -s 8080
    gf swagger -s 127.0.0.1:8080
    gf swagger -o ./document/swagger

说明
    "swagger"命令用于解析当前项目，并生成swagger api描述文件，可在swagger api服务器中使用。 
    如果与"-s/--server"选项一起使用，则监视当前项目的go文件的更改，并复制swagger文件，这对于本地API开发非常方便。
    如果命令"swag"失败，请首先检查系统路径是否包含go二进制路径，也可以参考以下内容手动安装"swag"工具：
    https://github.com/swaggo/swag
`))
}

func Run() {
	mlog.SetHeaderPrint(true)
	parser, err := gcmd.Parse(g.MapStrBool{
		"s,server": true,
		"o,output": true,
		"pack":     false,
	})
	if err != nil {
		mlog.Fatal(err)
	}
	server := parser.GetOpt("server")
	output := parser.GetOpt("output", defaultOutput)
	// Generate swagger files.
	if err := generateSwaggerFiles(output, parser.ContainsOpt("pack")); err != nil {
		mlog.Print(err)
	}
	// Watch the go file changes and regenerate the swagger files.
	dirty := gtype.NewBool()
	_, err = gfsnotify.Add(gfile.RealPath("."), func(event *gfsnotify.Event) {
		if gfile.ExtName(event.Path) != "go" || gstr.Contains(event.Path, "swagger") {
			return
		}
		// Variable <dirty> is used for running the changes only one in one second.
		if !dirty.Cas(false, true) {
			return
		}
		// With some delay in case of multiple code changes in very short interval.
		gtimer.SetTimeout(1500*gtime.MS, func() {
			mlog.Printf(`go文件改变：%s`, event.String())
			mlog.Print(`复制swagger文件...`)
			if err := generateSwaggerFiles(output, parser.ContainsOpt("pack")); err != nil {
				mlog.Print(err)
			} else {
				mlog.Print(`完成！`)
			}
			dirty.Set(false)
		})
	})
	if err != nil {
		mlog.Fatal(err)
	}
	// Swagger server starts.
	if server != "" {
		if gstr.IsNumeric(server) {
			server = ":" + server
		}
		s := g.Server()
		s.Plugin(&swagger.Swagger{})
		s.SetAddr(server)
		s.Run()
	}
}

// generateSwaggerFiles generates necessary swagger files.
func generateSwaggerFiles(output string, pack bool) error {
	mlog.Print(`生成swagger文件...`)
	// Temporary storing swagger files directory.
	tempOutputPath := gfile.Join(gfile.TempDir(), "swagger")
	if gfile.Exists(tempOutputPath) {
		gfile.Remove(tempOutputPath)
	}
	gfile.Mkdir(tempOutputPath)
	// Check and install swag tool.
	swag := gproc.SearchBinary("swag")
	if swag == "" {
		err := gproc.ShellRun(fmt.Sprintf(`go get -u -v %s`, swaggoRepoPath))
		if err != nil {
			return err
		}
	}
	// Generate swagger files using swag.
	command := fmt.Sprintf(`swag init -o %s`, tempOutputPath)
	result, err := gproc.ShellExec(command)
	if err != nil {
		return errors.New(result + err.Error())
	}
	if !gfile.Exists(gfile.Join(tempOutputPath, "swagger.json")) {
		return errors.New("make swagger files failed")
	}
	if !gfile.Exists(output) {
		gfile.Mkdir(output)
	}
	if err = gfile.CopyFile(
		gfile.Join(tempOutputPath, "swagger.json"),
		gfile.Join(output, "swagger.json"),
	); err != nil {
		return err
	}
	mlog.Print(`完成！`)
	// Auto pack into go file.
	if pack && gfile.Exists("swagger") {
		packCmd := fmt.Sprintf(`gf pack %s packed/%s -n packed`, "swagger", PackedGoFileName)
		mlog.Print(packCmd)
		if err := gproc.ShellRun(packCmd); err != nil {
			return err
		}
	}
	return nil
}
