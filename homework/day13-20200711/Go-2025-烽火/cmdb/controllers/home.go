package controllers

import (
	"cmdb/base/controllers/auth"
)

type HomeController struct {
	auth.AuthController
}

func (c *HomeController) Index() {
	c.TplName = "home/index.html"
}
