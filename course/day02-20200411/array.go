package main

import "fmt"

func main() {
	var names [3]string
	var signIns [3]bool
	var scores [3]float64

	// 类型
	fmt.Printf("%T\n", names)
	fmt.Printf("%T\n", signIns)
	fmt.Printf("%T\n", scores)

	// 零值
	fmt.Printf("%#v\n", names)
	fmt.Printf("%#v\n", signIns)
	fmt.Printf("%#v\n", scores)

	// 赋值
	// 字面量 => 0, 1, 2, 3, n-1
	// 第一种
	names = [3]string{"05-牛", "37-海日啦", "21-燕鑫"}

	// names = [1]string{"05-牛"}
	// 第二种
	names = [...]string{"05-牛", "37-海日啦", "21-燕鑫"}

	fmt.Printf("%#v\n", names)

	testnames := [...]string{"05-牛", "37-海日啦"}
	fmt.Printf("%T\n", testnames)

	// 第三种
	names = [3]string{1: "kk"} //[3]string{"", "kk", ""}

	fmt.Printf("%#v\n", [3]string{1: "kk"})

	// 操作
	// 关系运算 == !=
	fmt.Println(names == [3]string{})
	fmt.Println(names == [3]string{1: "kk"})

	// 元素
	// 访问 & 修改 索引(0, 1, ..., n - 1)
	fmt.Printf("%q\n", names[0])
	names[0] = "02-牛"
	fmt.Printf("%#v\n", names)

	// 函数 len
	fmt.Println(len(names))

	// 遍历
	for i := 0; i < len(names); i++ {
		fmt.Println(i, names[i])
	}

	for i, v := range names {
		fmt.Println(i, v)
	}

	// 定义一个数组 每个元素也是数组
	// 二维数组
	d2 := [3][2]int{1: [2]int{1, 2}, 0: [2]int{3, 4}, 2: [2]int{1: 5}}
	//[2]int = {0, 0}
	//{[2]int, [2]int, [2]int}
	//{{0, 0}, {0, 0}, {0, 0}}
	fmt.Printf("%#v\n", d2)

	fmt.Printf("%#v\n", d2[0])
	fmt.Printf("%#v\n", d2[0][1])
}
