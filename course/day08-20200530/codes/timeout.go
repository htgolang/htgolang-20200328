package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	result := make(chan interface{})

	go func() {
		r := rand.Intn(10)
		fmt.Println("timeout:", r)
		time.Sleep(time.Second * time.Duration(r))
		result <- r
	}()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("超时...")
		// 让任务例程结束 context
	case r := <-result:
		fmt.Println("执行成功:", r)
	}
}
