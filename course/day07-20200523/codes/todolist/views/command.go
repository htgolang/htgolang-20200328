package views

import (
	"os"
	"strconv"
	"todolist/commands/command"

	"github.com/olekukonko/tablewriter"
)

type commandView struct {
}

func (v *commandView) Menu(cmds map[int]*command.Command) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"选项", "说明"})
	for i := 1; i <= len(cmds); i++ {
		table.Append([]string{
			strconv.Itoa(i),
			cmds[i].Name,
		})
	}
	table.Render()
}

var CommandView = new(commandView)
