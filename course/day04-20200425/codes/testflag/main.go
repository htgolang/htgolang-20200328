package main

import (
	"flag"
	"fmt"
)

// 命令行中的参数进行解析 => 指定的命令参数 -> 变量
func main() {
	var (
		host string
		port int
		h    bool
		help bool
	)
	// -H host -P port
	// -h --help

	// 解析变量的指针, 命令行中指定的参数名, 默认值, 帮助
	flag.StringVar(&host, "H", "127.0.0.1", "连接地址")
	flag.IntVar(&port, "P", 22, "连接端口")
	flag.BoolVar(&h, "h", false, "帮助")
	flag.BoolVar(&help, "help", false, "帮助")

	flag.Usage = func() {
		fmt.Println("usage: testflag [-H 127.0.0.1] [-P 22]")
		flag.PrintDefaults()
	}

	flag.Parse() // 解析

	if h || help {
		flag.Usage()
		return
	}

	fmt.Println(host, port, h, help)
	fmt.Println(flag.NArg())

	fmt.Printf("%#v\n", flag.Args())
}
