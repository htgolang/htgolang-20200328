package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var (
	fname string
	remotefname string
)

func init() {
	flag.StringVar(&fname,"local","","local filename")
	flag.StringVar(&remotefname,"remote","","remote filename")
}

func main() {
	//fname := "1.txt"
	//remotefname := "1r.txt"
	flag.Parse()
	if flag.NFlag()!=2{
		flag.Usage()
		return
	}
	addr := "127.0.0.1:2222"
	conn, err := net.Dial("tcp", addr)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(fname)
	if err != nil {
		log.Printf("fail to open %s !\n", fname)
		return
	}
	conn.Write([]byte(remotefname+"\n"))
	reader:=bufio.NewReader(conn)
	rspB,_,err:=reader.ReadLine()
	if err==io.EOF{
		log.Println("Connection failed abnormally!")
		return
	}
	if string(rspB)=="FAIL"{
		log.Println("Uploading to remote fileserver failed!")
		return
	}
	if string(rspB)=="OK"{
		freader:=bufio.NewReader(f)
		for {
			rb:=make([]byte,1024)
			n,err:=freader.Read(rb)
			if err==io.EOF{
				log.Println("Uploading finished successfully!")
				break
			}
			conn.Write(rb[:n])
		}
	}
}
