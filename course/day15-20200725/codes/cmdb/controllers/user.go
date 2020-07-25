package controllers

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego"

	"cmdb/base/controllers/auth"
	"cmdb/forms"
	"cmdb/services"
)

// UserController 用户管理控制器
type UserController struct {
	auth.LayoutController
}

// Query 查询用户
func (c *UserController) Query() {
	flash := beego.ReadFromRequest(&c.Controller)
	fmt.Println(flash.Data)
	q := c.GetString("q")

	c.Data["users"] = services.UserService.Query(q)
	c.Data["q"] = q

	c.TplName = "user/query.html"
}

// Modify 修改用户
func (c *UserController) Modify() {
	// 假设当前用户不能修改其他人的信息
	// if 1 == 1 {

	// 	c.Abort("404")
	// 	return
	// }

	form := &forms.UserModifyForm{}
	// GET 获取数据
	// POST 修改用户
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			// 验证数据
			services.UserService.Modify(form)

			// 存储消息
			flash := beego.NewFlash()
			flash.Set("notice", "修改用户信息成功")
			// flash.Notice("修改用户信息成功")
			flash.Error("error")
			// flash.Set("error", "error")
			flash.Success("success")
			// flash.Set("success", "error")
			flash.Warning("warning")
			flash.Store(&c.Controller)

			c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
		}
	} else if pk, err := c.GetInt("pk"); err == nil {
		if user := services.UserService.GetByPk(pk); user != nil {
			form.ID = user.ID
			form.Name = user.Name
		}
	}

	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "user/modify.html"
}

// Delete 删除用户数据
func (c *UserController) Delete() {
	if pk, err := c.GetInt("pk"); err == nil && c.LoginUser.ID != pk {
		services.UserService.Delete(pk)
	}
	c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
}
