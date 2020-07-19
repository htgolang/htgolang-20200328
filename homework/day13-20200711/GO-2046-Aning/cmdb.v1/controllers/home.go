package controllers

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

//写一个专门做检查的结构体
type AuthorizationController struct {
	beego.Controller
}

//专门做parepare
func (c *AuthorizationController) Prepare() {
	fmt.Println("AuthorizationController.Prepare")
	sessionUser := c.GetSession("user")
	if sessionUser == nil {
		c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
		//停止使用prepare的以下部分
		c.StopRun()
	}
}

type HomeController struct {
	//匿名组合AuthorizationController
	AuthorizationController
	// beego.Controller
}

func (c *HomeController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ctx, controllerName, actionName, app)
	fmt.Println("Init:", controllerName, actionName)
}

// //用home写parpare
// func (c *HomeController) Prepare() {
// 	//session 检查
// 	fmt.Println("prepare")
// 	sessionUser := c.GetSession("user")
// 	if sessionUser == nil {
// 		c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
// 		c.StopRun()
// 	}
// }

func (c *HomeController) Index() {
	// sessionUser := c.GetSession("user")
	// if sessionUser == nil {
	// 	c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
	// 	// c.StopRun()
	// }

	fmt.Println("index")
	c.TplName = "home/index.html"
}

func (c *HomeController) Test() {
	fmt.Println("Test")
	c.Ctx.WriteString("test")
}

func (c *HomeController) Render() error {
	fmt.Println("render")
	return c.Controller.Render()
}

func (c *HomeController) Finish() {
	fmt.Println("finish")
}
