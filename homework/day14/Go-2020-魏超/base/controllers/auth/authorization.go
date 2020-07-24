package auth

import (
	"net/http"

	"cmdb/base/controllers/base"

	"cmdb/models"

	"github.com/astaxie/beego"
)

// Authorization 全局授权
type AuthorizationController struct {
	base.Base
	LoginUser *models.User
}

// Perare 全局授权perare钩子
func (c *AuthorizationController) Perare() {
	// 服务未开启session不做任何检查
	if beego.AppConfig.DefaultBool("session::SessionOn", false) {
		sessionValue := c.GetSession(
			beego.AppConfig.DefaultString("session::Name", "UserID"),
		)
		// 全局检查只检查tick是否为空
		if sessionValue != nil {
			if userID, ok := sessionValue.(int); ok {
				if user := models.GetUserByID(userID); user != nil {
					c.Data["loginUser"] = user
					c.LoginUser = user
					return
				}
			}
		}
		c.Redirect(
			beego.URLFor(beego.AppConfig.DefaultString(
				"auto::LoginAction",
				"AthController.Login",
			),
			),
			http.StatusFound,
		)
	}
}
