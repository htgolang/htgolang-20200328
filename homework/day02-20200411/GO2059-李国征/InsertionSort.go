package main

import "fmt"

func main() {
	// 插入排序
	// 插入排序从前进行比对， 如果大于前者将替换位置。
	var lst []int = []int{4, 2, 1, 6, 10, 3, 11}
	for j :=0; j<len(lst)-1; j++{
		for index:=0;index <len(lst)-1; index++{
			if index == 0 {
				continue
			} else if lst[index] < lst[index-1]{
				lst[index], lst[index-1] = lst[index-1], lst[index]
			}
		}
	}
	fmt.Println("排序结果:", lst)
}
