package controllers

import (
	"net/http"

	"cmdb/base/controllers/auth"
	"cmdb/forms"
	"cmdb/models"
	"cmdb/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	auth.AuthController
}

//添加用户
func (c *UserController) Add() {
	user := models.User{}
	if c.Ctx.Input.IsPost() {
		c.ParseForm(&user)
		tempPw := user.Password
		user.Password = utils.GeneratePassword(tempPw)
		user.StaffID = models.GenerateStaffID()
		ormer := orm.NewOrm()
		ormer.Insert(&user)
		action := beego.AppConfig.DefaultString("auth::UserAction", "UserController.Query")
		c.Redirect(beego.URLFor(action), http.StatusFound)
	}
	c.Data["user"] = user
	c.TplName = "user/adduser.html"
}

// 查询用户
func (c *UserController) Query() {
	q := c.GetString("query")
	users := models.QueryUser(q)
	c.Data["users"] = users
	c.Data["query"] = q
	c.TplName = "user/queryuser.html"
}

// 修改用户
func (c *UserController) Modify() {
	formUser := &forms.FormUser{}

	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(formUser); err == nil {
			models.ModifyUser(formUser)
			action := beego.AppConfig.DefaultString("auth::UserAction", "UserController.Query")
			c.Redirect(beego.URLFor(action), http.StatusFound)
		}
	} else if id, err := c.GetInt("id"); err == nil {
		if user := models.GetUserByPk(id); user != nil {
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
	c.TplName = "user/modifyuser.html"
}

// 删除用户
func (c *UserController) Delete() {
	if uid, err := c.GetInt("id"); err == nil {
		if uid != c.LoginUser.ID {
			models.DeleteUser(uid)
			c.Redirect(beego.URLFor(`UserController.Query`), http.StatusFound)
		}
	}
}
