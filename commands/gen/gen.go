package gen

import (
	"github.com/gogf/gf-cli/library/mlog"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/text/gstr"
)

func Help() {
	mlog.Print(gstr.TrimLeft(`
用法 
    gf gen 类型 [选项]

类型
    dao     生成dao和model文件。
    model   生成model文件，请注意这些生成的model文件不同于"gf gen dao"命令生成的model文件。

选项
    -/--path             生成文件的目录路径。
    -l, --link           数据库配置，请参考：https://goframe.org/database/gdb/config
    -t, --tables         仅为给定的表生成模型，多个表名用','分隔。
    -g, --group          指定数据库的配置组名，这是不必要的，默认值是"default"。
    -c, --config         用于指定数据库的配置文件，通常不需要。
                         如果没有通过"-l"，它将搜索"./config.toml"和"./config/config.toml"，默认在当前工作目录。
    -p, --prefix         为指定链接/数据库表的所有表添加前缀。
    -r, --remove-prefix  删除表的指定前缀，多个前缀以','分隔。 
    -m, --mod            生成的golang文件导入的模块名。
                  
配置支持
    配置文件也支持选项。配置节点名为"gf.gen"，它也支持多个数据库，例如:
    [gfcli]
        [[gfcli.gen.dao]]
            link   = "mysql:root:12345678@tcp(127.0.0.1:3306)/test"
            tables = "order,products"
        [[gfcli.gen.dao]]
            link   = "mysql:root:12345678@tcp(127.0.0.1:3306)/primary"
            path   = "./my-app"
            prefix = "primary_"
            tables = "user, userDetail"

示例
    gf gen dao
        gf gen dao
        gf gen dao -l "mysql:root:12345678@tcp(127.0.0.1:3306)/test"
        gf gen dao -path ./model -c config.yaml -g user-center -t user,user_detail,user_login
        gf gen dao -r user_
    gf gen model
        gf gen model
        gf gen model -l "mysql:root:12345678@tcp(127.0.0.1:3306)/test"
        gf gen model -path ./model -c config.yaml -g user-center -t user,user_detail,user_login
        gf gen model -r user_

说明
    "gen"命令具有多种用途。它目前支持为ORM模型生成go文件。
`))
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
