package auth

import "github.com/astaxie/beego"

// LayoutController 布局控制器基础
type LayoutController struct {
	AuthorizationController
}

// Prepare
func (c *LayoutController) Prepare() {
	c.AuthorizationController.Prepare()
	c.Layout = "base/layouts/layout.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["SectionStyle"] = ""
	c.LayoutSections["SectionScript"] = ""

	c.Data["title"] = beego.AppConfig.DefaultString("AppName", "CMDB")
}
