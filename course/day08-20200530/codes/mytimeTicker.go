package main

import (
	"fmt"
	"time"
)

func timeTicker(interval time.Duration) <-chan time.Time {
	timeChannel := make(chan time.Time)
	go func() {
		for {
			time.Sleep(interval)
			timeChannel <- time.Now()
		}
	}()

	return timeChannel
}

func main() {

	for now := range timeTicker(time.Second * 3) {
		fmt.Println(now)
	}

}
