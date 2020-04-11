package main

import "fmt"

func main() {
	/*
		Print multiplication table, like this:
		1 * 1 = 1
		2 * 1 = 2 2 * 2 = 4
		3 * 1 = 3 3 * 2 = 6 3 * 3 = 9
		...
		9 * 1 = 9 9 * 2 = 18 9 * 3 = 27 .......9 * 9 = 81
	*/

	for row := 1; row <= 9; row++ {
		for col := 1; col <= row; col++ {
			fmt.Printf("%d * %d = %2d\t", row, col, row*col)
		}
		fmt.Println()
	}
}
