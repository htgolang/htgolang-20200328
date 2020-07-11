package controllers

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego"

	"cmdb/models"
)

// UserController 用户管理控制器
type UserController struct {
	beego.Controller
}

func (c *UserController) Prepare() {
	sessionUser := c.GetSession("user")
	if sessionUser == nil {
		// 无session信息(未登录)
		c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
		c.StopRun()
	}
}

// Query 查询用户
func (c *UserController) Query() {
	fmt.Println("Query")
	q := c.GetString("q")

	users := models.QueryUser(q)
	c.Data["users"] = users
	c.Data["q"] = q
	c.TplName = "user/query.html"
}
