package auth

import (
	"cmdb/base/controllers/base"
	"fmt"
	"strings"

	"net/http"

	"cmdb/models"

	"github.com/astaxie/beego"
)

type AuthController struct {
	base.BaseController
	LoginUser *models.User
}

func (c *AuthController) getNav() string {
	controllerName, _ := c.GetControllerAndAction()
	return strings.ToLower(strings.TrimSuffix(controllerName, "Controller"))
}

// 用户认证检查
func (c *AuthController) Prepare() {
	sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
	sessionValue := c.GetSession(sessionKey)
	c.Data["loginUser"] = nil
	c.Data["nav"] = c.getNav()

	if sessionValue != nil {
		if pk, ok := sessionValue.(int); ok {
			if user := models.GetUserByPk(pk); user != nil {
				c.Data["loginUser"] = user
				c.LoginUser = user
				return
			}
		} else {
			fmt.Println("ok: ", ok)
		}
	}

	action := beego.AppConfig.DefaultString("auth::LoginAction", "AuthController.Login")
	c.Redirect(beego.URLFor(action), http.StatusFound)
}
