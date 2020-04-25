package main

import (
	"log"
	"os"
)

func main() {
	//设置格式的
	// flags
	log.SetFlags(log.Flags() | log.Lshortfile)
	// prefix
	log.SetPrefix("main: ")

	log.Println("我是第一条Println日志")

	// log.Fatalln("我是一个Fatal日志")
	// log.Panicln("我是一条Panic日志")

	log.Println("我是第二条Println日志")

	// DEBUG，INFO, WARNING, ERROR
	// logrus

	logger := log.New(os.Stdout, "logger:", log.Lshortfile|log.Ltime)
	logger2 := log.New(os.Stdout, "logger2:", log.Llongfile|log.Ldate)

	logger.Println("我是logger日志")
	logger2.Println("我是logger2日志")

	// 标准输入/输出 fmt.Scan fmt.Println
	// os.Stdin, os.Stdout, os.Stderr

}
