package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// 1. 定义处理器/处理器函数
	// 2. 绑定URL 处理器/处理器函数关系
	// 3. 启动web服务

	/*
		1. 处理器函数
		参数: http.ResponseWriter, *http.Request
	*/
	timeFunc := func(response http.ResponseWriter, request *http.Request) {
		fmt.Println(request)

		now := time.Now().Format("2006-01-02 15:04:05")
		// io.WriteString(response, now)
		fmt.Fprint(response, now)

	}

	/*
		2.绑定URL关系
		http.HandleFunc(path, 处理器函数)
	*/

	http.HandleFunc("/time/", timeFunc)
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		ctx, err := ioutil.ReadFile("index.html")
		if err == nil {
			response.Write(ctx)
		} else {
			fmt.Fprint(response, "欢迎")
		}
	})

	/*
		3. 启动web服务
		http.ListenAndServe(addr, nil)
	*/
	http.ListenAndServe(":9999", nil)

}
