package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		r := rand.Intn(100)
		guess := 0
		fmt.Println(r)

		for i := 1; i < 7; i++ {
			if i == 6 {
				fmt.Println("太笨了，游戏结束")
				break
			}

			fmt.Print("请在 0 - 100 之间猜测一个数字： ")
			fmt.Scan(&guess)

			if guess == r {
				fmt.Printf("经过 %d 次猜对了，太聪明了\n", i)
				break
			} else if guess < r {
				fmt.Println("猜小了！猜小了！")
			} else {
				fmt.Println("猜大了！猜大了！")
			}

		}
		fmt.Println("")
		fmt.Print("继续游戏？[Y,yes,y]：")
		isQuit := ""
		fmt.Scan(&isQuit)

		if isQuit != "y" && isQuit != "Y" && isQuit != "yes" {
			break
		}
	}

}
