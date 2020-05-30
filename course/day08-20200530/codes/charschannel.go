package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)
	fmt.Println("start")

	for i := 0; i < 2; i++ {
		go func(prefix int) {
			for ch := 'A'; ch <= 'Z'; ch++ {
				fmt.Printf("%d: %c\n", prefix, ch)
			}
			channel <- prefix // 例程结束后写入管道
		}(i)
	}

	// 管道中读取3次数据证明三个例程执行结束
	fmt.Println("before")
	for i := 0; i < 3; i++ {
		fmt.Println("over:", <-channel)
	}

	fmt.Println("over")
	time.Sleep(time.Second * 5)
}
