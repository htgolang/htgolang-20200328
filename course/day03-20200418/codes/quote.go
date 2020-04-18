package main

import "fmt"

// 赋值 <==> 实参/形参

func test1(n int) {
	n = 1
}

func test2(s []int) {
	fmt.Printf("%p\n", s)
	s[0] = 1
}

func main() {
	a := 0
	b := make([]int, 10)
	test1(a)
	test2(b)
	fmt.Println(a)
	fmt.Printf("%p\n", b)
	fmt.Println(b)
	// 在内存中申请内存新的空间，将a的值拷贝到b中
	// 在修改a 不影响 b
	// 在修改a 影响 b

	// 值类型  b = a (31)
	// 引用类型 b = a (地址)
	// b => a 存储内容
	// age := 30
	// tmpAge := age

	// tmpAge = 31
	// fmt.Println(age, tmpAge)
	// fmt.Println("%p, %p\n", &age, &tmpAge)

	// //
	// users := make([]string, 10)
	// tmpUsers := users
	// tmpUsers[0] = "kk"
	// fmt.Printf("%#v, %#v\n", users, tmpUsers)
	// fmt.Printf("%p, %p\n", users, tmpUsers)
	// fmt.Printf("%p, %p\n", &users, &tmpUsers)
	// // 值类型
	// // int, float, point, 数组, 结构体
	// // 引用
	// // 切片，映射，接口

	// array := [3]int{}
	// tmpArray := array
	// tmpArray[0] = 10
	// fmt.Println(array, tmpArray)
}
