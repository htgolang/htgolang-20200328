package main

import (
	"fmt"
	"time"
)

func timeAfter(interval time.Duration) <-chan time.Time {
	timeChannel := make(chan time.Time)
	go func() {
		time.Sleep(interval)
		timeChannel <- time.Now()
	}()

	return timeChannel
}

func main() {

	fmt.Println(time.Now())
	// 延迟3s 延迟interval后执行一次
	fmt.Println(<-time.After(time.Second * 3))

	fmt.Println(time.Now())
	fmt.Println(<-timeAfter(time.Second * 3))

}
