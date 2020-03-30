package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for i2 := 1; i2 <= i; i2++ {
			if i == i2 {
				fmt.Printf("%3d *%3d = %3d\n", i, i2, i*i2)
			} else {
				fmt.Printf("%3d *%3d = %3d", i, i2, i*i2)
			}
		}
	}
}
