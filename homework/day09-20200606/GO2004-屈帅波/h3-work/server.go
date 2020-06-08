package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	//"os"

)





type Client struct {
	Conn net.Conn
	Name string
	Addr string
}
var ClientMap map[string]Client = map[string]Client{}

func ToDo(conn net.Conn,msg string) {
	if _,ok := ClientMap[conn.RemoteAddr().String()];!ok{
		return
	}
	name := ClientMap[conn.RemoteAddr().String()].Name
	for _,v := range ClientMap {


		fmt.Fprintf(v.Conn,"%s  : %s \n",name,msg)
		fmt.Println(name,msg)
	}
}

func Quit(conn net.Conn) {
	//如果map里面没有连接信息说明已经关闭
	if _,ok := ClientMap[conn.RemoteAddr().String()];!ok{
		return
	}
	name := ClientMap[conn.RemoteAddr().String()].Name
	for _,v := range ClientMap {
		fmt.Fprintf(v.Conn,"%s is quit home \n",name)
	}
	delete(ClientMap,conn.RemoteAddr().String())
	conn.Close()
}

func main() {
	// 1. 监听服务 Listen/ListenPacket
	// 2. 接收客户端连接 Accept
	// 3. 处理客户端连接(数据交换)
	// 4. 关闭客户端 defer
	// 5. 关闭服务 defer

	addr := "0.0.0.0:9999"

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
		reader := bufio.NewReader(conn)
		line ,_,err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
			return
		}
		if _,ok := ClientMap[ conn.RemoteAddr().String()] ;!ok {
			ClientMap[conn.RemoteAddr().String()] = Client{
				Conn: conn,
				Name: string(line),
				Addr: conn.RemoteAddr().String(),
			}
		}
		fmt.Println(ClientMap)

		go 	func(conn  net.Conn) {
			// 延迟关闭客户端
			defer conn.Close()
			msg := "is add home "
			ToDo(conn,msg)
			//handleConnection(conn,20)
			log.Printf("client[%s] is connected...", conn.RemoteAddr())


			// 从客户端读取数据
			reader := bufio.NewReader(conn)
			//scanner := bufio.NewScanner(os.Stdin)
			for {
				//for _,client := range ClientMap {
				line_two, _, err := reader.ReadLine()
				if err != nil {
					log.Println(err)
					Quit(conn)
					break
				} else {
					if string(line_two) == "quit" {
						Quit(conn)
						break
					}
					msgs := string(line_two)
					ToDo(conn,msgs)
				}
			}
		}(conn)
	}

}
