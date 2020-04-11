package main

import "fmt"

func main() {
	// 队列
	// 先进先出
	queue := []string{}
	// push
	// append
	queue = append(queue, "a", "b")
	queue = append(queue, "c")

	// pop
	x := queue[0]
	queue = queue[1:]
	fmt.Println("1:", x)

	x = queue[0]
	queue = queue[1:]
	fmt.Println("2:", x)

	x = queue[0]
	queue = queue[1:]
	fmt.Println("3:", x)
}
