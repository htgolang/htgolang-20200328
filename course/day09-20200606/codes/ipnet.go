package main

import (
	"fmt"
	"net"
)

func main() {
	//IPNet IP范围，cidr 掩码 127.0.0.1/24 127.0.0.0 - 127.0.0.255 127.0.0.*
	ip, ipnet, err := net.ParseCIDR("192.168.1.1/24")
	fmt.Println(ip, ipnet, err)

	fmt.Println(ipnet.Contains(ip))
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.2.1")))
	fmt.Println(ipnet.Network())

}
