package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {

	var wg sync.WaitGroup // 定义等待组结构体变量
	var locker sync.Mutex // 定义互斥锁结构体变量

	var a, b = 10000, 10000

	var count = 1000

	wg.Add(2)
	// a->b
	// 转账100次
	go func() {
		for i := 0; i < count; i++ {
			money := rand.Intn(100)
			if a > money {
				func() {
					locker.Lock()         // 加锁
					defer locker.Unlock() //释放锁
					a -= money
					// 切换例程
					time.Sleep(time.Microsecond)
					b += money
				}()
			}
		}
		wg.Done()
	}()

	// b -> a
	// 转账100次
	go func() {
		for i := 0; i < count; i++ {
			money := rand.Intn(100)
			if b > money {
				func() {
					locker.Lock()         // 加锁
					defer locker.Unlock() // 释放锁
					b -= money
					time.Sleep(time.Microsecond)
					a += money
				}()
			}
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Printf("a: %d, b : %d, total: %d", a, b, a+b)
}
