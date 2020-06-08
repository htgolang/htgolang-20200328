package main
import (
	"fmt"
	
)

func main()  {
	var (
		nums = []int{1,3,4,6,7,9,11,12,17,20}
		target int = 17
		n = len(nums)
	)
	
	left := 0
	right := n-1
	flag := true
	for left < right {
		mid := (left+right)/2
		if nums[mid] > target {
			right = mid -1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			fmt.Println(target,"position is at index : ",mid)
			flag = false
			break
		}
	}
	if flag {
		fmt.Println(target,"not in the array ")
	}
		
	}




