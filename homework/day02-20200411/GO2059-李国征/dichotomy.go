package main

import (
	"fmt"
	"sort"
)

func main() {
	// 二分法
	// 要求，必须是有序的
	var lst []int = []int{1, 3, 4, 10, 32, 63, 66, 98, 130, 305, 783, 920, 1024}
	// 先进行排序
	sort.Ints(lst)
	insert := 3051
	// 查找中间值，然后匹配想要查询的数值
	for i := 0; i < len(lst); i++ {
		mid := len(lst) / 2
		if lst[mid] > insert {
			lst = lst[:mid]
			if len(lst) == 2 {
				if lst[0] == insert {
					fmt.Printf("查询到的结果:%v %v\n", lst[0], mid)
					break
				} else if lst[1] == insert {
					fmt.Printf("查询到的结果:%v %v\n", lst[0], mid)
					break
				} else {
					fmt.Println("没有查询到")
				}
			}
		} else if lst[mid] < insert {
			lst = lst[mid:]
			if len(lst) == 2 {
				if lst[0] == insert {
					fmt.Printf("查询到的结果:%v\n", lst[mid])
					break
				} else if lst[1] == insert {
					fmt.Printf("查询到的结果:%v\n", lst[mid])
					break
				} else {
					fmt.Println("没有查询到")
				}
			}
		} else if lst[mid] == insert {
			fmt.Printf("结果:%v\n", lst[mid])
			break
		}
	}
}