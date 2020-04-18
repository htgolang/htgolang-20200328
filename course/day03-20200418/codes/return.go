package main

import "fmt"

func add(n1, n2 int) int {
	return n1 + n2
}

func mult(n1, n2 int) int {
	return n1 * n2
}

// 多个返回值
func calc(n1, n2 int) (int, int) {
	// a, b := 1, 2
	r1 := add(n1, n2)
	r2 := mult(n1, n2)
	return r1, r2
}

// 命名返回值

func calc2(n1, n2 int) (r1 int, r2 int) {
	r1 = add(n1, n2)
	r2 = mult(n1, n2)
	return
}

// 合并返回值类型

func calc3(n1, n2 int) (r1, r2 int) {
	r1 = add(n1, n2)
	r2 = mult(n1, n2)
	return
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(calc(1, 2))
	fmt.Println(calc2(2, 3))
	fmt.Println(calc3(2, 3))
}
