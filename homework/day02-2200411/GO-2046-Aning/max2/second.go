package main

import "fmt"

func main() {
	nums := []int{54, 3, 1, 5, 8, 91, 74}
	firstnum, secondnum := nums[0], nums[0]

	for i, v := range nums {
		if v > firstnum {
			secondnum = firstnum
			firstnum = v
		} else {
			secondnum = v
		}
		fmt.Println(i, firstnum, secondnum)
	}
	fmt.Println(firstnum, secondnum)
}
