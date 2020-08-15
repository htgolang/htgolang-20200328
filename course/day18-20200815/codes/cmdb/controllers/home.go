package controllers

import (
	"cmdb/base/controllers/auth"
)

// HomeController 首页控制器
type HomeController struct {
	auth.LayoutController
}

// Index 首页显示方法
func (c *HomeController) Index() {
	c.TplName = "home/index.html"
}
