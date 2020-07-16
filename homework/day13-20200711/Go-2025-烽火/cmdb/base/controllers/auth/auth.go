package auth

import (
	"cmdb/base/controllers/base"

	"net/http"

	"github.com/astaxie/beego"
)

type AuthController struct {
	base.BaseController
}

// 用户认证检查
func (c *AuthController) Prepare() {
	sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")

	if sessionUser := c.GetSession(sessionKey); sessionUser == nil {
		action := beego.AppConfig.DefaultString("auth::LoginAction", "AuthController.Login")
		c.Redirect(beego.URLFor(action), http.StatusFound)
	}
}
