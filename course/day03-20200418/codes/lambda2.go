package main

import "fmt"

func main() {
	func() {
		fmt.Println("call")
	}()
	fmt.Println("1")

	func(i int) {
		fmt.Println(i)
	}(100)
}
