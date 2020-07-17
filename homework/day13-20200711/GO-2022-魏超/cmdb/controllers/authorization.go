package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

type authorization struct {
	base
}

func (c *authorization) Perare() {
	if beego.AppConfig.DefaultBool("session::SessionOn", false) {
		tick := c.GetSession(
			beego.AppConfig.DefaultString("session::Name", "TickID"),
		)
		if tick == nil {
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
}
