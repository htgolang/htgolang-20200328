package router

import (
	"github.com/astaxie/beego"
	"task/controller"
)

func init() {
	beego.AutoRouter(&controller.TasksController{})
}
