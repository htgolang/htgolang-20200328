package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.Gosched() // 让出CPU时间 time.Sleep()

	// var wg *sync.WaitGroup
	wg := new(sync.WaitGroup)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			for ch := 'A'; ch < 'Z'; ch++ {
				fmt.Printf("%d: %c\n", i, ch)
				runtime.Gosched()
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}
