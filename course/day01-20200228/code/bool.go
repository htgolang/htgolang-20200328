package main

import "fmt"

func main() {
	isGirl := false

	fmt.Printf("%T, %#v\n", isGirl, isGirl)

	// 操作
	// 逻辑运算
	a, b, c, d := true, true, false, false
	// 与: 左操作数与右操作数都为true，结果为true &&

	fmt.Println("a, b:", a && b) // true && true : true
	fmt.Println("a, c:", a && c) // true && false : false
	fmt.Println("c, b:", c && b) // false && true : false
	fmt.Println("c, d:", c && d) // false && false : false

	// 或: 左操作数与右操作数只要由一个为true，结果为true ||

	fmt.Println("a, b:", a || b) // true || true : true
	fmt.Println("a, c:", a || c) // true || false : true
	fmt.Println("c, b:", c || b) // false || true : true
	fmt.Println("c, d:", c || d) // false || false : false

	// 非: 取反 true=> false, false => true !
	fmt.Println("a:", !a) // !true : false
	fmt.Println("c:", !c) // !false : true

	// 关系
	fmt.Println(a == b) // true == true : true
	fmt.Println(a != c) // true != false : true
	fmt.Println(a == c) // true == false: false
	fmt.Println(c != b) // false != true : true

	fmt.Printf("%t, %t\n", a, c)

	var bbb bool
	fmt.Println(bbb)
}
