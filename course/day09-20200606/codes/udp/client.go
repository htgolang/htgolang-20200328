package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:9888")

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "Time: %s", time.Now().Format("2006-01-02"))
	ctx := make([]byte, 1024)
	n, _ := conn.Read(ctx)

	fmt.Println(string(ctx[:n]))
}
