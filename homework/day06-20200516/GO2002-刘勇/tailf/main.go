package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

var (
	filename string
	help     bool
)

//文件读取输出
func readfile(name string) {

	f, err := os.Open(name)
	if err != nil {
		if err == err.(*os.PathError) {
			fmt.Println("文件不存在，请检查文件名或路径")
			return
		}
		fmt.Println(err)
		return
	}

	b := bufio.NewReader(f)
	for {
		line, _, err := b.ReadLine()
		if err != nil {
			//当读取到文件末尾则休眠1秒，跳过当前循环继续读取文件
			if err == io.EOF {
				time.Sleep(time.Second * 1)
				continue
			}
			fmt.Println(err)
			return
		}
		//读到数据则输出
		fmt.Println(string(line))
	}
}

//命令行工具，解析输入
func MenuFlag() {
	flag.StringVar(&filename, "n", "", "file name")
	flag.BoolVar(&help, "h", false, "help")

	flag.Usage = func() {
		fmt.Println(`
Usage:gotail -n filename
Options:`)
		flag.PrintDefaults()
	}

	flag.Parse()

	if help || filename == "" {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {

	MenuFlag()

	readfile(filename)

}
