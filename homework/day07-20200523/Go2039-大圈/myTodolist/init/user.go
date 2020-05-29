package init

import (
	"myTodolist/commands"
	"myTodolist/controllers"
)

//对user进行初始化
//初始化函数会在导入包(init包)的时候执行
func init()  {
	//传入 Name 和 CallBack 然后注册到了mgr中
	commands.Register("退出",controllers.Logout)
	//传入一个Login函数，注册到mgr中
	commands.RegisterLoginCallBack(controllers.Login)
}
