package main

import "fmt"

// 参数类型合并
// 连续多个变量类型相同
// 保留最后一个元素类型，前面的类型都可以省略
// func add(n1 int, n2 int) int {
func add(n1, n2 int) int {
	return n1 + n2
}

func test(p1, p2 string, p3, p4 int, p5, p6 bool) {
	fmt.Printf("%T, %T, %T, %T, %T, %T\n", p1, p2, p3, p4, p5, p6)
	fmt.Println(p1, p2, p3, p4, p5, p6)
}

func main() {
	fmt.Println(add(1, 2))
	test("", "", 0, 0, false, false)
}
