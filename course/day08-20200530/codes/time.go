package main

import (
	"fmt"
	"time"
)

func main() {

	endTime := time.Now().Add(time.Second * 20)

	for now := range time.Tick(time.Second * 3) {
		fmt.Println(now)
		if now.After(endTime) {
			break
		}
	}
}
