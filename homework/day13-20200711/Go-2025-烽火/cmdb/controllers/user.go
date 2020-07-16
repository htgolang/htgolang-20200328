package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/models"
	"cmdb/utils"
	"net/http"

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
		action := beego.AppConfig.DefaultString("auth::HomeAction", "HomeController.Index")
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
	user := models.GetUserById(c.GetString("id"))
	if c.Ctx.Input.IsPost() {
		c.ParseForm(user)
		models.ModifyUser(user)
		action := beego.AppConfig.DefaultString("auth::HomeIndex", "HomeController.Index")
		c.Redirect(beego.URLFor(action), http.StatusFound)
	}
	c.Data["user"] = user
	c.TplName = "user/modifyuser.html"
}

// 删除用户
func (c *UserController) Delete() {
	models.DeleteUser(c.GetString("id"))
	action := beego.AppConfig.DefaultString("auth::HomeIndex", "HomeController.Index")
	c.Redirect(beego.URLFor(action), http.StatusFound)
}
