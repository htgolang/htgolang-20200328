package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	var filename string
	flag.StringVar(&filename,"f","","文件路径")
	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()
	if filename == ""{
		flag.PrintDefaults()
	}
	_,err := os.Stat(filename)
	if os.IsNotExist(err){
		fmt.Println("文件不存在")
	}
	file,err := os.OpenFile(filename,os.O_RDONLY,os.ModePerm)
	defer file.Close()

	buffReader :=bufio.NewReader(file)
	for {
		line,err := buffReader.ReadString('\n')
		if err != nil{
			if err == io.EOF{
				time.Sleep(time.Second)
			}else {
				fmt.Println(err)
				break
			}
		}
		fmt.Print(string(line))
	}

}
