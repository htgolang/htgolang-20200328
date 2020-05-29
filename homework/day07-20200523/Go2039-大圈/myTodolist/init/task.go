package init

import (
	"myTodolist/commands"
	"myTodolist/controllers"
)

func init() {
	//创建一个TaskController结构体实例
	//task := controllers.TaskController{}
	task := new(controllers.TaskController)
	//将功能注册到mgr
	commands.Register("添加任务",task.AddTask)
	commands.Register("删除任务",task.DelTask)
	commands.Register("编辑任务",task.EditTask)
	commands.Register("查看任务",task.SeeTask)
	commands.Register("修改任务状态",task.EditTaskStatus)

}
