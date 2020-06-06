package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const dateFormat = "2006-01-02 15:04:05"

func main() {
	addr := "127.0.0.1:9999"

	// 连接服务器
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Printf("connected")

	reader := bufio.NewReader(conn)
	for i := 0; i < 4; i++ {
		// 给服务器发送数据
		fmt.Fprintf(conn, "Time: %s\n", time.Now().Format(dateFormat))
		// 从服务端读取数据
		line, _, err := reader.ReadLine()

		if err != nil {
			log.Print(err)
			break
		}
		log.Printf("服务端响应: %s", string(line))
	}
	fmt.Fprintln(conn, "quit")

}
