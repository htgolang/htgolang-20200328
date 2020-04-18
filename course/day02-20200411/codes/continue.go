package main

import "fmt"

func main() {
CONTINUE:
	for j := 0; j < 3; j++ {
		fmt.Println(j, "-----------")
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			if i == 3 {
				continue CONTINUE
			}
		}
	}
}
