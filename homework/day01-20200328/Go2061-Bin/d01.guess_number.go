package main

import (
	"fmt"
	"math/rand"
)

func main() {
START:
	var num = rand.Intn(100)
	//var num = rand.New(rand.NewSource(time.Now().UnixNano())).Int() % 100
	const maxGuessTime = 5
	fmt.Println("猜猜你的幸运数字是啥？(1~100) : ")
	var input int
	for i := maxGuessTime; i > 0; i-- {
		fmt.Print("Please type in your guessing: ")
		fmt.Scan(&input)
		switch {
		case input == num:
			fmt.Println("Bingo! 猜对了！")
			goto END
		case input > num:
			fmt.Print("大了……")
		default:
			fmt.Print("小了……")
		}
		fmt.Println("剩余", i-1, "次机会")
		if i == 1 {
			var restart string
			fmt.Println("Game Over...再玩一次(Y/N)？")
			fmt.Scan(&restart)
			if restart == "Y" || restart == "y" {
				fmt.Println("----------------------------")
				goto START
			} else {
				fmt.Println("再见！")
			}
		}
	}
END:
}
