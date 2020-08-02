package controllers

import (
	"cmdb/base/controllers/base"
	"cmdb/base/errors"
	"cmdb/forms"
	"cmdb/services"
	"net/http"

	"github.com/astaxie/beego"
)

type AuthController struct {
	base.BaseController
}

func (c *AuthController) Login() {
	sessionKey := beego.AppConfig.DefaultString("auth::sessionKey", "user")
	if sessionUser := c.GetSession(sessionKey); sessionUser != nil {
		action := beego.AppConfig.DefaultString("auth::HomeAction", "HomeController.Index")
		c.Redirect(beego.URLFor(action), http.StatusFound)
		return
	}
	// 用户登陆认证
	form := &forms.LoginForm{}
	errs := errors.New()

	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			user := services.UserService.GetByName(form.Name)

			if user != nil && user.ValidPassword(form.Password) {
				sessionKey := beego.AppConfig.DefaultString("auth::sessionKey", "user")
				action := beego.AppConfig.DefaultString("auth::HomeAction", "HomeController.Index")
				c.SetSession(sessionKey, user.ID)
				c.Redirect(beego.URLFor(action), http.StatusFound)
			} else {
				beego.Error("用户名或密码错误")
				errs.Add("default", "用户名或密码错误")
			}
		}
	}
	c.Data["user"] = form
	c.Data["errors"] = errs
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "auth/login.html"
}

// Logout 用户退出登录
func (c *AuthController) Logout() {
	c.DestroySession()
	action := beego.AppConfig.DefaultString("auth::LogoutAction", "AuthController.Login")
	c.Redirect(beego.URLFor(action), http.StatusFound)
}
