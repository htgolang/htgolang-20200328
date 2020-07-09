package main

import "fmt"

func main() {
	// 冒泡排序
	var bs []int = []int{3, 5, 1, 65, 4, 10}
	for j:=0;j<len(bs)-1;j++{
		for i:=0; i<len(bs)-1; i++{
			if bs[i] > bs[i+1]{
				bs[i], bs[i+1] = bs[i+1], bs[i]
			}
		}
	}
	fmt.Println("排序后结果:", bs)
}
