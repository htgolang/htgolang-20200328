package main

import "fmt"

func main() {
	// 移除切片操作
	// 第一个元素或最后一个元素
	nums := []int{1, 2, 3, 4, 5}

	// start=0 start省略
	// end = len(nums) end省略

	// 移除第一个元素
	nums = nums[1:]
	fmt.Println(nums)

	// 移除最后一个元素
	nums = nums[:len(nums)-1]

	fmt.Println(nums)

	// 移除中间的元素怎么办？
	// [2, 3, 4] 移除3(index)
	// 切片操作, copy
	// [index:] [3, 4]
	// [index+1:] [4]
	// copy([index:], [index+1:]) [2, 4, 4]

	copy(nums[1:], nums[2:])
	nums = nums[:len(nums)-1]
	fmt.Println(nums)
}
