package main

import (
	"fmt"
	"net"
)

func main() {
	// 网络接口的
	intfs, _ := net.Interfaces()
	for _, intf := range intfs {
		fmt.Println(intf.Index)
		fmt.Println("\t", intf.Name, intf.HardwareAddr, intf.MTU, intf.Flags)
		addrs, _ := intf.Addrs()
		for _, addr := range addrs {
			fmt.Println("\t", addr)
		}
	}

}
