package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func Read(conn net.Conn)  {
	Reader := bufio.NewReader(conn)
	for {
		line,_,err := Reader.ReadLine()
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(string(line))
	}
}

func Write(conn net.Conn,name string)  {
	Reader := bufio.NewReader(os.Stdin)

	for {
		line,_,_ := Reader.ReadLine()
		mesg :=strings.TrimSpace(string(line))
		if  mesg== ""{
			continue
		}else if mesg=="quit"{
			os.Exit(0)
		}

		conn.Write([]byte(mesg+"\n"))
	}
}

func UserInput(mesg string) string  {
	Reader := bufio.NewReader(os.Stdin)
	fmt.Println(mesg)
	line,_,_ :=Reader.ReadLine()
	return string(line)
}

func main() {



	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var name string
	for {
		name = UserInput("请输入你的昵称")
		conn.Write([]byte(name))
		res := make([]byte,1024*1024)
		n,err := conn.Read(res)
		if err !=nil{
			log.Fatal(err)
		}
		resOut :=string(res[:n])
		if resOut == "0"{
			break
		}else  {
			fmt.Println("昵称已被占用")
			continue
		}
	}
	fmt.Println("准备开始聊天吧")
	go Read(conn)
	for{
		Write(conn,name)
	}

}