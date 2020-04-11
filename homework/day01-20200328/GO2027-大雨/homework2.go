package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 种子，只需要设置一次
	rand.Seed(time.Now().Unix())
	// rand.Seed(1)
	random := rand.Intn(10)
	//fmt.Println(random)
	fmt.Println("猜数字游戏请输入猜测的数字")
	for i := 1; i <= 5; i++ {
		enternum := 0
		fmt.Scan(&enternum)
		if enternum > random && i < 5 {
			fmt.Println("太大了!")
		}
		if enternum < random && i < 5 {
			fmt.Println("太小了!")
		}
		if enternum == random {
			fmt.Printf("经过%d次对了，太聪明了!", i)
			break
		}
		if enternum != random && i == 5 {
			fmt.Println("提示太笨了，游戏结束!")
		}
	}
}
