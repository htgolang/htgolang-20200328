package controllers

import (
	"cmdb/base/controllers/base"
	"cmdb/forms"
	"cmdb/models"
	"cmdb/utils"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

// AuthController 登录验证控制器
type AuthController struct {
	base.Base
}

// Login 用户登录
func (c AuthController) Login() {
	var (
		authForm forms.AuthForm
		errors   map[string]string
	)
	if sessionUser := c.GetSession(beego.AppConfig.DefaultString("session::Name", "UserID")); sessionUser != nil {
		c.Redirect(beego.URLFor(beego.AppConfig.DefaultString("auth::HomeAction", "HomeController.Index")), http.StatusFound)
		return
	}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(&authForm); err == nil {
			if user := models.GetUserByNickname(authForm.Name); user == nil {
				errors["submit"] = "username not exist"
			} else if utils.ValidPassword(authForm.Password, user.Password) {
				c.SetSession(beego.AppConfig.DefaultString("session::Name", "UserID"), user.ID)
				c.Redirect(beego.URLFor(beego.AppConfig.DefaultString("auth::HomeAction", "HomeController.Index")), http.StatusFound)
			} else {
				errors["submit"] = "password error"
			}
		} else {
			errors["submit"] = "username or password error"
		}
	}
	c.Data["Form"] = authForm
	c.Data["Errors"] = errors
	c.TplName = "auth/login.html"
}

// Logout 用户退出
func (c AuthController) Logout() {
	c.DestroySession()
	c.Redirect(beego.URLFor(
		beego.AppConfig.DefaultString(
			"auth::LogoutAction",
			"AuthController.Login",
		)),
		http.StatusFound,
	)
}

// Register 用户注册
func (c AuthController) Register() {
	var (
		userForm forms.AddUserForm
		errors   map[string]string
	)
	if c.Ctx.Input.IsPost() {
		err := c.ParseForm(userForm)
		if err != nil {
			logs.Warn("parse form AddUserForm error, %s", err)
			errors["submit"] = "server error"
		} else {
			valid := validation.Validation{}
			if ok, err := valid.Valid(&userForm); err != nil {
				logs.Warn("valid handle error, %s", err)
				errors["submit"] = "server error"
			} else if !ok {
				for _, err := range valid.Errors {
					errors[err.Key] = err.Message
				}
			} else {
				user := models.User{
					StaffID:      userForm.StaffID,
					Name:         userForm.Name,
					NickName:     userForm.NickName,
					Gender:       userForm.Gender,
					Tel:          userForm.Tel,
					Email:        userForm.Email,
					Addr:         userForm.Addr,
					DepartmentID: userForm.DepartmentID,
					Title:        userForm.Title,
					StatusID:     0,
					Password:     utils.HashPasswd(userForm.Password),
				}
				if err := models.AddUser(user); err == nil {
					c.Redirect(c.URLFor("AuthController.Login"), http.StatusFound)
				} else {
					logs.Error("update data error, %s", err)
					errors["submit"] = "update data error"
				}
			}
		}
	}
	c.Data["Errors"] = errors
	c.Data["Department"] = userForm
	c.TplName = "user/register.html"
}
