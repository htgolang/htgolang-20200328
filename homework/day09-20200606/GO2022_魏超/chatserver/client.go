package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

const (
	_ILLEGAL_CHAR_MAP = `\|/'":*? $:`
	MAX_MSG_SIZE      = 1024
)

// 检测输入的名字不能包含特殊字符
// 限制发送数据的大小
// 发送的消息需要进行确认
func main() {
	// 设置日志
	logfile, err := os.OpenFile("client.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("open log file failed, %s\n", err)
		return
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	// 创建与服务端的连接
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Printf("connect chatserver failed, %s", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	fmt.Print("Please input your name:")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		name := scanner.Text()
		if strings.ContainsAny(name, _ILLEGAL_CHAR_MAP) {
			fmt.Printf("Your input name include illegal char [%s]", _ILLEGAL_CHAR_MAP)
			fmt.Print("Please re-enter:")
			continue
		}
		fmt.Fprintf(conn, "%s|", name)
		ackText, err := reader.ReadString('|')
		if err != nil {
			if err == io.EOF {
				fmt.Println("chatserver exit...")
				return
			}
			log.Println("server data receive err, %s", err)
			return
		}
		if ackText[:len(ackText)-1] == "0" {
			fmt.Println("username exist.")
			fmt.Print("Please re-enter!")
			continue
		} else {
			fmt.Println("Sign in Success! [quit]")
			break
		}
	}

	go func() {
		ctx := make([]byte, 1024)
		for {
			n, err := conn.Read(ctx)
			if err != nil {
				log.Printf("receive data failed, %s", err)
				continue
			}
			fmt.Println(string(ctx[:n]))
		}
	}()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		msg := scanner.Text()
		msgSize := len(msg)
		if msgSize > MAX_MSG_SIZE {
			fmt.Println("Data length exceeds the limit.")
			continue
		}
		if strings.ToLower(msg) == "quit" {
			fmt.Fprint(conn, "4|quit")
			fmt.Println("Quit chatserver...")
			break
		}
		fmt.Fprintf(conn, "%d|%s", msgSize, msg)
	}
}
