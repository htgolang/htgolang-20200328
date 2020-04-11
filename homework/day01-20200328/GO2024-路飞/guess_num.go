package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		Guess a number between 0 and 100, 5 times chance.
	*/
	fmt.Println("***** Now let's guess a number between 0 and 100 *****")
	rand.Seed(time.Now().Unix())
	randInt := rand.Intn(100) // 生成随机数

	var chance int = 5 // 猜数的机会

	var guessNum int
	fmt.Print("<<< Input your guess: ")

	for i := 1; i <= chance; i++ {
		fmt.Scan(&guessNum)
		switch {
		case guessNum < 0: // 判断是否输入了负数
			fmt.Println(">>> Your guess should be in a 0 to 100 range.")
			// 判断本次是否是最后一次机会，若是，则结束游戏，否则继续猜数。
			if i < chance {
				fmt.Print("<<< Try again: ")
				continue
			} else if i == chance {
				fmt.Printf(">>> Game over, the answer is %d.", randInt)
				break
			}
		case guessNum > randInt:
			fmt.Println(">>> Bigger.")
			if i < chance {
				fmt.Print("<<< Try again: ")
				continue
			} else if i == chance {
				fmt.Printf(">>> Game over, the answer is %d.", randInt)
				break
			}
		case guessNum < randInt:
			fmt.Println(">>> Smaller.")
			if i < chance {
				fmt.Print("<<< Try again: ")
				continue
			} else if i == chance {
				fmt.Printf(">>> Game over, the answer is %d.", randInt)
				break
			}
		case guessNum == randInt:
			fmt.Println(">>> Bingo! You get it.")
			break
		default:
			fmt.Println(">>> Unknown error.")
		}
	}
}
