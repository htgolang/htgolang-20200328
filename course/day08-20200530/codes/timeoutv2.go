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
	timeout := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 5)
		timeout <- struct{}{}
	}()

	go func() {
		r := rand.Intn(10)
		fmt.Println("timeout:", r)
		time.Sleep(time.Second * time.Duration(r))
		result <- r
	}()

	select {
	case <-timeout:
		fmt.Println("超时...")
		// 让任务例程结束 context
	case r := <-result:
		fmt.Println("执行成功:", r)
	}
}
