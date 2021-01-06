package pack

import (
	"github.com/gogf/gf-cli/library/allyes"
	"github.com/gogf/gf-cli/library/mlog"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gres"
	"github.com/gogf/gf/text/gstr"
	"strings"
)

func Help() {
	mlog.Print(gstr.TrimLeft(`
用法 
    gf pack SRC DST

主题
    SRC  打包的源路径，可以是多个源路径。
    DST  打包文件的目标文件路径。如果文件名的扩展名是".go"，并且给出了"-n”选项，则它允许将SRC打包到文件中，否则将SRC打包到二进制文件中。

选项
    -n, --name      输出go文件的包名，如果没有传递名称，则设置为其目录名。
    -p, --prefix    打包到资源文件中的每个文件的前缀。

示例
    gf pack public data.bin
    gf pack public,template data.bin
    gf pack public,template packed/data.go
    gf pack public,template,config packed/data.go
    gf pack public,template,config packed/data.go -n=packed -p=/var/www/my-app
    gf pack /var/www/public packed/data.go -n=packed
`))
}

func Run() {
	parser, err := gcmd.Parse(g.MapStrBool{
		"n,name":   true,
		"p,prefix": true,
	})
	if err != nil {
		mlog.Fatal(err)
	}
	srcPath := parser.GetArg(2)
	dstPath := parser.GetArg(3)
	if srcPath == "" {
		mlog.Fatal("SRC路径不能为空")
	}
	if dstPath == "" {
		mlog.Fatal("DST路径不能为空")
	}
	if gfile.Exists(dstPath) && gfile.IsDir(dstPath) {
		mlog.Fatalf("DST路径'%s'不能为目录", dstPath)
	}
	if !gfile.IsEmpty(dstPath) && !allyes.Check() {
		s := gcmd.Scanf("路径'%s'不为空，文件可能被覆盖，继续？[y/n]: ", dstPath)
		if strings.EqualFold(s, "n") {
			return
		}
	}
	var (
		name   = parser.GetOpt("name")
		prefix = parser.GetOpt("prefix")
	)
	if name == "" && gfile.ExtName(dstPath) == "go" {
		name = gfile.Basename(gfile.Dir(dstPath))
	}
	if name != "" {
		if err := gres.PackToGoFile(srcPath, dstPath, name, prefix); err != nil {
			mlog.Fatalf("打包失败：%v", err)
		}
	} else {
		if err := gres.PackToFile(srcPath, dstPath, prefix); err != nil {
			mlog.Fatalf("打包失败：%v", err)
		}
	}
	mlog.Print("完成！")
}
