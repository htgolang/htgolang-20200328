package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 重镜像，让浏览器重新发起请求到新的地址上
	http.HandleFunc("/home/", func(resp http.ResponseWriter, req *http.Request) {
		http.Redirect(resp, req, "/login/", 302)
		// fmt.Fprint(resp, "首页")
	})

	http.HandleFunc("/login/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprint(resp, "登录页面")
	})

	http.ListenAndServe(":8888", nil)
}
