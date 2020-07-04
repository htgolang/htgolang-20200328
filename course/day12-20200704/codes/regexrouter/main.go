package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	// 正则路由
	// URL中定义正则字符串方式进行匹配
	// 匹配 /数字/ 格式的路由 => 并把匹配的值放入到:id参数

	beego.Get("/name/?:name(\\w+)/", func(ctx *context.Context) {
		name := ctx.Input.Param(":name")
		ctx.WriteString(fmt.Sprintf("匹配name: %s", name))
	})

	beego.Get("/id/:id(\\d+)/", func(ctx *context.Context) {
		id := ctx.Input.Param(":id")
		ctx.WriteString(fmt.Sprintf("匹配id: %s", id))
	})

	beego.Get("/any/:content/", func(ctx *context.Context) {
		content := ctx.Input.Param(":content")
		ctx.WriteString(fmt.Sprintf("匹配context: %s", content))
	})

	beego.Get("/file/*", func(ctx *context.Context) {
		splat := ctx.Input.Param(":splat")
		ctx.WriteString(fmt.Sprintf("匹配file: %s", splat))
	})

	beego.Get("/ext/*.*", func(ctx *context.Context) {
		path := ctx.Input.Param(":path")
		ext := ctx.Input.Param(":ext")
		ctx.WriteString(fmt.Sprintf("匹配ext: %s[%s]", path, ext))
	})

	beego.Run()
}
