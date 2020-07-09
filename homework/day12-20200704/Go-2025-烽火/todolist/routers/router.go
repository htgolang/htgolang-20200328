package routers

import (
	"time"
	"todolist/base"
	"todolist/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.AddFuncMap("datetime", func(tt *time.Time) string {
		if tt == nil {
			return "--"
		}
		return tt.Format(base.TimeLayout)
	})

	beego.Router("/", &controllers.TaskController{}, "get:Index")
	beego.Router("/task/add/", &controllers.TaskController{}, "get,post:Add")
	beego.Router("/task/modify/", &controllers.TaskController{}, "get,post:Modify")
	beego.Router("/task/query/", &controllers.TaskController{}, "get,post:Query")
	beego.Router("/task/delete/", &controllers.TaskController{}, "get:Del")
	// beego.AutoRouter(&controllers.TaskController{})
}
