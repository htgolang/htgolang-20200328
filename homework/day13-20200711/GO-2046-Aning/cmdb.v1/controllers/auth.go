package controllers

import (
	"net/http"

	"github.com/astaxie/beego"

	"cmdb/base/errors"
	"cmdb/forms"
	"cmdb/modles"
)

//认证控制器
type AuthController struct {
	beego.Controller
}

func (c *AuthController) Login() {
	//登陆前判断,若已存在直接登录
	sessionUser := c.GetSession("user")
	if sessionUser != nil {
		c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
	}
	//get 直接加载
	//post   验证
	// 验证成功
	// 验证失败
	form := &forms.LoginForm{}
	errs := errors.New()
	if c.Ctx.Input.IsPost() {
		//获取提交的数据
		// c.GetString("name")
		// c.GetString("password")
		if err := c.ParseForm(form); err == nil {
			user := modles.GetUserByName(form.Name)
			// fmt.Println(user)
			if user == nil {
				// errs.Add("default", "name  or passwd error")
				errs.Add("default", "用户名或密码错误")
			} else if user.ValidPassword(form.Password) {
				// c.Redirect("/home/index", http.StatusFound)

				//设置userID为session
				c.SetSession("user", user.ID)

				c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
			} else {
				// errs.Add("default", "name  or passwd error")
				errs.Add("default", "用户名或密码错误")
			}
		} else {
			// fmt.Println(err)
			// errs.Add("default", "name  or passwd error")
			errs.Add("default", "用户名或密码错误")
		}
	}

	c.Data["form"] = form
	c.Data["errors"] = errs
	//定义加载页面
	c.TplName = "auth/login.html"
}

//点击退出销毁session
func (c *AuthController) Logout() {
	c.DestroySession()
	c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
}
