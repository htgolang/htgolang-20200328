package controllers

import (
	"net/http"

	"github.com/astaxie/beego"

	"cmdb/base/controllers/auth"
	"cmdb/forms"
	"cmdb/models"
)

// UserController 用户管理控制器
type UserController struct {
	auth.AuthorizationController
}

// Query 查询用户
func (c *UserController) Query() {
	q := c.GetString("q")

	users := models.QueryUser(q)
	c.Data["users"] = users
	c.Data["q"] = q
	c.TplName = "user/query.html"
}

// Modify 修改用户
func (c *UserController) Modify() {
	form := &forms.UserModifyForm{}
	// GET 获取数据
	// POST 修改用户
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			// 验证数据
			models.ModifyUser(form)
			c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
		}
	} else if pk, err := c.GetInt("pk"); err == nil {
		if user := models.GetUserByPk(pk); user != nil {
			form.ID = user.ID
			form.Name = user.Name
		}
	}

	c.Data["form"] = form
	c.TplName = "user/modify.html"
}

// Delete 删除用户数据
func (c *UserController) Delete() {
	if pk, err := c.GetInt("pk"); err == nil && c.LoginUser.ID != pk {
		models.DeleteUser(pk)
	}
	c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
}
