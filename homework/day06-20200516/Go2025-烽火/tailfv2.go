package main

import (
	"github.com/hpcloud/tail"

	"flag"
	"fmt"
	"os"
)

/*
v2版本：引用hp tail包
*/
var h bool

func traceFile(path string) {
	t, err := tail.TailFile(path, tail.Config{
		ReOpen: true,
		Follow: true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 0},
		MustExist: true,
		Poll: true,
	})
	if err != nil {
		fmt.Println(err)
	}
	for line := range t.Lines {
    	fmt.Println(line.Text)
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
	traceFile(path)
}
