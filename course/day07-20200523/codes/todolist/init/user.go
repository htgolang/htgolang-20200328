package init

import (
	"todolist/commands"
	"todolist/controllers"
)

func init() {
	commands.RegisterLoginCallback(controllers.Login)
	commands.Register("退出", controllers.Logout)
}
