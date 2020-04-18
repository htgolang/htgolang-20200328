package main

import (
	"errors"
	"fmt"
	"strconv"
)

func div(n1, n2 int) (int, error) {
	if n2 == 0 {
		return -1, errors.New("除数为0")
	}
	return n1 / n2, nil
}

func main() {
	value, err := strconv.Atoi("xxx")
	fmt.Printf("%T\n", err)
	fmt.Println(err)
	fmt.Println(value)

	e := fmt.Errorf("自定错误")
	fmt.Printf("%T %#v\n", e, e)

	e2 := errors.New("自定错误2")
	fmt.Printf("%T %#v\n", e2, e2)

	// go 语言
	// 希望程序内部如果有错误
	// 通过最后一个返回值显示返回给调用者
	// 由调用者决定如何处理

	if rt, err := div(1, 0); err == nil {
		fmt.Println(rt)
	} else {
		fmt.Println(err)
	}
}
