package controllers

import (
	"cmdb/modles"
	"cmdb/utils"
	"net/http"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

//执行完不往下执行
func (c *UserController) Prepare() {
	sessionUser := c.GetSession("user")
	if sessionUser == nil {
		c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
		c.StopRun()
	}
}

func (c *UserController) Query() {
	//获取session判断是否为空

	//此处判断有prepare判断
	// sessionUser := c.GetSession("user")
	// //sesson为空返回登陆页面
	// if sessionUser == nil {
	// 	//无sesion
	// 	//断言  =》 int
	// 	c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
	// 	return
	// }

	q := utils.Like(c.GetString("q"))
	// q := c.GetString("q")
	users := modles.QueryUser(q)
	c.Data["users"] = users
	c.TplName = "user/query.html"
}
