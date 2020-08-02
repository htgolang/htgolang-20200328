package controllers

import (
	"cmdb/base/controllers/auth"
)

type HomeController struct {
	auth.AuthController
}

// func (c *HomeController) Prepare() {
// 	c.AuthController.Prepare()
// 	c.Data["nav"] = "home"
// }

func (c *HomeController) Index() {
	c.TplName = "home/index.html"
}
