package main

import (
	"fmt"
	"sort"
)

func main() {
	// 选取切片中第二大元素
	// 1、不去重 选取第二大元素
	// 2、去重 选取第二大元素
	var lst []int = []int{1, 2, 4, 5, 5}
	fmt.Printf("列表:%v\n", lst)
	sort.Ints(lst)

	// 第一种情况
	fmt.Printf("不去重，提取第二大元素 %#v\n", lst[len(lst)-2])

	// 第二种情况
	if lst[len(lst)-1] == lst[len(lst)-2]{
		fmt.Printf("去重，提取第二大元素 %#v\n", lst[len(lst)-3])
	} else {
		fmt.Printf("去重，提取第二大元素 %#v\n", lst[len(lst)-2])
	}
}
