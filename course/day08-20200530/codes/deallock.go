package main

func main() {

	// 定义管道 元素为int
	var channel chan int

	channel = make(chan int)

	<-channel
}
