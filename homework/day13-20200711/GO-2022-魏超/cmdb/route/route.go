package route

import (
	"github.com/astaxie/beego"

	"cmdb/controllers"
)

func init() {
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.HomeController{})
	beego.AutoRouter(&controllers.UserController{})
}
