package main

import "fmt"

func main() {

	var (
		index = 0
		total = 0
	)

	for {
		total += index
		index++
		if index > 100 {
			break
		}
	}

	fmt.Println(total)
}
