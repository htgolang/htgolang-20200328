package auth

import (
	"cmdb/base/controllers/base"
	"net/http"

	"github.com/astaxie/beego"
)

// AuthorizationController 所有需要认证才能访问的基础控制器
type AuthorizationController struct {
	base.BaseController
}

// Prepare 用户认证检查
func (c *AuthorizationController) Prepare() {
	sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
	user := c.GetSession(sessionKey)
	if user == nil {
		action := beego.AppConfig.DefaultString("auth::LoginAction",
			"AthController.Login")
		c.Redirect(beego.URLFor(action), http.StatusFound)
	}
}
