package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func ls(conn net.Conn) {
	fmt.Fprintf(conn, "ls|0|")
	reader := bufio.NewReader(conn)

	sizeText, err := reader.ReadString('|')
	log.Printf("reade size err: %s", err)

	size, err := strconv.Atoi(sizeText[:len(sizeText)-1]) // 1
	for size > 0 {
		name, err := reader.ReadString(':')
		log.Printf("read filename err: %s", err)
		fmt.Println(name[:len(name)-1])
		size--
	}
}

func cat(conn net.Conn, name string) {
	fmt.Fprintf(conn, "cat|1|%s|", name)
	reader := bufio.NewReader(conn)
	sizeText, err := reader.ReadString('|')
	log.Printf("read size err: %v", err)

	size, err := strconv.Atoi(sizeText[:len(sizeText)-1])
	log.Printf("atoi err: %v", err)
	if size > 0 {
		ctx := make([]byte, size)
		n, err := reader.Read(ctx)
		log.Printf("read content err: %v", err)
		fmt.Printf("文件内容: %s\n", string(ctx[:n]))
	} else {
		fmt.Println("文件内容为空")
	}
}

func main() {
	logfile, _ := os.OpenFile("client.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer logfile.Close()

	log.SetOutput(logfile)

	addr := "127.0.0.1:9999"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Print("connected fileserver")

	scanner := bufio.NewScanner(os.Stdin)
END:
	for {
		fmt.Print("请输入指令:")
		scanner.Scan()
		input := scanner.Text()
		cmds := strings.Split(input, " ")
		switch cmds[0] {
		case "ls":
			ls(conn)
		case "cat":
			cat(conn, cmds[1])
		case "quit":
			fmt.Fprint(conn, "quit|0|")
			break END
		default:
			fmt.Println("输入指令错误")
		}
	}
}
