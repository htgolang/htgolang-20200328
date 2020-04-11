package main

import (
	"fmt"
)

// 包级别
var packageVar string = "package Var"

func main() {
	// 函数级别的
	var funcVar string = "func Var"
	var packageVar string = "func package var"

	{
		var packageVar string = "block package var"
		fmt.Println("1", packageVar)
	}

	fmt.Println("2", packageVar, funcVar)
}
