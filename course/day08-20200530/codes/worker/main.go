package main

import (
	"fmt"
	"sync"
	"worker/pool"
)

func main() {
	worker := pool.NewPool(2)

	createTask := func(i int) func() interface{} {
		return func() interface{} {
			return i
		}
	}

	for i := 0; i < 5; i++ {
		worker.AddTask(createTask(i))
	}

	worker.Start()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for result := range worker.Results {
			fmt.Println(result)
		}
		wg.Done()

	}()
	worker.Wait()

	wg.Wait()
}
