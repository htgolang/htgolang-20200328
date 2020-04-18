package main

import "fmt"

func main() {
BREAK:
	for z := 0; z < 2; z++ {
		fmt.Println(z, "*************")
		for j := 0; j < 3; j++ {
			fmt.Println(j, "-----------")
			for i := 0; i < 5; i++ {
				fmt.Println(i)
				if i == 3 {
					break BREAK
				}
			}
		}
	}
}
