package main

import (
	"fmt"
)

func main() {
	var height float32 = 1.68
	var heightType = 1.68 // 1e3 1 * 10 ^ 3

	fmt.Printf("%T %#v %f\n", height, height, height)
	fmt.Printf("%T %#v %f\n", heightType, heightType, heightType)

	var k = 1e3
	fmt.Println(k)
	// 操作
	// 算术运算: + - * / ++ --
	var (
		f1 = 1.2
		f2 = 2.36
	)

	fmt.Println(f1 + f2)
	fmt.Println(f1 - f2)
	fmt.Println(f1 * f2)
	fmt.Println(f1 / f2)

	f1++
	f2--

	fmt.Println(f1, f2)

	// 关系运算 > < >= <=  如果想要判断== != 判断差值再一定区间范围内
	fmt.Println(f1 > f2)
	fmt.Println(f1 >= f2)
	fmt.Println(f1 < f2)
	fmt.Println(f1 <= f2)

	// 赋值运算 =  += -= /= *=
	f1 = 1.32
	f1 += f2 //f1 = f1 + f2

	fmt.Println(f1)

}