package main

import (
	"fmt"
)


func SecondMaxNum(nums []int) int{
	if len(nums) == 0 {
		panic("xxxx")
	}

	if len(nums) == 1 {
		return nums[0]
	}

	var max, secondmax int
	if nums[0] > nums[1] {
		max = nums[0]
		secondmax = nums[1]
	} else {
		max = nums[1]
		secondmax = nums[0]
	}

	for i := 2;i < len(nums);i++ {
		if nums[i] > secondmax {
			if nums[i] < max {
				secondmax = nums[i]
			} else if nums[i] == max {  //有并列第一时，跳出该次判断，继续下次判断
				continue
			} else {
				secondmax,max = max,nums[i]
			}
		}
	}

	return secondmax
}



func SecondMaxNum2(nums []int) int{
	if len(nums) == 0 {
		panic("xxxx")
	}

	if len(nums) == 1 {
		return nums[0]
	}

	var max, secondmax int
	if nums[0] > nums[1] {
		max = nums[0]
		secondmax = nums[1]
	} else {
		max = nums[1]
		secondmax = nums[0]
	}

	for i := 2;i < len(nums);i++ {
		if nums[i] > secondmax {
			if nums[i] <= max {
				secondmax = nums[i]
			} else {
				secondmax,max = max,nums[i]
			}
		}
	}

	return secondmax
}



func BubbleSort(nums []int) []int {
	for i:=0;i<len(nums);i++ {
		for j := 0;j<len(nums);j++ {
			if nums[i] > nums[j] {
				nums[j],nums[i] = nums[i],nums[j]
			}
		}
	}

	return nums
}


func InsertSort(nums []int) []int {
	if len(nums) <= 1 {
		return nil
	}

	for i := 1;i < len(nums);i++ {
		tmp := nums[i]
		key := i -1

		for key >= 0 && tmp < nums[key] {
			nums[key+1] = nums[key]
			key--
		}
		//fmt.Println(nums)
		if key + 1 != i {
			nums[key+1] = tmp
		}
	}

	return nums
}


func main() {
	var nums []int = []int{1,2,3,4,5,101,5,100,2,10,101,46,100}



    Secondmax := SecondMaxNum(nums)
    fmt.Println("第二大的数是：",Secondmax)

    Secondmax = SecondMaxNum2(nums)
    fmt.Println("第二大的数是：",Secondmax)

    BubbleSort(nums)
	fmt.Println(nums)

    InsertSort(nums)
    fmt.Println(nums)
}
