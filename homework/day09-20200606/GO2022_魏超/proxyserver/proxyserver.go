package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
)

func main() {
	var (
		ip   string
		port int
		url  string
		err  error
		addr string
	)
	flag.StringVar(&ip, "i", "127.0.0.1", "input listen ip.")
	// flag.StringVar(&url, "u", "https://baidu.com", "input proxy server host:ip.")
	flag.StringVar(&url, "u", "127.0.0.1:8080", "input proxy server host:ip.")
	flag.IntVar(&port, "p", 80, "input listen port.")

	// 生成监听地址
	addr = net.JoinHostPort(ip, strconv.Itoa(port))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{}

		url = fmt.Sprintf("%s%s", url, r.URL.RequestURI())
		req, err := http.NewRequest(r.Method, url, r.Body)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}

		// 设置cookie
		for _, c := range r.Cookies() {
			req.AddCookie(c)
		}

		// 设置header
		for key := range r.Header {
			for _, value := range r.Header[key] {
				req.Header.Set(key, value)
			}
		}
		resp, err := client.Do(req)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}
		for key := range resp.Header {
			for _, value := range resp.Header[key] {
				w.Header().Set(key, value)
			}
		}

		body := make([]byte, 0)
		ctx := make([]byte, 1024)
		for {
			n, err := resp.Body.Read(ctx)
			if err != nil {
				if err == io.EOF {
					break
				}
				w.Write([]byte(fmt.Sprintf("%s", err)))
				return
			}
			body = append(body, ctx[:n]...)
		}
		w.Write(body)
		return
	})
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("listen %s failed, %s", addr, err)
	}
}
