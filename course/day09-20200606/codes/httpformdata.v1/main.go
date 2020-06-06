package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// http.post
	// 提交数据 请求体
	// 有编码格式
	// application/x-www-form-urlencoded
	// k=v&k2=v2
	// 上传文件 => multipart/form-data
	// application/json {"a" : 1}

	// 上传文件 => multipart/form-data

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		// 1.解析提交内容
		req.ParseMultipartForm(1024 * 1024)

		fmt.Println(req.MultipartForm) //Value, File
		// url 数据 => Form, FormValue
		// body 值类型 => Form, FormValue, PostForm, PostFormValue, req.MultiprtForm.Value
		// file req.MultiprtForm.File["name"][0].Open()

		// Todo: 参数检查
		file, _ := req.MultipartForm.File["a"][0].Open()
		io.Copy(os.Stdout, file)

	})

	http.ListenAndServe(":8888", nil)

}
