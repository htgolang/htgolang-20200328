package main

import (
	"fmt"
	"net/http"
)

func main() {

	// http协议
	// GET
	// url?params
	// params k=>v&k2=v2

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		//1. 解析参数
		req.ParseForm()
		//2. 获取
		fmt.Println(req.Form)
		fmt.Println(req.Form.Get("a"))
		fmt.Println(req.Form["a"])
	})

	http.ListenAndServe(":8888", nil)
}
