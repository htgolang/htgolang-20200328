package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 13, 55, 22, 11, 44, 5, 10}

	for index := 0; index < len(nums)-1; index++ {
		for i := 0; i < len(nums)-1-index; i++ {
			if nums[i+1] > nums[i] { //降序
				// if nums[i+1] < nums[i] { //升序
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
		fmt.Printf("第%d次 %v\n", index+1, nums)
	}
	fmt.Println("最后结果:", nums)
}
