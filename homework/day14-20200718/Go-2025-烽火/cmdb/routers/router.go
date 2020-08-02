package routers

import (
	"cmdb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "*:Index")
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.TaskController{})
}
