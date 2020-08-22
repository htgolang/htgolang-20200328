package auth

import (
	"strings"

	"github.com/astaxie/beego"
)

// LayoutController 布局控制器基础
type LayoutController struct {
	AuthorizationController
}

func (c *LayoutController) GetNav() string {
	controllerName, _ := c.GetControllerAndAction()
	return strings.ToLower(strings.TrimSuffix(controllerName, "Controller"))
}

// Prepare
func (c *LayoutController) Prepare() {
	c.AuthorizationController.Prepare()
	c.Layout = "base/layouts/layout.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["SectionStyle"] = ""
	c.LayoutSections["SectionScript"] = ""

	c.Data["nav"] = c.GetNav()
	c.Data["subnav"] = ""
	c.Data["title"] = beego.AppConfig.DefaultString("AppName", "CMDB")
}
