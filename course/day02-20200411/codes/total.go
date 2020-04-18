package main

import "fmt"

func main() {
	total := 0
	index := 0
	max := 100

START:
	index += 1
	total += index
	if index == max {
		goto END // 已经加到100 结束
	}
	goto START // 循环
END:
	fmt.Println(total)
}
