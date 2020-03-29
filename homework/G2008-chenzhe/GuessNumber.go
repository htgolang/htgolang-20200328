package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for{
		rand.Seed(time.Now().Unix())
		var TmpRandNum=rand.Intn(100)
		var Right bool = false
		var InNum int
		for i:=0;i<5;i++{
			
			fmt.Println("请输入你要猜测的数字")
			fmt.Scan(&InNum)
			if InNum < TmpRandNum{
				fmt.Println("你输入的数字太小了")
			}else if InNum > TmpRandNum{
				fmt.Println("你输入的数字太大了")
			}else {
				fmt.Printf("恭喜你，经过%d次，终于答对了,太聪明了",i+1)
				Right = true
				break
			}
		}
		if Right == false{
			fmt.Println("太笨了，游戏结束")
		}
		var RunOrStop string
		for {
			var IfBreak bool = false
			fmt.Println("你是否要继续游戏？（yes,y,Y/no,n,N）")
			fmt.Scan(&RunOrStop)
			switch RunOrStop {
			case "yes","y","Y":
				fmt.Println("马上开始下一轮游戏")
				IfBreak=true
			case "no","n","N":
				fmt.Println("即将推出游戏")
				return
			default:
				fmt.Println("错误指令,请重新输入")
			}
			if IfBreak{
				break
			}
		}

	}
}