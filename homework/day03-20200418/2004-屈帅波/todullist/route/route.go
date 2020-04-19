package route

import (
	"github.com/astaxie/beego"
	"todulist/controls"
)

func init()  {
	beego.Router("/",&controls.LoginControllers{},"get:Get")
	//登陆页面
	beego.Router("/login",&controls.LoginControllers{},"get:LoginGet;post:LoginPost")
	//注册页面
	beego.Router("/register",&controls.RegisterControllers{},"get:RegGet;post:RegPost")
	//总操作列表
	beego.Router("/operation",&controls.OperationCotrollers{},"get:AllOperGet")

	//用户管理操作
	beego.Router("/user",&controls.OperationCotrollers{},"get:UserOperGet")
	//添加用户
	beego.Router("/user/add",&controls.UserControoler{},"get:AddUserGet;post:AddUserPost")
	//显示所有用户
	beego.Router("/user/show",&controls.UserControoler{},"get:ShowUserGet")
	//查询条件
	beego.Router("/user/seluser",&controls.UserControoler{},"get:SelUserGet;post:SelUserPost")
	//修改用户
	beego.Router("/user/change",&controls.UserControoler{},"get:ChangeUserGet;post:ChangeUserPost")
	//删除用户
	beego.Router("/user/del",&controls.UserControoler{},"get:DelUserGet")

	//任务管理操作
	beego.Router("/task",&controls.OperationCotrollers{},"get:TaskOperGet")
	//添加任务
	beego.Router("/task/add",&controls.Taskcontrollers{},"get:AddTaskGet;post:AddTaskPost")
	//添加任务
	beego.Router("/task/show",&controls.Taskcontrollers{},"get:ShowTaskGet")
	//查询任务
	beego.Router("/task/seltask" ,&controls.Taskcontrollers{},"get:SelTaskGet;post:SelTaskPost")
	//删除任务
	beego.Router("/task/del",&controls.Taskcontrollers{},"get:DelTaskGet")


}


