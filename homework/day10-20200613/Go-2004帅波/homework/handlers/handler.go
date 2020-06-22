package handlers

import (
	"fmt"
	"github.com/astaxie/beego/context"
)

var supportMethod = [6]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}

func RestfulHandler() func(ctx *context.Context) {
	var restfulHandler = func(ctx *context.Context) {
		// 获取隐藏请求
		requestMethod := ctx.Input.Query("_method")
		fmt.Println(requestMethod)
		if requestMethod ==  ""{
			// 正常请求
			requestMethod = ctx.Input.Method()
		}

		// 判断当前请求是否在允许请求内
		flag := false
		for _, method := range supportMethod{
			if method == requestMethod {
				flag = true
				break
			}
		}

		// 方法请求
		if flag == false {
			ctx.ResponseWriter.WriteHeader(405)
			ctx.Output.Body([]byte("Method Not Allow"))
			return
		}

		// 伪造请求方式
		if requestMethod != "" && ctx.Input.IsPost() {
			fmt.Println(requestMethod)
			ctx.Request.Method = requestMethod
		}
	}
	return restfulHandler
}
