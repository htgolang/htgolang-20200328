package routers

// 先到内置
// 第三方
// 当前项目

import (
	"github.com/astaxie/beego"

	"cmdb/controllers"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.HomeController{}, "*:Index")

	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.HomeController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.PasswordController{})
}
