package main

import "fmt"

func test() (rt string) {
	// defer 函数调用
	// 延迟执行 return调用以后，在函数退出之前
	defer func() {
		fmt.Println("defer")
		rt = "defer" // 写代码注意，在defer中不要修改返回值
	}()

	defer func() {
		fmt.Println("defer A")
	}()

	defer func() {
		fmt.Println("defer B")
	}()
	fmt.Println("test")
	rt = "test"

	return
}

func test2(n1, n2 int) {
	defer func() {
		// 函数体内不管是否发生错误，都会执行
		fmt.Println("test2 defer")
	}()
	fmt.Println("before")
	fmt.Println(n1 / n2)
	fmt.Println("after")
}

func main() {
	fmt.Println(test())

	test2(1, 0)

}
