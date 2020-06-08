package main

import (
	"fmt"
)

func main()  {
	var(
		nums = []int{4,1,6,3,2,9,10,5,7}

	)
	for i:=1;i<len(nums);i++{   // starts with index 1
		value := nums[i]
		idx := i
		for idx > 0 && value<nums[idx-1]{
			idx = idx -1
			nums[i] = nums[i-1]
		}
		nums[idx] = value
	}
	fmt.Println(nums)
}