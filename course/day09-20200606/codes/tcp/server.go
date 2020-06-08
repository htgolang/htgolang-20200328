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
	// 1. 监听服务 Listen/ListenPacket
	// 2. 接收客户端连接 Accept
	// 3. 处理客户端连接(数据交换)
	// 4. 关闭客户端 defer
	// 5. 关闭服务 defer

	addr := "127.0.0.1:9999"

	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal(err)
	}

	// 延迟关闭
	defer listener.Close()
	log.Printf("listen on: [%s]", addr)

	// 循环接受
	for {
		// 接收客户端连接
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		func() {
			// 延迟关闭客户端
			defer conn.Close()
			log.Printf("client[%s] is connected...", conn.RemoteAddr())
			// 从客户端读取数据
			reader := bufio.NewReader(conn)
			for {
				line, _, err := reader.ReadLine()
				if err != nil {
					log.Println(err)
					break
				} else {
					if string(line) == "quit" {
						break
					}
					log.Printf("接收到数据: %s", string(line))
					// 回复数据
					fmt.Fprintf(conn, "Over: %s\n", time.Now().Format(dateFormat))
				}
			}
		}()
	}

}
