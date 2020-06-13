package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"rpcserver/service"
)

func main() {
	addr := ":9999"

	// 注册服务 指定服务名称
	rpc.RegisterName("calc", &service.Calculator{})
	// 注册服务 未指定服务名称，默认结构体名
	rpc.Register(&service.Calculator{})

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Printf("[+] listen on: %s", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("[-]error client: %s\n", err.Error())
			continue
		}
		log.Printf("[+] client connected: %s\n", conn.RemoteAddr())

		// 使用例程启动jsonrpc处理客户端请求
		go jsonrpc.ServeConn(conn)
	}
}
