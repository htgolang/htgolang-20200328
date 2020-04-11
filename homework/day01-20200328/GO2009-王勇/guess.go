package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		guessNum int
		randInt  int
		isGo     string
		isRight  bool
	)
	//设置种子
	rand.Seed(time.Now().UnixNano())
	for {
		isRight = false
		fmt.Println("欢迎来到猜数字游戏，请输入你的答案：")
		randInt = rand.Intn(100)
		fmt.Println(randInt)
		for i := 1; i <= 5; i++ {
			fmt.Scan(&guessNum)
			switch {
			case guessNum == randInt:
				fmt.Println("恭喜您，第", i, "次就答对了！")
				isRight = true
				// break
			case guessNum > randInt:
				fmt.Println("大了，请再次输入答案：")
			case guessNum < randInt:
				fmt.Println("小了，请再次输入答案：")
			}
			//5次用完都没答对的出发提示语句
			if i == 5 && guessNum != randInt {
				fmt.Println("5次机会用完了，很遗憾您都错了")
			}
			//5此以内完成了答案，跳出for循环
			if isRight {
				break
			}
		}
		fmt.Println("是否需要继续玩游戏呢？请输入y/n")
		fmt.Scan(&isGo)
		if isGo == "y" {
			continue
		} else {
			break
		}
	}

}
