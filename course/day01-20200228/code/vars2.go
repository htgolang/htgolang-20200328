package main

import "fmt"

func main() {
	var name string = "kk" //定义了类型并且初始化了值

	var zeroString string // 定义变量类型，但不初始化值
	// 初始化使用类型对应的零值（空字符串""）

	var typeString = "kk" // 定义变量省略类型, 不能省略初始化值
	// 通过对应的值类型推到变量的类型

	// 短声明(必须在函数内包含函数内子块使用， 不能再包级别使用)
	shortString := "kk"
	// 通过对应的值类型推到变量的类型

	fmt.Println(name, zeroString, typeString, shortString)
}
