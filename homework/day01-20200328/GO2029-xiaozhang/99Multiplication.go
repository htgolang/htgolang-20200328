package main 

import "fmt"

func main () {
	for n := 1; n < 10; n++ {
		for m := 1; m <= n; m++ {
			fmt.Printf(" %d, %d, %d", n, m, n*m)
		}
		fmt.Println("")
	}
}
