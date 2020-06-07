package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:8888"
	listener, err := net.Listen("tcp", addr)
	defer listener.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listen on: [%s]\n", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			//log.Fatal(err)
			log.Println(err)
			continue
		}
		log.Printf("[%v] is connected!\n", conn.RemoteAddr())
		go recvRequest(conn)
	}
}

func recvRequest(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	requestData := make([]byte, 0)
	for {
		rb := make([]byte, 1024)
		n, err := reader.Read(rb)
		if err == io.EOF {
			log.Printf("[%v] has left!\n", conn.RemoteAddr())
			break
		}
		requestData = append(requestData, rb[:n]...)
	}

	fmt.Println(string(requestData))
}
