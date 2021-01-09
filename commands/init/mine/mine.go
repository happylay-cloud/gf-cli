package mine

import (
	"github.com/gogf/gf-cli/library/allyes"
	"github.com/gogf/gf-cli/library/mlog"
	"github.com/gogf/gf/encoding/gcompress"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
	"strings"
)

const (
	emptyProject     = "github.com/happylay-cloud/gf-empty"
	emptyProjectName = "gf-empty"
)

var (
	cdnUrl  = g.Config("url").GetString("cdn.url")
	homeUrl = g.Config("url").GetString("home.url")
)

func init() {
	if cdnUrl == "" {
		mlog.Fatal("CDN配置不能为空")
	}
	if homeUrl == "" {
		mlog.Fatal("Home配置不能为空")
	}
}

func Help() {
	mlog.Print(gstr.TrimLeft(`
用法    
    gfctl init@mine 名称

主题 
    名称 项目名称。它将在当前目录中创建一个同名的文件夹。该名称也将是项目的模块名称。

示例
    gfctl init@mine my-app
    gfctl init@mine my-project-name
`))
}

func Run() {
	parser, err := gcmd.Parse(nil)
	if err != nil {
		mlog.Fatal(err)
	}
	projectName := parser.GetArg(2)
	if projectName == "" {
		mlog.Fatal("项目名称不应为空")
	}
	dirPath := projectName
	if !gfile.IsEmpty(dirPath) && !allyes.Check() {
		s := gcmd.Scanf(`文件夹"%s"不是空的，文件可能被覆盖，继续吗？ [y/n]: `, projectName)
		if strings.EqualFold(s, "n") {
			return
		}
	}
	mlog.Print("正在初始化...")

	// Zip data retrieving.
	respData, err := ghttp.Get("https://github.com/happylay-cloud/gf-empty/archive/mine.zip")
	if err != nil {
		mlog.Fatal("获取项目zip数据失败：%s", err.Error())
	}
	defer respData.Close()
	zipData := respData.ReadAll()
	if len(zipData) == 0 {
		mlog.Fatal("获取项目数据失败：数据值为空。可能是网络问题，再试一次？")
	}

	// Unzip the zip data.
	if err = gcompress.UnZipContent(zipData, dirPath, emptyProjectName+"-mine"); err != nil {
		mlog.Fatal("解压缩项目数据失败，", err.Error())
	}
	// Replace project name.
	if err = gfile.ReplaceDir(emptyProject, projectName, dirPath, "Dockerfile,*.go,*.MD,*.mod", true); err != nil {
		mlog.Fatal("内容替换失败，", err.Error())
	}
	if err = gfile.ReplaceDir(emptyProjectName, projectName, dirPath, "Dockerfile,*.go,*.MD,*.mod", true); err != nil {
		mlog.Fatal("内容替换失败，", err.Error())
	}
	mlog.Print("初始化完成！")
	mlog.Print("你现在可以运行以下命令开始你的旅程吧，享受吧！")
	mlog.Printf("cd %s", projectName)
	mlog.Printf("git init")
	mlog.Printf("%s", "gfctl run main.go")

}
