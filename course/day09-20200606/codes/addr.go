package main

import (
	"fmt"
	"net"
)

func main() {
	// 获取网络地址
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		fmt.Println(addr.Network(), addr.String())
	}
}
