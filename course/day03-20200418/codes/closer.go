package main

import "fmt"

func addBase(base int) func(int) int {
	return func(n int) int {
		return n + base
	}
}

func main() {
	add1 := addBase(1)

	fmt.Println(add1(5))

	add10 := addBase(10)
	fmt.Println(add10(3))
}
