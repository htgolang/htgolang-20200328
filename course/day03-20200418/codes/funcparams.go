package main

import "fmt"

func calc(n1 int, n2 int, callback func(int, int) int) int {
	// 不定义是什么运算
	// 通过函数参数传递我要进行运算
	rt := callback(n1, n2)
	// 检查结果在0, 100范围内，超过 -1
	if rt >= 0 && rt <= 100 {
		return rt
	}
	return -1
}

func add(n1, n2 int) int {
	return n1 + n2
}

func mult(n1, n2 int) int {
	return n1 * n2
}

func main() {
	rt := calc(1, 2, add)
	fmt.Println(rt)

	rt = calc(1, 2, mult)
	fmt.Println(rt)

	rt = calc(50, 30, add)
	fmt.Println(rt)

	rt = calc(50, 30, mult)
	fmt.Println(rt)
}
