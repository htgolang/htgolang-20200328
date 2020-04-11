package main

import (
	"fmt"
)

// 包级别
var packageVar string = "package Var"

func main() {
	// 函数级别的
	var funcVar string = "func Var"
	{
		// 块级别
		var blockVar string = "block Var"

		{ // 限定变量的使用范围

			// 子块级别
			var innerBlockVar string = "inner block var"
			fmt.Println(packageVar, funcVar, blockVar, innerBlockVar)
		}

		fmt.Println(packageVar, funcVar, blockVar)
	}

	fmt.Println(packageVar, funcVar)
}
