package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
v1版本：读取所有内容
*/
var h bool

func traceFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	for {
		line, _ := reader.ReadString('\n')
		if strings.TrimSpace(line) != "" {
			fmt.Println(strings.TrimSpace(line))
		}
	}

}

func IsFileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
	}
}

func main() {
	flag.BoolVar(&h, "h", false, "帮助信息")
	flag.Usage = func() {
		fmt.Println("tailf [file] 查看文件内容")
		flag.PrintDefaults()
	}

	flag.Parse()

	if h || flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	path := flag.Arg(0)
	fmt.Printf("file: %s\n", path)
	if IsFileExists(path) {
		traceFile(path)
	} else {
		fmt.Printf("file:[%s] not find!", path)
		os.Exit(1)
	}
}
