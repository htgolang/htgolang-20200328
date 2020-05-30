package main

import (
	"fmt"
	"time"
)

func chars(prefix string) {
	for c := 'A'; c <= 'Z'; c++ {
		fmt.Printf("%s: %c\n", prefix, c)
		time.Sleep(time.Microsecond)
	}
}

func main() {

	// go 函数调用
	go chars("gorouting")

	// 主例程
	chars("main")

	// 休眠
	// time.Sleep(time.Second * 3)

	// 1. 只打印main
	// (*)2. 乱序打印main和gorouting main A-Z， gorouting不一定
	// 3. 乱序打印main和gorouting main 都A-Z
}
