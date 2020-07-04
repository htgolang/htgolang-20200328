package main

// 导入beego包
import (
	"fmt"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

func main() {

	// 路由 url => controller
	// 处理器函数 处理器: 函数 结构体
	// 函数, 结构体
	// beego/context/Context

	// 绑定函数
	// 以GET方式请求/通过绑定函数处理
	// 固定路由
	beego.Get("/", func(ctx *context.Context) {
		// 用户数据的获取
		name := ctx.Input.Query("name")

		// 给用户响应数据
		ctx.Output.Context.WriteString(fmt.Sprintf("你输入的名字是: %s", name))
	})
	beego.Post("/", func(ctx *context.Context) {
		name := ctx.Input.Query("name")
		ctx.Output.Context.WriteString(fmt.Sprintf("(POST)你输入的名字是: %s", name))

	})
	// beego.Delete
	// beego.Put
	// beego.Options
	// beego.Head
	// beego.Patch

	beego.Any("/any", func(ctx *context.Context) {
		name := ctx.Input.Query("name")
		ctx.Output.Context.WriteString(fmt.Sprintf("(%s)你输入的名字是: %s", ctx.Input.Method(), name))
	})

	// 启动beego程序
	beego.Run()
}
