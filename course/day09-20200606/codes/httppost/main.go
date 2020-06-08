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
		fmt.Println(req.PostFormValue("a"))
	})

	http.ListenAndServe(":8888", nil)

}
