package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.post
	// 提交数据 请求体
	// 有编码格式
	// application/x-www-form-urlencoded
	// k=v&k2=v2
	// 上传文件 => multipart/form-data
	// application/json {"a" : 1}

	// 有编码格式
	// application/x-www-form-urlencoded

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		//1. 解析参数
		req.ParseForm()
		//2. 获取
		fmt.Println(req.Form) // 包含请求体中和URL中数据
		fmt.Println(req.Form.Get("a"))
		fmt.Println(req.Form["a"])
		fmt.Println(req.PostForm) // 只包含请求体数据
	})

	http.ListenAndServe(":8888", nil)

}
