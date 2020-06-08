package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const dateFormat = "2006-01-02 15:04:05"

func main() {
	addr := "127.0.0.1:9999"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Printf("connected")
	fmt.Fprintf(conn, "Time1: %s\n", time.Now().Format(dateFormat))
	fmt.Fprintf(conn, "Time2: %s\n", time.Now().Format(dateFormat))
	fmt.Fprintf(conn, "Time3: %s\n", time.Now().Format(dateFormat))
}
