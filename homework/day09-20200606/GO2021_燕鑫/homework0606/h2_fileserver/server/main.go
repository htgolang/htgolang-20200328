package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	addr := ":2222"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listen on: [%s]\n", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("[%v] is connected!\n", conn.RemoteAddr())
		go func(conn net.Conn) {
			defer conn.Close()
			reader := bufio.NewReader(conn)
			filenameB, _, err := reader.ReadLine()
			if err == io.EOF {
				log.Printf("[%v] has left!\n", conn.RemoteAddr())
				return
			}
			f, err := os.OpenFile(strings.TrimSpace(string(filenameB)), os.O_CREATE|os.O_TRUNC, os.ModePerm)
			defer f.Close()
			if err != nil {
				conn.Write([]byte("FAIL\n"))
				log.Println("Terminate connection!")
				return
			}
			conn.Write([]byte("OK\n"))
			writer := bufio.NewWriter(f)
			defer writer.Flush()
			for {
				rb := make([]byte, 1024)
				n, err := reader.Read(rb)
				if err == io.EOF {
					log.Printf("[%v] has left!\n", conn.RemoteAddr())
					return
				}
				writer.Write(rb[:n])
			}
		}(conn)
	}
}
