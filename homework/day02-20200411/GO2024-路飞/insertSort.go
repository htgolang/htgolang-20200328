package main

import "fmt"

func main(){
	/*
		3, 6, 2, 8, 5
		A  B  C  D  E   位置

		安排6的位置
			6>3, 6放到3的后面，nums[A+1]=6		3 6 。。。
		安排2的位置
			2<6, 6后移一位，nums[B+1]=nums[B]	3 6 6 。。
			2<3, 3后移一位，nums[A+1]=nums[A]	3 3 6 。。
			nums[0]=2						   2 3 6 。。
		安排8的位置
			8>6, 8放到6的后面, nums[C+1]=nums[C]	2 3 6 8 。
			break
		安排5的位置
			5<8, 8后移一位，nums[D+1]=nums[D]	2 3 6 8 8
			5<6, 6后移一位, nums[C+1]=nums[C]	2 3 6 6 8
			5>3, 5放到3的后面，nums[B+1]=5		2 3 5 6 8
	*/
		var nums = []int{33, 65, 21, 89, 12, 23, 99, 41}
		var tmp int
		for i:=1;i<=len(nums)-1;i++{
			tmp = nums[i]
			for j:=i-1;j>=0;j--{
				if tmp < nums[j]{
					nums[j+1] = nums[j]
					if j==0{
						nums[j] = tmp
					}
				}else if tmp >= nums[j]{
					nums[j+1] = tmp
					break
				}
			}
		}
		fmt.Println(nums)
}
