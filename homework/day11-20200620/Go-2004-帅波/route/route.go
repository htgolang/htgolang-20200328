package route

import (
	"github.com/astaxie/beego"
	"github.com/strive-after/go-kubernetes/controls"
	"github.com/astaxie/beego/context"
)

//这里切记contexy不是go自带的包 这里需要导入	"github.com/astaxie/beego/context"
var filtFunc = func(ctx *context.Context) {
	url := ctx.Request.URL.Path
	//当请求路径是login 和reg 的时候跳过不做过滤函数检查
	if url == "/login" || url == "/register" {
		return
	}
	useremail  := ctx.GetCookie("UserEmail")
	user := ctx.Input.Session(useremail)
	if user == nil {
		ctx.Redirect(301,"/login")
	}
}

func init()  {
	//过滤函数
	//访问这些页面得时候 需要通过session做判断没有登陆 无法访问后台页面
	beego.InsertFilter("/*",beego.BeforeRouter,filtFunc)
	//beego.InsertFilter("/user/show",beego.BeforeRouter,handlers.RestfulHandler())


	//运维平台总页面
	beego.Router("/",&controls.LoginControllers{},"get:Operation")
	//登陆页面
	beego.Router("/login",&controls.LoginControllers{},"get:LoginGet;post:LoginPost")
	//注册页面
	beego.Router("/register",&controls.RegisterControllers{},"get:RegGet;post:RegPost")
	//退出登陆删除session
	beego.Router("/LoginOut",&controls.LoginControllers{},"get:LoginOut")
}


