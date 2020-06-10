package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

type Client struct {
	conn net.Conn
	name string
}

var (
	mesChan chan string =make(chan string)
	lock	sync.Mutex
)

var ClientMap	map[string]net.Conn = make(map[string]net.Conn)


func Chat(name string, client net.Conn)  {
	defer client.Close()
	defer func() {
		lock.Lock()
		delete(ClientMap,name)
		lock.Unlock()
	}()

	bufReader := bufio.NewReader(client)
	mesChan <- "欢迎"+name+"加入聊天室"
	for {


		mesg,_,err := bufReader.ReadLine()
		if err != nil{
			log.Println(err)
			mesChan <- name+"离开聊天室"
			break
		}
		//fmt.Println("start  -> chan")
		//fmt.Println("data -> ",string(mesg))
		mesChan <- name+":"+string(mesg)
		//fmt.Println("end  -> chan")
	}
}

func main() {
	listener,err := net.Listen("tcp","127.0.0.1:9999")
	if err !=nil{
		log.Fatal(err)
	}
	defer listener.Close()
	fmt.Println("服务已开始")

	go func() {
		for {
			//fmt.Println("start <- chan")
			mesg := <- mesChan
			//fmt.Println("data <-",mesg)
			//fmt.Println("end <- chan")
			fmt.Println(mesg)
			for _,value := range ClientMap {
				value.Write([]byte(mesg+"\n"))
			}
		}
	}()

	for {
		conn,err := listener.Accept()
		if err != nil{
			log.Println(err)
			continue
		}
		name := make([]byte,1024*1024)
		fullname :=""
OUTNAME:
		for{
			n,err := conn.Read(name)

			fullname = string(name[:n])

			if err != nil{
				fmt.Println(err)
				//fmt.Println("read err")
			}
			//fmt.Println("开始循环")
			//for _,c := range ClientSlice{
			for Name,_ := range ClientMap{

				if Name==fullname{
					//判断重名
					//fmt.Println("重名")
					conn.Write([]byte("1"))
					continue OUTNAME
				}
			}
			fmt.Println("结束循环")
			//判断无重名
			conn.Write([]byte("0"))
			break
		}

		//c :=Client{
		//	conn:conn,
		//	name:fullname,
		//}
		lock.Lock()
		//fmt.Println("加入队列")
		ClientMap[fullname]=conn
		//ClientSlice = append(ClientSlice,c)
		lock.Unlock()
		//读取客户端消息，并写入管道
		//go Chat(fullname,c)
		go Chat(fullname,conn)
	}

}