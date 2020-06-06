package main

import "fmt"

func main(){
	/*
		[4, 7, 2, 1, 9]
		 A  B  C  D  E   位置
		第1轮：
			A B比较，不交换，4,7,2,1,9
			B C比较，交换，4,2,7,1,9
			C D比较，交换，4,2,1,7,9
			D E比较，不交换，4,2,1,7,9
		第2轮：
			A B比较，
			B C比较，
			C D比较，
			D E比较，
		第3轮：
			A B比较，
			B C比较，
			C D比较，
			D E比较，
		第4轮：
			A B比较，
			B C比较，
			C D比较，
			D E比较，

		len()-1 轮，每轮比较len()-1次
	*/

	var nums = []int{33, 65, 21, 89, 12, 23, 99, 41}
	for i:=0;i<len(nums)-1;i++{
		for j:=0;j<len(nums)-1;j++{
			if nums[j] > nums[j+1]{
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Println(nums)
}
