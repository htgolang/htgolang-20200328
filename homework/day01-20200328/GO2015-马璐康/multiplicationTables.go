package main

import "fmt"

func main() {
	for j := 1; j <= 9; j++ {
		for i := 1; i <= j; i++ {
			result := i * j
			fmt.Printf("%d * %d = %-2d ", i, j, result)
		}
		fmt.Println()

	}

}
