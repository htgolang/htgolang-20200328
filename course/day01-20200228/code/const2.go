package main

import "fmt"

func main() {

	// const B 不允许
	const (
		A = "test"
		B //使用前一个常量的值进行初始化(B)
		C //使用前一个常量的值进行初始化(C=>B)
		D = "testD"
		E //使用前一个常量的值进行初始化(D)
		F //使用前一个常量的值进行初始化(E=>D)
	)

	fmt.Println(B, C)
	fmt.Println(E, F)
}
