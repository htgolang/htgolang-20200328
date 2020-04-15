package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 4, 7, 8, 7, 6, 5, 8}
	sort.Ints(nums)
	var second int
	max := nums[len(nums)-1]
	index := sort.SearchInts(nums, max)

	// 并列最大
	if index == len(nums)-1 {
		second = nums[len(nums)-2]
	} else if index < len(nums)-1 {
		second = nums[index-1]
	}
	fmt.Println("并列第一的情况下，second Max:", second)

	// 不并排
	if index == len(nums)-1 {
		second = nums[len(nums)-2]
	} else if index <= len(nums)-2 {
		second = max
	}
	fmt.Println("不并列第一的情况下，second Max:", second)

}
