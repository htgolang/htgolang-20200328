package main

import "fmt"

func test(flag bool) int {
	if flag {
		return 1
	}
	fmt.Println("return before")
	return 2
}

func main() {
	fmt.Println(test(true))
	fmt.Println(test(false))
}
