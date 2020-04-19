package main

import "fmt"

func main(){
	/*
		1. 排序
		2. 去重
		3. 取值
	*/
	var nums = []int{41,88,32,99,32,99,21,13}
	// 排序
	for i:=0;i<len(nums)-1;i++{
		for j:=0;j<len(nums)-1;j++{
			if nums[j] > nums[j+1]{
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Println(nums)
	// 去重
	newNums := []int{nums[0]}
	for i:=1;i<len(nums);i++{
		if newNums[len(newNums)-1] != nums[i]{
			newNums = append(newNums, nums[i])
		}
	}
	fmt.Println(newNums)
	// 取第二大值
	fmt.Println(newNums[len(newNums)-2])
}
