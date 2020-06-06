package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 当请求URL未绑定, 按照URL中最近匹配的绑定关系去处理
	/*
		/ indexHandleFunc
		/time/ timeHandleFunc

		/time/ => timeHandleFunc
		/time/xxx/xxx => /time/ => timeHandleFunc

		/abc/abc/ => / => indexHandleFunc
	*/

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println(req.Method, req.URL, req.Proto)

		fmt.Printf("%T, %#v\n", req.Header, req.Header)
		fmt.Println(req.Header.Get("User-Agent"))
	})

	http.ListenAndServe(":8888", nil)
}
