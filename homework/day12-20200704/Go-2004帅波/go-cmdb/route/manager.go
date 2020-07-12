package route
//
//import (
//	"github.com/astaxie/beego"
//	"todulist/controls"
//)
//
//func init() {
//
//	//添加任务
//	beego.Router("/task/add",&controls.Taskcontrollers{},"get:AddTaskGet;post:AddTaskPost")
//	//任务展示
//	beego.Router("/task/show",&controls.Taskcontrollers{},"get:ShowTaskGet")
//	//查询任务
//	beego.Router("/task/seltask" ,&controls.Taskcontrollers{},"get:SelTaskGet;post:SelTaskPost")
//	//暂停任务
//	beego.Router("/task/pause",&controls.Taskcontrollers{},"get:PauseTaskGet")
//	//取消任务
//	beego.Router("/task/cancel",&controls.Taskcontrollers{},"get:CancelTaskGet")
//	//重试任务
//	beego.Router("/task/retry",&controls.Taskcontrollers{},"get:RetryTaskGet")
//	//删除任务
//	beego.Router("/task/del",&controls.Taskcontrollers{},"get:DelTaskGet")
//}
