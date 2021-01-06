package get

import (
	"fmt"
	"github.com/gogf/gf-cli/library/mlog"
	"github.com/gogf/gf/os/gproc"
	"github.com/gogf/gf/text/gstr"
	"os"
)

func Help() {
	mlog.Print(gstr.TrimLeft(`
用法    
    gf get 包

主题 
    包 远程golang包路径，例如: github.com/gogf/gf。

示例
    gf get github.com/gogf/gf
    gf get github.com/gogf/gf@latest
    gf get github.com/gogf/gf@master
    gf get golang.org/x/sys

`))
}

func Run() {
	if len(os.Args) > 2 {
		gproc.ShellRun(fmt.Sprintf(`go get -u %s`, gstr.Join(os.Args[2:], " ")))
	} else {
		mlog.Fatal("please input the package path for get")
	}
}
