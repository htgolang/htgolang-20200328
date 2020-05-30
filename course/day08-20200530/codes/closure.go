package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// wg.Add(10)
	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(i int) {
			fmt.Println(i)

			wg.Done()
		}(i)
	}
	// time.Sleep(time.Second)
	fmt.Println("wait")
	wg.Wait()
	fmt.Println("over")

	// 1. 打印0-9
	// 2. 打印全是0 不可能
	// 3. 打印全是10 (可能性比较大)
	// 随机打印0-9, n个10
}
