package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		randomInt   int
		inputInt    int
		inputString string
		i           = 1
		x           = 1
	)
	rand.Seed(time.Now().Unix())
	randomInt = rand.Intn(100)
	//fmt.Println(randomInt)
	fmt.Println("欢迎光临，这里是猜数字（0-100的整数）游戏。")
	//第一个五次机会
	for {
		if i <= 5 {
			fmt.Println("请输入你的答案：")
			fmt.Scan(&inputInt)
			if inputInt == randomInt {
				fmt.Printf("恭喜你！总共用了%d次就答对了！！！\n", i)
				break
			} else if inputInt > randomInt {
				i++
				fmt.Println("太大了！")
				fmt.Println(i)
				continue
			} else if inputInt < randomInt {
				i++
				fmt.Println("太小了")
				fmt.Println(i)
				continue
			}
		}
		fmt.Println("太笨了，五次机会已经用完！")
		break
	}
	//fmt.Println(rand.Intn(100))
	//第二次机会
	fmt.Println("请问是否需要再玩一次？是请输入Y，否请输入N。")
	fmt.Scan(&inputString)
	if inputString == "Y" {
		randomInt = rand.Intn(100)
		fmt.Println(randomInt)
		fmt.Println("欢迎再次光临猜数字（0-100的整数）游戏！")

		for {
			if x <= 5 {
				fmt.Println("请输入你的答案：")
				fmt.Scan(&inputInt)
				if inputInt == randomInt {
					fmt.Printf("恭喜你！总共用了%d次就答对了！！！\n", x)
					break
				} else if inputInt > randomInt {
					x++
					fmt.Println("太大了！")
					fmt.Println(x)
					continue
				} else if inputInt < randomInt {
					x++
					fmt.Println("太小了")
					fmt.Println(x)
					continue
				}
			}
			fmt.Println("太笨了，五次机会已经用完！")
			break
		}
	}
	if inputString == "N" {
		fmt.Println("不想玩了，太没意思了！")
	}

}
