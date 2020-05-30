package main

import (
	"fmt"
	"time"
)

func main() {

	// 定义管道 元素为int
	var channel chan int

	fmt.Printf("%T %#v\n", channel, channel)

	// 初始化(无缓冲区的管道)
	channel = make(chan int)
	fmt.Printf("%T %#v\n", channel, channel)

	go func() {

		fmt.Println("before <-")
		// 写
		channel <- 1 // 需要有另外的例程进行读取, 在不能写入数据时会进行阻塞
		fmt.Println("after <-")
	}()

	fmt.Println("before")

	// 读
	time.Sleep(time.Second * 5)
	num := <-channel // 需要有另外的例程进行写入, 在没有读到数据会进行阻塞
	fmt.Println(num)
	time.Sleep(time.Second * 5)
}
