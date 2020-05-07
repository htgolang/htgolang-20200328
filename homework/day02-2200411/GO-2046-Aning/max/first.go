package main

import "fmt"

func main() {
	nums := []int{1, 50, 20, 30, 45, 19, 32}
	maxnum := nums[0]
	for i, v := range nums {
		if v > maxnum {
			maxnum = v
		}
		fmt.Println(i, ":", maxnum)
	}
	fmt.Println(maxnum)
}
