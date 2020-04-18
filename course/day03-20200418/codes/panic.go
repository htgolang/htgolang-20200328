package main

import (
	"fmt"
)

func test() (err error) {
	// recover 必须在延迟执行函数内
	defer func() {
		fmt.Println("defer")
		if panicErr := recover(); panicErr != nil {
			err = fmt.Errorf("%s", panicErr)
		}
	}()
	fmt.Println("before")
	panic("自定义panic")
	fmt.Println("after")
	return
}

func main() {
	fmt.Println("before main")
	err := test()
	fmt.Println("after main", err)
}

// raise
// throw
