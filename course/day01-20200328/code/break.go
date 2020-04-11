package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break //循环结束
		}
		fmt.Println(i)
	}
}
