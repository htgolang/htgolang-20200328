package main

import (
	"fmt"
	"math/rand"
	"time"
)
var flag bool = false
func guessNumber() {
	for {
		rand.Seed(time.Now().Unix()) //生成一个随机数种子
		var answer int = rand.Intn(100) //生成一个随机数
		var maxGuessNumber int = 5

		for i := 0; i < maxGuessNumber; i++ {
			var guess int
			fmt.Println("请输入你猜的数字：")
			fmt.Scan(&guess)
			if guess == answer {
				fmt.Printf("恭喜你只用了 %d 次就猜对了\n", i+1)
				flag = true
				break //猜对了的话则跳出内层循环
			} else if guess > answer {
				fmt.Printf("猜大了,还有 %d 次机会\n", maxGuessNumber-i-1)
			} else {
				fmt.Printf("猜小了,还有 %d 次机会\n", maxGuessNumber-i-1)
			}
			if i == maxGuessNumber -1 {
				fmt.Println("5次都没猜对，你太笨了，退出游戏！")
				fmt.Printf("告诉你答案：%d\t", answer)
				flag = true
				break //5次都没猜对了的话则跳出内层循环
			}
		}
		if flag == true {
			var choiceAgent string
			fmt.Println("小帅哥，要不要再来一次？yes or y or Y")
			fmt.Scan(&choiceAgent)
			if choiceAgent == "yes" || choiceAgent == "y" || choiceAgent == "Y" {
				fmt.Println("再来一次，加油！")
			}else {
				fmt.Println("别走啊。。。")
				break //选择了不玩了，则退出死循环
			}
		}
	}
}

func main() {
	guessNumber()
}

/*
猜数字游戏 生成随机整数[0, 100) 提示用户再控制台输入猜测的数字 比较，当用户输入较大，提示太大了 当用户输入太小，提示太小了 当用户输入正确，
提示经过N次对了，太聪明了 用户最多猜5次，如果5次内都没有猜正确，提示太笨了，游戏结束。
扩展: 当成功或失败后，提示用户是否继续，输入：yes, y, Y则继续，重新生成随机数，让用户猜测
*/
