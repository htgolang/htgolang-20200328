package main

import "fmt"

// n个盘子 start(开始) end(终点) temp(借助)

// n  start -> temp -> end
// n - 1 start -> end -> temp
//
// start -> end
// n - 1 temp -> start -> end

// 终止条件 1 -> start -> end

func tower(start string, end string, temp string, layer int) {
	if layer == 1 {
		fmt.Println(start, "->", end)
		return // 无返回值，表示函数结束
	}
	tower(start, temp, end, layer-1)
	fmt.Println(start, "->", end)
	tower(temp, end, start, layer-1)
}

func main() {
	tower("塔1", "塔3", "塔2", 5)
}
