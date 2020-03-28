package main

import "fmt"

func main() {
	var age8 int8 = 31 // 字面量 10 进制 8进制 16进制
	var age = 31
	fmt.Printf("%T, %#v, %d\n", age8, age8, age8)
	fmt.Printf("%T, %#v, %d\n", age, age, age)

	/* 了解
	8进制: 0?? ? < 8
	16进制 0X?? ? < 16 0-9A-F
	abc
	a * base^2 + b * base ^ 1 + c * base ^ 0

	070 => 10进制: 56
	078 => 10进制: 78

	// 二进制存储
	base 2
	7 => 0111 => 1 * 2^2 + 1 * 2 ^1 + 1 * 2 ^ 0
	*/

	fmt.Println(070, 78)

	// 常用操作
	// 算数运算 + - * / % ++ --

	// 4 % 3 = 4 => 3 * 1 + 1
	// 2 % 4 = 2 => 4 * 0 + 2
	//a = b*c + d
	//a % c = d
	a, b := 2, 4
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) // 0  4/3
	fmt.Println(a % b) // 2

	a++
	b--
	fmt.Println(a, b) // 33

	// 关系运算 > < >= <= != ==
	fmt.Println(a > b)  // false
	fmt.Println(a < b)  // false
	fmt.Println(a >= b) // true
	fmt.Println(a <= b) // true
	fmt.Println(a == b) // true
	fmt.Println(a != b) //false

	/*
		仅了解 位运算
		a * base^2 + b * base ^ 1 + c * base ^ 0
		7 => 0111
		2 => 0010

		负数：二进制表示补码 对应正数 取反+1
		-3 3=> 0011 => 1101

		按位与：&, 两个都为1结果为1, 0010 7&2 =>2
		按位或：|, 只要由一个为1结果为1, 0111 7|2 => 7
		取非: ^, 1=>0, 0=>1？ 1101 => -3
		右移位：>> 7 >> 2 0001 => 1
		左移位: << 7 << 2 0000 0000 0000 0111 => 01 1100 =>28
		and not: &^ 7 &^2 0101 => 5
	*/
	fmt.Println(7 & 2)
	fmt.Println(7 | 2)
	fmt.Println(^2)
	fmt.Println(7 >> 2)
	fmt.Println(7 << 2)
	fmt.Println(7 &^ 2)

	var (
		i   int   = 1
		i32 int32 = 1
		i64 int64 = 1
	)

	// 类型转换 type(value) int32(i) int(i32) int64(i) int(i64)
	fmt.Printf("%T\n", i+int(i32))
	fmt.Printf("%T\n", i+int(i64))
	fmt.Printf("%T\n", int32(i)+i32)
	fmt.Printf("%T\n", int64(i)+i64)

	var (
		achar        byte = 'A'
		aint         byte = 65
		unicodePoint rune = '中'

		// 字符串 => 内存(01) => 转换 => 编码(utf8, utf16编码, gbk, gb2312)
	)

	fmt.Println(achar, aint)
	fmt.Println(unicodePoint)

	fmt.Printf("%d %b %o %x %U %c %c", achar, 15, 15, 15, unicodePoint, achar, 65)
}
