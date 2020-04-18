package main

import (
	"fmt"
)

func main() {

	// [168 180 166 176 169]
	// heigh := []int{168, 180, 166, 176, 169}

	// a := 1
	// b := 2
	// a, b := 1, 2

	// a, b = b, a
	// fmt.Println(a, b)

	// 交换
	// tmp := a
	// a = b
	// b = tmp
	// fmt.Println(a, b)
	heigh := []int{168, 180, 166, 170, 169}

	for j := 0; j < len(heigh)-1; j++ {
		for i := 0; i < len(heigh)-1; i++ {
			if heigh[i] > heigh[i+1] {
				heigh[i], heigh[i+1] = heigh[i+1], heigh[i]

			}
			fmt.Println(i, heigh)
		}
	}

	fmt.Println(heigh)

}
