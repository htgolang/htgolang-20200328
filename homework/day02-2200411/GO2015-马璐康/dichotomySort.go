// 自己写的有bug，不能对所有数字都能判断出位置，就走网上参考别人的代码。
package main

import "fmt"

func dichotomyFind(arr *[]int, firstIndex int, lastIndex int, findValue int) {
	if firstIndex > lastIndex {
		fmt.Println("数字不在列表中！")
	}
	// 先找到中间的下标
	midNub := (firstIndex + lastIndex) / 2

	if (*arr)[midNub] > findValue {
		dichotomyFind(arr, firstIndex, midNub-1, findValue)
	} else if (*arr)[midNub] < findValue {
		dichotomyFind(arr, midNub+1, lastIndex, findValue)
	} else {
		fmt.Println("找到选择的数字，下标是：", midNub)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	dichotomyFind(&arr, 0, len(arr)-1, 9)

}

//package main
//
//import "fmt"
//
//func main() {
//	nums := []int{1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15, 16, 17, 18, 19}
//	findNum := 7
//	leftIndex := 0
//	rightIndex := len(nums) - 1
//	fmt.Println(rightIndex)
//	midNum := (leftIndex + rightIndex) / 2
//	fmt.Println(midNum)
//	for i := 1; i < len(nums); i++ {
//		if leftIndex > rightIndex {
//			fmt.Println("没有在找到")
//			return
//		}
//		if nums[midNum] > findNum {
//			fmt.Println("中间的数字为：", nums[midNum], "，输入的数大于实际数字")
//			midNum = (leftIndex + (midNum - 1)) / 2
//			//leftIndex=leftIndex+1
//			fmt.Println("现在的中间数下标为：", midNum)
//		} else if nums[midNum] < findNum {
//			fmt.Println("中间的数字为：", nums[midNum], "，输入的数小于实际数字")
//			midNum = (midNum + 1 + rightIndex) / 2
//			fmt.Println("现在的中间数下标为：", midNum)
//		} else {
//			fmt.Println("恭喜！中间的数字为：", nums[midNum], "，输入的数等于实际数字")
//			fmt.Println("现在的中间数下标为：", midNum)
//			break
//		}
//	}
//
//}
