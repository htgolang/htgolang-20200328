package main

import (
	"fmt"
)

// 空接口
type EmptyIntf interface {
}

func PrintType(v interface{}) {
	switch value := v.(type) {
	case int:
		fmt.Println("int", value)
	case bool:
		fmt.Println("bool", value)
	case [3]int:
		fmt.Println("[3]int", value)
	case []int:
		fmt.Println("[]int", value)
	case map[string]string:
		fmt.Println("map[string][string]", value)
	default:
		fmt.Println("unknow", value)
	}
}

func main() {
	var emptyIntf EmptyIntf
	emptyIntf = 1
	emptyIntf = true
	emptyIntf = "test"
	fmt.Println(emptyIntf)

	var emptyInfo2 interface{}
	emptyInfo2 = 1
	emptyInfo2 = "test"
	fmt.Println(emptyInfo2)
	fmt.Println()

	PrintType(1)
	PrintType("test")
	PrintType(false)
	PrintType([3]int{1, 2, 3})
	PrintType([2]int{1, 3})
	PrintType([]int{1, 2, 3})
	PrintType(map[string]string{"kk": "1"})
	PrintType(map[string]int{"kk": 1})
}
