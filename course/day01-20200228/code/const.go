package main

import "fmt"

const (
	packageName string = "package"
	packageMsg         = "package"
)

func main() {
	const name string = "kk"
	const msg = "msg" // 常量可以不使用
	fmt.Println(name)

	// name = "silence" 常量一旦定义不能修改

}
