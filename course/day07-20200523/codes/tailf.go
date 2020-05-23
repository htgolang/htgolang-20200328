package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func tailf(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	// seek 文件末尾
	file.Seek(0, os.SEEK_END)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(time.Second)
			continue
		}
		fmt.Println(string(line))
	}
}

func main() {

	var path string
	var h, help bool
	flag.StringVar(&path, "p", "tail.log", "path")
	flag.BoolVar(&h, "h", false, "help")
	flag.BoolVar(&help, "help", false, "help")

	flag.Usage = func() {
		fmt.Println("tailf -p path")
		flag.PrintDefaults()
	}

	flag.Parse()
	if h || help {
		flag.Usage()
		os.Exit(0)
	}

	// path => 检查
	// path 输入为空
	// path 文件不存在
	// path 目录
	tailf(path)
}
