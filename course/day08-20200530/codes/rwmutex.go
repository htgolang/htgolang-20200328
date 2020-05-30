package main

import (
	"log"
	"sync"
	"time"
)

func main() {

	// a = 1
	// 有: 写 读
	// 读 读
	// 在并发什么场景会出现数据混乱 = 改

	var wg sync.WaitGroup

	// 读写锁
	var locker sync.RWMutex
	// 获取锁 Lock, Rlock
	// 释放锁 Unlock, RUnlock

	// 写 写 Lock Lock 互斥
	// 写 读 Rlock Lock 互斥
	// 读 读 Rlock Rlock

	wg.Add(2)
	go func() {
		log.Println("A: Lock Before")
		locker.RLock()

		log.Println("A: Locked")
		time.Sleep(5 * time.Second)

		locker.RUnlock()
		log.Println("A: Unlocked")
		wg.Done()
	}()

	go func() {
		log.Println("B: Lock Before")
		locker.RLock()
		log.Println("B: Locked")
		time.Sleep(5 * time.Second)

		locker.RUnlock()
		log.Println("B: Unlocked")
		wg.Done()
	}()
	wg.Wait()
}
