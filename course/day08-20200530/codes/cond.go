package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	var locker sync.Mutex
	cond := sync.NewCond(&locker)

	wg.Add(2)
	go func() {
		fmt.Println("Cond Wait")
		cond.L.Lock()
		defer cond.L.Unlock()
		// 判断是否满足条件，如果不满足
		cond.Wait()
		fmt.Println("Cond After")
		wg.Done()
	}()

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("notice")
		cond.Broadcast() // 广播条件
		wg.Done()
	}()

	wg.Wait()

}
