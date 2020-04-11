//猜数字
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	var(
		randNumber int
		userNumber int
		count int = 1
		game bool = true
		again string
	)

	rand.Seed(time.Now().Unix())
	randNumber = rand.Intn(100)

	for game == true {

		fmt.Print("请输入您猜测的数字,本轮共五次机会:")
		fmt.Scan(&userNumber)

        if count < 5{
			if userNumber > randNumber{
				fmt.Printf("您猜测的数字太大了,您还有%d次机会\n",(5-count))
				count ++
			}else if userNumber < randNumber{
				fmt.Printf("您猜测的数字太小了,您还有%d次机会\n",(5-count))
				count ++
			}else if userNumber == randNumber{
				fmt.Printf("恭喜您经过%d次猜对了", count)
				break
			}
	   }else{	
			fmt.Printf("超过5次，游戏结束。还玩吗？[y/n]")			
			fmt.Scan(&again)
			switch again{
			case "yes", "y", "Y":
				randNumber = rand.Intn(100)
				count = 0
				game = true
			case "n","NO","no":
				fmt.Print("游戏结束。")
				game = false
			default:
				fmt.Print("游戏结束。")
                game = false

			}
		}
	
	}

}