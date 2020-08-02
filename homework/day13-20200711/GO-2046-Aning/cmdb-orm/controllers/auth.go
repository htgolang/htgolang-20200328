package controllers

import (
	"net/http"

	"cmdb/base/controllers/base"

	"github.com/astaxie/beego"

	"cmdb/base/errors"
	"cmdb/forms"
	"cmdb/models"
)

// AuthController 认证控制器
type AuthController struct {
	base.BaseController
}

// Login 认证登录
func (c *AuthController) Login() {
	sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
	sessionUser := c.GetSession(sessionKey)
	if sessionUser != nil {
		action := beego.AppConfig.DefaultString("auth::HomeAction", "HomeController.Index")
		c.Redirect(beego.URLFor(action), http.StatusFound)
		return
	}

	form := &forms.LoginForm{}
	errs := errors.New()
	// Get请求直接加载页面
	// Post请求进行数据验证
	if c.Ctx.Input.IsPost() {
		// 获取用户提交数据
		if err := c.ParseForm(form); err == nil {
			user := models.GetUserByName(form.Name)

			if user == nil {
				errs.Add("default", "用户名或密码错误")
				// 用户不存在
			} else if user.ValidPassword(form.Password) {
				// 用户密码正确
				// 记录用户状态(session 记录服务器端)
				sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
				action := beego.AppConfig.DefaultString("auth::HomeAction", "HomeController.Index")

				c.SetSession(sessionKey, user.ID)
				c.Redirect(beego.URLFor(action), http.StatusFound)
			} else {
				// 用户密码不正确
				errs.Add("default", "用户名或密码错误")
			}
		} else {
			errs.Add("default", "用户名或密码错误")
		}
	}

	c.Data["form"] = form
	c.Data["errors"] = errs
	// 定义加载页面
	c.TplName = "auth/login.html"
}

// Logout 用户退出登录
func (c *AuthController) Logout() {
	c.DestroySession()
	action := beego.AppConfig.DefaultString("auth::LogoutAction", "AuthController.Login")
	c.Redirect(beego.URLFor(action), http.StatusFound)
}
