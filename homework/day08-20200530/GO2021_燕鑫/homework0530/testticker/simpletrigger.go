package testticker

import (
	"fmt"
	"time"
)

func HWTrigger() {
	i := 0
	syncChan := make(chan struct{}, 1)
	finishChan := make(chan struct{}, 1)
	intChan := make(chan int, 1)
	triggertime := time.Second
	trigger := time.NewTicker(triggertime)
	go func() {
	LOOP:
		for range trigger.C {
			select {
			case intChan <- 1:
				fmt.Println("SEND: 1")
			case intChan <- 2:
				fmt.Println("SEND: 2")
			case intChan <- 3:
				fmt.Println("SEND: 3")
			case <-finishChan:
				break LOOP
			}
		}
		syncChan<- struct{}{}
		fmt.Println("END. [sender]")
	}()

	var sum int
	for e := range intChan {
		i++
		fmt.Println("Received: ", e)
		sum += e
		if sum > 10 {
			finishChan<- struct{}{}
			break
		}
	}
	<-syncChan
	fmt.Println("End. [receiver]", i)
}
