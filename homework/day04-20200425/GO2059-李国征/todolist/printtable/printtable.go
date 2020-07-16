package printtable

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func PrintTask(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"序号", "名称", "开始时间", "结束时间", "用户", "状态"})
	table.SetBorder(false)
	table.AppendBulk(data)
	table.Render()
}
