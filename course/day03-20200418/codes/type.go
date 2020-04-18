package main

import "fmt"

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

func test(a int, b string) int {
	fmt.Println(a, b)
	return 1
}

func main() {
	// test() // 函数调用

	a := test
	b := test(1, "")
	fmt.Printf("%T\n", test)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	a(1, "kk")

	var callback func(int, int) int
	fmt.Printf("%T, %#v\n", callback, callback)

	fmt.Printf("%T\n", sayHello)

	callback = add
	fmt.Printf("%#v\n", callback)
	rt := callback(1, 4)
	fmt.Println(rt)
}
