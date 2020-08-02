package auth

import (
	"cmdb/base/controllers/base"
	"net/http"

	"github.com/astaxie/beego"
)

//认证需要的控制器
type AuthorizationController struct {
	base.BaseController
}

//prepare 用户认证检查
func (c *AuthorizationController) Prepare() {
	user := c.GetSession(beego.AppConfig.DefaultString("auth::SessionKey", "user"))
	if user == nil {
		c.Redirect(beego.URLFor(beego.AppConfig.DefaultString("auth::LoginController", "AuthController.Login")), http.StatusFound)
	}
}
