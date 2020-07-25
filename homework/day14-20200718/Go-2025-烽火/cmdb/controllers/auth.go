package controllers

import (
	"cmdb/base/controllers/base"
	"cmdb/base/errors"
	"cmdb/forms"
	"cmdb/models"
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
			user := models.GetUserByName(form.Name)
			// fmt.Printf("%#v\n", user)
			if user != nil && user.ValidatePassword(form.Password) {
				sessionKey := beego.AppConfig.DefaultString("auth::sessionKey", "user")
				action := beego.AppConfig.DefaultString("auth::HomeAction", "HomeController.Index")
				c.SetSession(sessionKey, user.ID)
				// fmt.Println("session:", sessionKey, c.GetSession(sessionKey))
				// fmt.Println("-----set session------")
				c.Redirect(beego.URLFor(action), http.StatusFound)
			} else {
				errs.Add("default", "用户名或密码错误")
			}
		}
	}
	c.Data["user"] = form
	c.Data["errors"] = errs
	c.TplName = "auth/login.html"
}
