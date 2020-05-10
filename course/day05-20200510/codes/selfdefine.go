package main

import "fmt"

type Counter int

type Task map[string]string

type Callback func(...string)

func main() {

	// var name Type
	var cnt Counter
	fmt.Printf("%T\n", cnt)
	fmt.Printf("%#v\n", cnt)

	cnt = 1
	fmt.Printf("%#v\n", cnt)

	// var total int = 100
	// fmt.Println(total / cnt)

	// var task map[string]string
	var task Task
	fmt.Printf("%T\n", task)
	fmt.Printf("%#v\n", task)
	task = map[string]string{"name": "完成Todo管理"}

	fmt.Printf("%#v\n", task)

	// 函数类型
	var print Callback

	print = func(args ...string) {
		for i, v := range args {
			fmt.Println(i, v)
		}

	}

	print("a", "b", "c")
}
