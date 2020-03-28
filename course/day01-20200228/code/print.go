package main

import "fmt"

func main() {
	var name = "kk"

	fmt.Println("*")
	fmt.Println(name) // kk 一行 打印变量加换行
	fmt.Println("*")
	fmt.Print(name) // kk* 一行 只打印变量不加换行
	fmt.Println("*")

	fmt.Printf("%T, %v, %#v", name, name, name)
}
