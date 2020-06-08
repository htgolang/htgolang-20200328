package main

import "fmt"

func main() {
	nums := []int{10, 2, 6, 3, 9, 1, 34, 7, 4}
	for j := 0; j < len(nums)-1; j++ {
		fmt.Printf("第%d轮\n", j)
		for i := 0; i < len(nums)-j-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
			fmt.Println(nums)
		}
	}
	fmt.Println("最终结果：", nums)

}
