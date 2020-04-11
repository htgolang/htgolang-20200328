package main

import "fmt"

var version = "1.0"

func main() {
	// 函数内(块)定义的变量必须使用
	/*
		var name string = "kk"
		var msg = "hello world"
		var desc string
	*/

	var (
		name string = "kk"
		msg         = "hello world"
		desc string
	)

	/*
		x := "x"
		y := "y"
	*/
	x, y := "x", "y"
	fmt.Println(name, msg, desc, x, y)
}
