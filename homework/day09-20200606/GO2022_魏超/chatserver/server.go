package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
)

type ConnChat struct {
	connMap map[string]net.Conn
	lock    sync.RWMutex
}

func (this *ConnChat) AddConn(name string, conn net.Conn) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, ok := this.connMap[name]
	if ok {
		return false
	}
	this.connMap[name] = conn
	return true
}

func (this *ConnChat) DelConn(name string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.connMap, name)
}

func (this ConnChat) Broadcast(name, msg string) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	for user := range this.connMap {
		log.Println(msg)
		if user != name {
			fmt.Fprint(this.connMap[user], msg)
		}
	}
}

func (this ConnChat) Len() int {
	return len(this.connMap)
}

var connChat = &ConnChat{
	connMap: make(map[string]net.Conn),
}

func main() {
	logfile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("open log file failed, %s\n", err)
		return
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	addr := "0.0.0.0:9999"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("listen chat server failed, %s\n", err)
		return
	}
	defer listener.Close()
	log.Printf("listener on :[%s]", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept client failed, %s\n", err)
			continue
		}
		log.Printf("accept %s", conn.RemoteAddr().String())
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	var (
		name     string
		err      error
		n        int
		sizeText string
		msgSize  int
	)

	ctx := make([]byte, 1024)

	// 针对加入的用户设置新的名字
	for {
		nameText, err := reader.ReadString('|')
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Printf("get client name failed, %s", err)
		}
		name = nameText[:len(nameText)-1]
		if connChat.AddConn(name, conn) {
			fmt.Fprint(conn, "1|")
			fmt.Printf("[%s] join...\n", name)
			break
		}
		fmt.Fprint(conn, "0|")
	}
	connChat.Broadcast(name, fmt.Sprintf("<welcome, %s join chat room!>", name))
	defer func() {
		connChat.Broadcast(name, fmt.Sprintf("<%s leave chat room...>", name))
		connChat.DelConn(name)
	}()
	for {
		sizeText, err = reader.ReadString('|')
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Printf("get [%s] msg size failed, %s", name, err)
			return
		}
		msgSize, err = strconv.Atoi(sizeText[:len(sizeText)-1])
		if err != nil {
			fmt.Printf("[%s] protocol error convert mesg size failed, %s", name, err)
			continue
		}
		n, err = reader.Read(ctx)
		if err != nil {
			fmt.Printf("receive [%s] info failed, %s", name, err)
			continue
		}
		if n != msgSize {
			fmt.Printf("[%s] send info incomplete", name)
			continue
		}
		// 当接收到用户的输入的信息是quit，直接退出监听
		if string(ctx[:n]) == "quit" {
			break
		}
		connChat.Broadcast(name, fmt.Sprintf("[%s] %s", name, string(ctx[:n])))
	}

}
