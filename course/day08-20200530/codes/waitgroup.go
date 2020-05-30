package main

import (
	"fmt"
	"sync"
	"time"
)

func mainChars(prefix string) {
	for c := 'A'; c <= 'Z'; c++ {
		fmt.Printf("%s: %c\n", prefix, c)
		time.Sleep(time.Microsecond)
	}
}

func chars(prefix string, wg *sync.WaitGroup) {
	for c := 'A'; c <= 'Z'; c++ {
		fmt.Printf("%s: %c\n", prefix, c)
		time.Sleep(time.Microsecond)
	}
	wg.Done()
}

func main() {

	// 定义等待组结构体的变量（计数信号量）

	// 定义值类型，在goroutine之间传递需要传递指针
	// 定义指针类型
	var wg sync.WaitGroup

	// Add(n) 添加n信号
	// Done() 处理完成一个信号
	// Wait() 等待计数器归零， 当所有信号都处理完成就结束, 当信号量没有处理完成就等待

	// go 函数调用
	wg.Add(1)
	go chars("gorouting", &wg)

	// 主例程
	mainChars("main")

	fmt.Println("wait")
	wg.Wait() // 程序阻塞，等待信号量全部处理完成
	fmt.Println("over")

}
