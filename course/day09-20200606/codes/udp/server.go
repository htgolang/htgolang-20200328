package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := ":9888"

	packetConn, err := net.ListenPacket("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer packetConn.Close()

	ctx := make([]byte, 1024)

	for {
		n, addr, err := packetConn.ReadFrom(ctx)
		if err != nil {
			log.Printf("read from err: %s", err)
			continue
		}
		fmt.Println(addr)
		fmt.Println(string(ctx[:n]))
		packetConn.WriteTo([]byte("xxxxxxx"), addr)
	}

}
