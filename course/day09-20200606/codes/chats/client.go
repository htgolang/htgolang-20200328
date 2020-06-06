package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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

	scanner := bufio.NewScanner(os.Stdin)

	reader := bufio.NewReader(conn)
	for i := 0; i < 4; i++ {
		// 给服务器发送数据
		fmt.Print("请输入消息: ")
		scanner.Scan()
		fmt.Fprintf(conn, "%s\n", scanner.Text())
		// 从服务端读取数据
		line, _, err := reader.ReadLine()

		if err != nil {
			log.Print(err)
			break
		}
		fmt.Printf("回应: %s\n", string(line))
	}
	fmt.Fprintln(conn, "quit")

}
