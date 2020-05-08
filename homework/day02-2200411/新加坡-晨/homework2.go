package main
import (
	"fmt"
	"sort"
)


func main()  {
	var (
		nums = []int{4,2,6,7,9,5,3,5,9,8,8,10}
	)
	sort.Ints(nums)
	fmt.Println(nums)
	max := nums[len(nums)-1]
	for i:= len(nums)-2 ; i>=0;i-- {
		if nums[i] != max{
			fmt.Println("2nd largest number is : ",nums[i])
			break
		} 
	}



}