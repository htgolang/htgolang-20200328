package controllers

import (
	"net/http"

	"cmdb/base/controllers/auth"
	"cmdb/forms"
	"cmdb/services"

	"github.com/astaxie/beego"
)

type UserController struct {
	auth.AuthController
}

// func (c *UserController) Prepare() {
// 	if c.LoginUser.IsAdmin == 0 {
// 		beego.Warning(fmt.Sprintf("current user:{} has no privilege to access current page", c.LoginUser.Name))
// 		c.Redirect(beego.URLFor("HomeController.Index"), http.StatusFound)
// 	}
// }

//添加用户
func (c *UserController) Add() {
	form := &forms.FormUser{}
	if c.Ctx.Input.IsPost() {
		c.ParseForm(form)
		if err := services.UserService.Add(form); err != nil {
			beego.Error(err.Error())
		} else {
			beego.Informational("添加新用户成功")
			action := beego.AppConfig.DefaultString("auth::UserAction", "UserController.Query")
			c.Redirect(beego.URLFor(action), http.StatusFound)
		}
	}
	c.Data["user"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "user/add.html"
}

// 查询用户
func (c *UserController) Query() {
	q := c.GetString("query")
	users := services.UserService.Query(q)
	c.Data["users"] = users
	c.Data["query"] = q
	c.TplName = "user/query.html"
}

// 修改用户
func (c *UserController) Modify() {
	formUser := &forms.FormUser{}

	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(formUser); err == nil {
			services.UserService.Modify(formUser)
			action := beego.AppConfig.DefaultString("auth::UserAction", "UserController.Query")
			c.Redirect(beego.URLFor(action), http.StatusFound)
		}
	} else if id, err := c.GetInt("id"); err == nil {
		if user := services.UserService.GetByPk(id); user != nil {
			formUser.ID = user.ID
			formUser.Name = user.Name
			formUser.Password = user.Password
			formUser.Nickname = user.Nickname
			formUser.Gender = user.Gender
			formUser.Tel = user.Tel
			formUser.Addr = user.Addr
			formUser.Email = user.Email
			formUser.Status = user.Status
			formUser.Department = user.Department
		}
	}

	c.Data["user"] = formUser
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "user/modify.html"
}

// 删除用户
func (c *UserController) Delete() {
	if uid, err := c.GetInt("id"); err == nil {
		if uid != c.LoginUser.ID {
			services.UserService.Delete(uid)
			c.Redirect(beego.URLFor(`UserController.Query`), http.StatusFound)
		}
	}
}
