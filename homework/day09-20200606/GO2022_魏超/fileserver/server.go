package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	addr := "127.0.0.1:9999"
	logfile, err := os.OpenFile("fileserver.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("open logfile faild,", err)
		return
	}
	log.SetOutput(logfile)

	// 启动监听服务
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("server runing failed, %s\n", err)
		return
	}

	// 延迟关闭
	defer listener.Close()
	log.Printf("listen on: %s", addr)

	for {
		// 接受客户端请求
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("err accept: %s", err)
			continue
		}

		// 处理客户端链接
		log.Printf("client connected: %s", conn.RemoteAddr())
		go handleConn(conn)
	}
}

func cmd(conn net.Conn) (string, []string) {
	reader := bufio.NewReader(conn)
	op, err := reader.ReadString('|')
	if err != nil {
		log.Printf("read op error: %s\n", err)
	}
	cntText, err := reader.ReadString('|')
	if err != nil {
		log.Printf("read cnt error: %s\n", err)
	}
	if cntText == "" {
		return "quit", []string{}
	}
	cnt, err := strconv.Atoi(cntText[:len(cntText)-1])
	if err != nil {
		log.Printf("atoi error: %s\n", err)
	}
	args := make([]string, 0, cnt)
	for cnt > 0 {
		param, _ := reader.ReadString('|')
		args = append(args, param[:len(param)-1])
		cnt--
	}
	return op[:len(op)-1], args
}

func handleConn(conn net.Conn) {
	defer conn.Close()
END:
	for {
		op, args := cmd(conn)
		log.Printf("op: %s, args: %#v\n", op, args)
		switch op {
		case "ls":
			ls(conn)
		case "cat":
			cat(conn, args[0])
		case "delete":
			delete(conn, args[0])
		case "download":
			download(conn, args[0])
		case "upload":
			upload(conn, args[0])
		case "quit":
			break END
		}
	}
	log.Printf("Client closed: %s", conn.RemoteAddr())
}

func ls(conn net.Conn) {
	file, err := os.Open(".")
	if err != nil {
		log.Printf("open fill err,%s", err)
	}
	defer file.Close()
	names, err := file.Readdirnames(-1)
	if err != nil {
		log.Printf("cmd format err,%s", err)
	}
	suffix := ""
	if len(names) > 0 {
		suffix = ":"
	}
	fmt.Fprintf(conn, "%d|%s%s", len(names), strings.Join(names, ":"), suffix)
}

func cat(conn net.Conn, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		fmt.Fprint(conn, "0|0|")
		return
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		log.Println(err)
		fmt.Fprint(conn, "0|0|")
		return
	}
	fileSize := fileinfo.Size()
	sliceSize := 10
	fmt.Fprintf(conn, "%d|%d|", fileSize, sliceSize)
	ctx := make([]byte, sliceSize)
	for fileSize > 0 {
		n, _ := file.Read(ctx)
		fmt.Fprint(conn, string(ctx[:n]))
		fileSize -= int64(n)
	}
}

func delete(conn net.Conn, name string) {
	fileinfo, err := os.Stat(name)
	if err != nil {
		log.Printf("file stat failed, %s", err)
		fmt.Fprint(conn, "File information acquisition failed.|")
		return
	}
	if fileinfo.IsDir() {
		fmt.Fprint(conn, "Cann't delete dir.|")
	} else {
		err := os.Remove(name)
		if err != nil {
			log.Panicf("delete %s failed", name)
			fmt.Fprint(conn, "Delete file failed.|")
		} else {
			fmt.Fprint(conn, "Delete file success.|")
		}

	}
}

func download(conn net.Conn, name string) {
	sliceSize := 1024
	file, err := os.Open(name)
	if err != nil {
		log.Printf("file open failed, %s", err)
		fmt.Fprint(conn, "0|0|")
		return
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		log.Printf("file stat failed, %s", err)
		fmt.Fprint(conn, "0|0|")
		return
	}
	fileSize := fileinfo.Size()
	fmt.Fprintf(conn, "%d|%d|", fileSize, sliceSize)
	ctx := make([]byte, sliceSize)
	for fileSize > 0 {
		n, _ := file.Read(ctx)
		m := md5.Sum(ctx[:n])
		fmt.Fprintf(conn, "%s|%s", hex.EncodeToString(m[:]), string(ctx[:n]))
		fileSize -= int64(n)
	}
}

func upload(conn net.Conn, name string) {
	isErr := false
	file, err := os.Create(name)
	if err != nil {
		log.Printf("file open failed, %s", err)
		fmt.Fprint(conn, "0|")
		return
	}
	defer file.Close()
	reader := bufio.NewReader(conn)
	sizeText, err := reader.ReadString('|')
	if err != nil {
		log.Println(err)
		isErr = true
	}
	fileSize, err := strconv.Atoi(sizeText[:len(sizeText)-1])
	if err != nil {
		log.Println(err)
		isErr = true
	}
	sizeText, err = reader.ReadString('|')
	if err != nil {
		log.Println(err)
		isErr = true
	}
	sliceSize, err := strconv.Atoi(sizeText[:len(sizeText)-1])
	if err != nil {
		log.Println(err)
		isErr = true
	}
	ctx := make([]byte, sliceSize)

	for isErr != true && fileSize > 0 {
		var n int
		md5Text, err := reader.ReadString('|')
		if err != nil {
			log.Println(err)
			isErr = true
			break
		}
		if fileSize < sliceSize {
			n, err = reader.Read(ctx)
		} else {
			n, err = io.ReadFull(reader, ctx)
		}
		if err != nil {
			log.Println(err)
			isErr = true
			break
		}

		m := md5.Sum(ctx[:n])
		if md5Text[:len(md5Text)-1] != hex.EncodeToString(m[:]) {
			log.Println("data md5 failed.")
			isErr = true
			break
		}

		_, err = file.Write(ctx[:n])
		if err != nil {
			log.Println("data write file failed")
			isErr = true
			break
		}
		fileSize -= n
	}

	if isErr {
		fmt.Fprint(conn, "1|")
	} else {
		fmt.Fprint(conn, "0|")
	}

}
