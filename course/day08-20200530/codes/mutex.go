package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var locker sync.Mutex

	wg.Add(2)
	go func() {
		log.Println("A: Lock Before")
		locker.Lock()

		log.Println("A: Locked")
		time.Sleep(5 * time.Second)

		locker.Unlock()
		log.Println("A: Unlocked")
		wg.Done()
	}()

	go func() {
		log.Println("B: Lock Before")
		locker.Lock()
		log.Println("B: Locked")
		time.Sleep(5 * time.Second)

		locker.Unlock()
		log.Println("B: Unlocked")
		wg.Done()
	}()
	wg.Wait()
}
