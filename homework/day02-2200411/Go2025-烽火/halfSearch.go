package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 10, 5, 7, 22, 9, 11}
	target := 9
	sort.Ints(nums)
	fmt.Println("sorted nums:", nums)

	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := left + ((right - left) >> 1)
		fmt.Println("middle", middle)

		if nums[middle] > target {
			right = middle - 1
			fmt.Println(left, right)
		} else if nums[middle] < target {
			left = middle + 1
			fmt.Println(left, right)
		} else {
			fmt.Println("find target:", target, "index:", middle)
			return
		}
	}
	fmt.Println("not find target:", target)
}
