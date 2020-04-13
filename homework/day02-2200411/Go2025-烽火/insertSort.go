package main

import (
	"fmt"
	"strings"
)

/*
插入排序：
	以第一个元素为准，后面的元素依次与之比较，进行值交换， 循环len(n)-1
	第二次，再以第二个元素为准，后面的与之比较, 循环len(n)-1-i
	...
*/

func main() {
	nums := []int{10, 4, 2, 8, 3, 5, 9}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] { //升序
				nums[i], nums[j] = nums[j], nums[i]
			}
			// fmt.Println(i, j, nums)
		}
		fmt.Println(i, nums)
		fmt.Println(strings.Repeat("--", 20))
	}
	fmt.Println(nums)
}
