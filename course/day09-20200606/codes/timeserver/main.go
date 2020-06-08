package main

import (
	"io"
	"net/http"
	"time"
)

// 处理器
// Handler接口
// ServeHTTP(http.ResponseWriter, *http.Request)

type TimeHandler struct {
}

func (h *TimeHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	now := time.Now().Format("2006-01-02 15:04:05")
	io.WriteString(response, now)
}

func main() {
	http.Handle("/time/", &TimeHandler{})

	http.ListenAndServe(":9998", nil)
}
