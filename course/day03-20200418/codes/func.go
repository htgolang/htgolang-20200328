package main

import "fmt"

// 导入包(标准包，自定义包，第三方包)

// 包级别的变量，常量，函数

// 无参 无返回值
func sayHello() {
	fmt.Println("Hello World")
}

// 有参 无返回值
func sayHi(name string, name2 string) {
	fmt.Println("Hi:", name, name2)
}

// 有参，有返回值
func add(n1 int, n2 int) int {
	return n1 + n2
}

func test(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	// sayHello
	// main -> func -> ... -> sayHello

	// 调用 方法名()
	sayHello()

	// 调用 方法名(参数, )
	sayHi("金哥", "魏超")

	var n int
	n = add(1, 2)
	fmt.Println(n)

	test(1, "kk")

}
