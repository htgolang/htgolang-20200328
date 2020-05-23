package init

import (
	"todolist/commands"
	"todolist/controllers"
)

func init() {
	task := &controllers.TaskController{}

	commands.Register("添加任务", task.Add)
	commands.Register("删除任务", task.Delete)
	commands.Register("修改任务", task.Modify)
	commands.Register("修改任务状态", task.ModifyStatus)
	commands.Register("查询任务", task.Query)
}
