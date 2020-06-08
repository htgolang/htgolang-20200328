package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)


func Read(conn  net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		line, _, err := reader.ReadLine()

		if err != nil {
			fmt.Println(err)
			conn.Close()
			os.Exit(0)
		}
		fmt.Println(string(line))
	}
}



func main() {
	addr := "127.0.0.1:9999"

	// 连接服务器
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
		conn.Close()
		return
	}
	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("name is :")
	scanner.Scan()
	if scanner.Text() == "" {
		fmt.Println("name is not nil")
		return
	}
	fmt.Fprintf(conn,"%s\n",scanner.Text())
	go Read(conn)
	for {
		scanner.Scan()
		if scanner.Text() == "quit"{
			fmt.Println("退出")
			conn.Close()
			return
		}
		if scanner.Text() == "" {
			fmt.Println("输入不为空")
			continue
		}
		_,err := fmt.Fprintf(conn,"%s\n",scanner.Text())
		if err != nil {
			fmt.Println("服务器已关闭")
			conn.Close()
			return
		}
	}
}
