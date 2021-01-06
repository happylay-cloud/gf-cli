package env

import (
	"bytes"
	"github.com/gogf/gf-cli/library/mlog"
	"github.com/gogf/gf/os/gproc"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/text/gstr"
	"github.com/olekukonko/tablewriter"
)

func Run() {
	result, err := gproc.ShellExec("go env")
	if err != nil {
		mlog.Fatal(err)
	}
	if result == "" {
		mlog.Fatal(`检索Golang环境变量失败，您是否安装了Golang？`)
	}
	var (
		lines  = gstr.Split(result, "\n")
		buffer = bytes.NewBuffer(nil)
	)
	array := make([][]string, 0)
	for _, line := range lines {
		line = gstr.Trim(line)
		if line == "" {
			continue
		}
		if gstr.Pos(line, "set ") == 0 {
			line = line[4:]
		}
		match, _ := gregex.MatchString(`(.+?)=(.*)`, line)
		if len(match) < 3 {
			mlog.Fatalf(`无效的Golang环境变量："%s"`, line)
		}
		array = append(array, []string{gstr.Trim(match[1]), gstr.Trim(match[2])})
	}
	tw := tablewriter.NewWriter(buffer)
	tw.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT})
	tw.AppendBulk(array)
	tw.Render()
	mlog.Print(buffer.String())
}
