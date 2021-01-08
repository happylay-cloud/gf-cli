package gen

import (
	"github.com/gogf/gf-cli/library/mlog"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/text/gstr"
)

func Help() {
	switch gcmd.GetArg(2) {
	case "dao":
		HelpDao()
	case "model":
		HelpModel()
	default:
		mlog.Print(gstr.TrimLeft(`
用法 
    gf gen 类型 [选项]    

类型
    dao     生成dao和model文件。
    model   生成model文件，请注意这些生成的model文件不同于"gf gen dao"命令生成的model文件。

说明
    "gen"命令具有多种用途。它目前支持为ORM模型生成go文件。
    请使用"gf gen dao -h"或"gf gen model -h"获得指定类型的帮助。

`))
	}
}

func Run() {
	parser, err := gcmd.Parse(g.MapStrBool{
		"path":            true,
		"m,mod":           true,
		"l,link":          true,
		"t,tables":        true,
		"g,group":         true,
		"c,config":        true,
		"p,prefix":        true,
		"r,remove-prefix": true,
	})
	if err != nil {
		mlog.Fatal(err)
	}
	genType := parser.GetArg(2)
	if genType == "" {
		mlog.Print("generating type cannot be empty")
		return
	}
	switch genType {
	case "model":
		doGenModel(parser)

	case "dao":
		doGenDao(parser)
	}
}
