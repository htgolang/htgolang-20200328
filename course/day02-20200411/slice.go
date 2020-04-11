package main

import "fmt"

func main() {
	var names []string

	// 类型
	fmt.Printf("%T\n", names)

	// 零值
	fmt.Printf("%#v\n", names) // nil

	// 初始化
	// 字面量

	// 第一种
	names = []string{} //空切片 已经初始化但是元素数量为0
	fmt.Printf("%#v\n", names)

	names = []string{"05-牛", "25-烽火", "37-海日拉"}

	fmt.Printf("%#v\n", names)

	// 第二种
	names = []string{1: "05-牛", 100: "37-海日啦"}
	fmt.Printf("%#v\n", names)

	// 第三种 make函数

	// make() // 2个参数
	// make() // 3个参数

	names = make([]string, 0) // 申请有5个字符串元素的切片

	fmt.Printf("%#v\n", names)

	names = make([]string, 0, 10)

	fmt.Printf("%#v\n", names)

	names = make([]string, 3)
	names[0] = "a"
	names[1] = "b"
	names[2] = "c"
	// names[3] = "d"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	// fmt.Println(names[3])

	// 获取长度和容量
	fmt.Println(len(names), cap(names))

	// 添加元素
	names = append(names, "d")
	fmt.Printf("%#v\n", names)

	fmt.Println(len(names), cap(names)) // 4 6

	// 遍历
	for i := 0; i < len(names); i++ {
		fmt.Println(i, names[i])
	}

	for i, v := range names {
		fmt.Println(i, v)
	}

	// copy 切片之间的赋值

	aSlice := []string{"a", "b", "c"}
	bSlice := []string{"x", "y"}

	// 长度相等
	// 目的， 源
	copy(aSlice, bSlice)

	fmt.Printf("%#v, %#v\n", aSlice, bSlice)

	// 切片操作 => 数组，切片 => 新生成一个切片

	nums := []int{0, 1, 2, 3, 4, 5} // len = 6, cap = 6
	// nums = make([]int, 6, 10)       // len = 6, cap=10

	//start <= end <= cap  new_cap = cap - start
	numChildren := nums[3:4] //[start: end]
	numChildren = append(numChildren, 100)

	fmt.Printf("%#v, %#v\n", nums, numChildren)

	fmt.Printf("%T, %#v\n", numChildren, numChildren) // [start, end)
	fmt.Println(cap(numChildren))

	// start <= end <= max <= cap, new_cap = max - start
	nums = []int{0, 1, 2, 3, 4, 5}
	numChildren = nums[3:4:4]

	fmt.Println(cap(numChildren))
	numChildren = append(numChildren, 100)
	fmt.Println(cap(numChildren))
	fmt.Printf("%#v, %#v\n", nums, numChildren)

	numArray := [6]int{0, 1, 2, 3, 4, 5}

	// start <= end <= len new_cap = len - start
	arrayChildren := numArray[3:4]
	fmt.Printf("%T, %#v\n", arrayChildren, arrayChildren)
	fmt.Println(cap(arrayChildren))
	arrayChildren = append(arrayChildren, 100)
	fmt.Printf("%#v, %#v\n", arrayChildren, numArray)

	numArray = [6]int{0, 1, 2, 3, 4, 5}

	// start <= end <= len new_cap = len - start
	arrayChildren = numArray[3:4:4]
	fmt.Printf("%T, %#v\n", arrayChildren, arrayChildren)
	fmt.Println(cap(arrayChildren))
	arrayChildren = append(arrayChildren, 100)
	fmt.Printf("%#v, %#v\n", arrayChildren, numArray)
}
