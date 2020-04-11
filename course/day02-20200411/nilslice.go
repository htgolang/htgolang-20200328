package main

import "fmt"

func main() {
	var nilSlice []int
	var emptySlice []int = []int{} // emptySlice := []int{}

	fmt.Printf("%T, %#v\n", nilSlice, nilSlice)
	fmt.Printf("%T, %#v\n", emptySlice, emptySlice)

	nilSlice = append(nilSlice, 1)
	emptySlice = append(emptySlice, 1)

	fmt.Printf("%T, %#v\n", nilSlice, nilSlice)
	fmt.Printf("%T, %#v\n", emptySlice, emptySlice)
}
