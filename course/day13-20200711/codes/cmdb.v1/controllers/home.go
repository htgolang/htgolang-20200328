package controllers

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type AuthorizationController struct {
	beego.Controller
}

func (c *AuthorizationController) Prepare() {
	fmt.Println("AuthorizationController.Prepare")
	sessionUser := c.GetSession("user")
	if sessionUser == nil {
		c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
		c.StopRun()
	}
}

type HomeController struct {
	AuthorizationController
}

func (c *HomeController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ctx, controllerName, actionName, app)
	fmt.Println("Init:", controllerName, actionName)
}

func (c *HomeController) Index() {
	fmt.Println("Index")
	// session检查
	c.TplName = "home/index.html"
}

func (c *HomeController) Test() {
	fmt.Println("Test")
	c.Ctx.WriteString("test")
}

func (c *HomeController) Render() error {
	fmt.Println("Render")
	return c.Controller.Render()
}

func (c *HomeController) Finish() {
	fmt.Println("Finish")
}
