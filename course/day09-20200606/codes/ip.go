package main

import (
	"fmt"
	"net"
)

func main() {
	// IP结构体
	ip := net.ParseIP("61.135.169.121")
	// 3*16(48) + 13 => 61
	fmt.Printf("%T, %#v\n", ip, ip)
	ip = net.ParseIP("::1")
	fmt.Printf("%T, %#v\n", ip, ip)
	fmt.Println(net.ParseIP("192.168.2.x"))

	hosts, err := net.LookupIP("www.baidu.com")
	fmt.Println(hosts, err)
}
