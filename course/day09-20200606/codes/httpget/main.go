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
		fmt.Println(req.FormValue("a"))
	})

	http.ListenAndServe(":8888", nil)
}
