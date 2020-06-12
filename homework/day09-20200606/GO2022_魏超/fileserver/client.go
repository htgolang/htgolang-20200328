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
	logfile, err := os.OpenFile("client.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("open logfile faild,", err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	addr := "127.0.0.1:9999"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal("connect server faild, ", err)
	}
	defer conn.Close()
	log.Println("connected fileserver")

	scanner := bufio.NewScanner(os.Stdin)
END:
	for {
		fmt.Print("plese input command:")
		scanner.Scan()
		input := scanner.Text()
		cmds := strings.Split(input, " ")
		switch cmds[0] {
		case "ls":
			ls(conn)
		case "cat":
			if len(cmds) < 2 {
				fmt.Println("cat FILENAME")
			} else {
				cat(conn, cmds[1])
			}
		case "delete":
			if len(cmds) < 2 {
				fmt.Println("delete FILENAME")
			} else {
				delete(conn, cmds[1])
			}
		case "download":
			if len(cmds) < 3 {
				fmt.Println("download REMOTEFILE LOCALFILE")
			} else {
				download(conn, cmds[1], cmds[2])
			}
		case "upload":
			if len(cmds) < 3 {
				fmt.Println("upload LOCALFILE REMOTEFILE")
			} else {
				upload(conn, cmds[1], cmds[2])
			}
		case "quit":
			fmt.Fprint(conn, "quit|0|")
			break END
		default:
			fmt.Println("请入指令错误.")
		}
	}
}

func ls(conn net.Conn) {
	fmt.Fprintf(conn, "ls|0|")
	reader := bufio.NewReader(conn)
	sizetext, err := reader.ReadString('|')
	if err != nil {
		log.Println(err)
	}
	size, err := strconv.Atoi(sizetext[:len(sizetext)-1])
	if err != nil {
		log.Println(err)
	}
	for size > 0 {
		name, err := reader.ReadString(':')
		if err != nil {
			log.Printf("read filename err: %s", err)
		}
		fmt.Println(name[:len(name)-1])
		size--
	}
}

func cat(conn net.Conn, name string) {
	var sendSize int
	fmt.Fprintf(conn, "cat|1|%s|", name)
	reader := bufio.NewReader(conn)
	sizeText, err := reader.ReadString('|')
	if err != nil {
		log.Println(err)
	}
	fileSize, err := strconv.Atoi(sizeText[:len(sizeText)-1])
	if err != nil {
		log.Println(err)
	}
	sizeText, err = reader.ReadString('|')
	if err != nil {
		log.Println(err)
	}
	sliceSize, err := strconv.Atoi(sizeText[:len(sizeText)-1])
	if err != nil {
		log.Println(err)
	}

	if fileSize > 0 {
		fmt.Print("文件内容:")
		for fileSize > sendSize {
			ctx := make([]byte, sliceSize)
			n, err := reader.Read(ctx)
			if err != nil {
				log.Println(err)
				break
			}
			fmt.Printf("%s", string(ctx[:n]))
			sendSize += n
		}
		fmt.Print("\n")
	} else {
		fmt.Println("文件内容为空")
	}
}

func delete(conn net.Conn, name string) {
	fmt.Fprintf(conn, "delete|1|%s|", name)
	reader := bufio.NewReader(conn)
	result, err := reader.ReadString('|')
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(result[:len(result)-1])
}

func download(conn net.Conn, remotefile, localfile string) {
	file, err := os.Create(localfile)
	if err != nil {
		log.Printf("%s don't openf file, %s", localfile, err)
		return
	}
	defer file.Close()
	fmt.Fprintf(conn, "download|1|%s|", remotefile)
	reader := bufio.NewReader(conn)
	sizeText, err := reader.ReadString('|')
	if err != nil {
		log.Println(err)
	}
	fileSize, err := strconv.Atoi(sizeText[:len(sizeText)-1])
	if err != nil {
		log.Println(err)
	}
	sizeText, err = reader.ReadString('|')
	if err != nil {
		log.Println(err)
	}
	sliceSize, err := strconv.Atoi(sizeText[:len(sizeText)-1])
	if err != nil {
		log.Println(err)
	}

	ctx := make([]byte, sliceSize)
	if fileSize > 0 {
		for fileSize > 0 {
			var n int
			md5Text, err := reader.ReadString('|')
			if err != nil {
				log.Println(err)
				break
			}
			if fileSize < sliceSize {
				n, err = reader.Read(ctx)
			} else {
				n, err = io.ReadFull(reader, ctx)
			}
			if err != nil {
				log.Println(err)
				break
			}

			m := md5.Sum(ctx[:n])
			if md5Text[:len(md5Text)-1] != hex.EncodeToString(m[:]) {
				fmt.Println("data md5 failed.")
				break
			}
			_, err = file.Write(ctx[:n])
			if err != nil {
				log.Println(err)
				break
			}
			fileSize -= n
		}
	} else {
		fmt.Println("文件内容为空")
	}
}

func upload(conn net.Conn, localfile, remotefile string) {
	sliceSize := 1024
	file, err := os.Open(localfile)
	if err != nil {
		log.Printf("%s don't openf file, %s", localfile, err)
		return
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		log.Printf("%s stat failed, %s", localfile, err)
		return
	}
	fileSize := fileinfo.Size()
	reader := bufio.NewReader(conn)
	fmt.Fprintf(conn, "upload|1|%s|", remotefile)
	fmt.Fprintf(conn, "%d|%d|", fileSize, sliceSize)

	ctx := make([]byte, sliceSize)
	for fileSize > 0 {
		n, err := file.Read(ctx)
		if err == io.EOF {
			break
		}
		m := md5.Sum(ctx[:n])

		n1, err := fmt.Fprintf(conn, "%s|%s", hex.EncodeToString(m[:]), string(ctx[:n]))
		log.Println("n1:", n1)
		log.Println("err:", err)
		fileSize -= int64(n)
	}

	ack, _ := reader.ReadString('|')
	if len(ack) < 1 || ack[:len(ack)-1] == "1" {
		fmt.Println("上传失败")
		return
	}
	fmt.Println("上传成功")
}
