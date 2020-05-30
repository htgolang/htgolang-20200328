package main

import (
	"fmt"
	"time"
)

func main() {

	// time.Tick()

	endTime := time.Now().Add(time.Second * 20)
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	for now := range ticker.C {
		fmt.Println(now)
		if now.After(endTime) {
			break
		}
	}
}
