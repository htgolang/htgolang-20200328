package route

import (
	"github.com/astaxie/beego"
	"github.com/strive-after/go-kubernetes/controls"
)

func init() {

	//添加用户
	//beego.Router("/user/add",&controls.UserControoler{},"get:AddUserGet;post:AddUserPost")
	//显示所有用户
	beego.Router("/user/show",&controls.UserControoler{},"get:ShowUserGet")
	//查看用户信息
	beego.Router("/user/info",&controls.UserControoler{},"get:UserInfo")
	//修改用户
	beego.Router("/user/change",&controls.UserControoler{},"get:ChangeUserGet;post:ChangeUserPost")
	//删除用户
	beego.Router("/user/del",&controls.UserControoler{},"get:Del")
	//查看自己的信息
	beego.Router("/user/my/info",&controls.UserControoler{},"get:MyInfoGet;post:MyInfoPost")
	//修改当前用户密码
	beego.Router("/user/change/mypass",&controls.UserControoler{},"get:MyPassGet;post:MyPassPost")
	//管理员对用户做密码重置
	beego.Router("/user/change/userpass",&controls.UserControoler{},"get:UserPassGet;post:UserPassPost")
}