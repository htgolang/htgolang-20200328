package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var wg sync.WaitGroup
	var counter int64

	var count = 5
	var ceil = 10000

	// a = 0

	// a += 1
	// a += 2

	// 3 1 2

	// 10个例程 5个counter++ 5个count--
	for i := 0; i < count; i++ {
		wg.Add(2)
		go func() {
			for i := 0; i < ceil; i++ {
				// counter++
				atomic.AddInt64(&counter, 1)
				time.Sleep(time.Microsecond)
			}
			wg.Done()
		}()
		go func() {
			for i := 0; i < ceil; i++ {
				// counter--
				atomic.AddInt64(&counter, -1)
				time.Sleep(time.Microsecond)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
