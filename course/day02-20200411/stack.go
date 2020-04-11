package main

import "fmt"

func main() {
	// 堆栈
	// 先进后出

	stack := []string{}
	// push
	// append
	stack = append(stack, "a")
	stack = append(stack, "b")
	stack = append(stack, "c")

	// pop
	// 后面移除
	x := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println("发射:", x)

	x = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println("发射:", x)

	x = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println("发射:", x)
}
