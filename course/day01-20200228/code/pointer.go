package main

import "fmt"

func main() {
	// 零值 nil
	var (
		pointerInt    *int
		pointerString *string
	)

	fmt.Printf("%T %#v\n", pointerInt, pointerInt)
	fmt.Printf("%T %#v\n", pointerString, pointerString)

	// 赋值
	// 使用现有变量 取地址 &name
	age := 32

	age2 := age

	fmt.Printf("%T, %#v\n", &age, &age)
	fmt.Printf("%T, %#v\n", &age2, &age2)
	pointerInt = &age

	fmt.Println(pointerInt)
	fmt.Println(*pointerInt)

	*pointerInt = 330000
	fmt.Println(age, age2)

	pointerString = new(string)
	fmt.Printf("%#v, %#v\n", pointerString, *pointerString)

	pp := &pointerString
	fmt.Printf("%T\n", pp)
	**pp = "kk"
	fmt.Println(*pointerString)
}
