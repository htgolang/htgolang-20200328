package auth

import (
	"cmdb/base/controllers/base"
	"cmdb/base/response"
	"fmt"

	"github.com/astaxie/beego"
)

type APIController struct {
	base.BaseController
}

func (c *APIController) Prepare() {
	c.EnableXSRF = false // 针对Controller关闭XSRF检查

	token := fmt.Sprintf("Token %s", beego.AppConfig.DefaultString("api::token", ""))
	headerToken := c.Ctx.Input.Header("Authorization")

	fmt.Println(token)
	fmt.Println(headerToken)

	if token != headerToken {
		c.Data["json"] = response.Unauthorization
	}
}

func (c *APIController) Render() error {
	c.ServeJSON()
	return nil
}
