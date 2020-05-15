package route

import (
	"github.com/astaxie/beego"
	"todulist/controls"
)

func init() {
	beego.Router("/LoginErr",&controls.ErrControls{},"get:LoginErr")
	beego.Router("/RegErr",&controls.ErrControls{},"get:RegErr")
	beego.Router("/user/err",&controls.ErrControls{},"get:UserErr")
	beego.Router("/task/err",&controls.ErrControls{},"get:TaskErr")
	beego.Router("/role/err",&controls.ErrControls{},"get:RoleErr")
}