package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("for before", i)
		// 打开文件
		// 延迟关闭
		// 处理(处理出现错误)
		defer func(i int) {
			fmt.Println("defer", i)
		}(i)

		fmt.Println("for after", i)
	}

	fmt.Println("main")
}
