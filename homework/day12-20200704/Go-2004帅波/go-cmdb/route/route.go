package route

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/strive-after/go-cmdb/controls"
)

//这里切记contexy不是go自带的包 这里需要导入	"github.com/astaxie/beego/context"
var filtFunc = func(ctx *context.Context) {
	url := ctx.Request.RequestURI
	//当请求路径是login 和reg 的时候跳过不做过滤函数检查
	if url == "/auth/login" || url == "/auth/reg" {
		return
	}
	useremail,_  := ctx.GetSecureCookie(controls.Secret,"UserEmail")
	user := ctx.Input.Session(useremail)
	if user == nil {
		ctx.Redirect(301,"/auth/login")
	}
}


func init()  {
	//过滤函数
	//访问这些页面得时候 需要通过session做判断没有登陆 无法访问后台页面
	beego.InsertFilter("/*",beego.BeforeRouter,filtFunc)
	//beego.InsertFilter("/user/show",beego.BeforeRouter,handlers.RestfulHandler())

	// 登陆页面包含 reg跟login
	beego.AutoRouter(&controls.AuthController{})
	//运维平台总页面
	beego.Router("/",&controls.Operation{})
	//用户页面
	beego.AutoRouter(&controls.UserController{})
}



