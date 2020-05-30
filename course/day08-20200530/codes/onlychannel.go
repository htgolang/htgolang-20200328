package main

import "fmt"

func main() {
	channel := make(chan int, 10)

	channel <- 1
	fmt.Println(<-channel)

	// // 只读
	// var readonly <-chan int
	// // 只写
	// var writeonly chan<- int

	// readonly = channel
	// writeonly = channel

	// 再这个函数中只写
	func(channel chan<- int) {
		channel <- 2
	}(channel)

	// 再这个函数中只读
	func(channel <-chan int) {
		fmt.Println(<-channel)
	}(channel)

}
