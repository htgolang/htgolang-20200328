package main

import (
	"flag"
	"fmt"
)

// 命令行中的参数进行解析 => 指定的命令参数 -> 变量
func main() {
	// -H host -P port
	// -h --help

	// 解析变量的指针, 命令行中指定的参数名, 默认值, 帮助
	host := flag.String("H", "127.0.0.1", "连接地址")
	port := flag.Int("P", 22, "连接端口")
	h := flag.Bool("h", false, "帮助")
	help := flag.Bool("help", false, "帮助")

	flag.Usage = func() {
		fmt.Println("usage: testflag [-H 127.0.0.1] [-P 22]")
		flag.PrintDefaults()
	}

	flag.Parse() // 解析
	fmt.Printf("%T, %T, %T, %T\n", host, port, h, help)

	if *h || *help {
		flag.Usage()
		return
	}

	fmt.Println(*host, *port, *h, *help)
	fmt.Println(flag.NArg())

	fmt.Printf("%#v\n", flag.Args())
}
