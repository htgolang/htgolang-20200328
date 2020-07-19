package controllers

import (
	"cmdb/base/controllers/auth"
)

type HomeController struct {
	auth.AuthorizationController
}

func (c *HomeController) Index() {
	c.TplName = "home/index.html"
}
