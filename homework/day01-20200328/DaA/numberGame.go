package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	/*
		猜数字游戏 生成随机整数[0, 100) 提示用户再控制台输入猜测的数字 比较，当用户输入较大，提示太大了
		当用户输入太小，提示太小了 当用户输入正确，提示经过N次对了，太聪明了 用户最多猜5次，如果5次内都没有猜正确，
		提示太笨了，游戏结束

		扩展: 当成功或失败后，提示用户是否继续，输入：yes, y, Y则继续，重新生成随机数，让用户猜测

		1.进入游戏
		2.输入字符，转换类型并存储
		3.判断准确度，提示大小
		4.yes重新开始游戏，其他退出

		代码逻辑：
		for死循环，通过isQuit变量的值，控制是否开始游戏
		进入游戏后，通过for死循环进行单词游戏，通过usedGameNum计数控制最大尝试次数，
		通过userStr获取用户输入并判断是否符合数字要求，符合计算游戏次数，不符合要求用户重新输入。
		符合数字规则时，进行游戏逻辑判断。
	*/

	rand.Seed(time.Now().Unix()) //给rand函数设置一个随机种子
	userStr := ""                //初始化一个错误的值，让逻辑进入重新输入环节。也用来防止用户输入非数字
	isQuit := false              //是否退出游戏
	maxGameNum := 6              //每轮最多猜几次

	for {
		if isQuit { //退出游戏
			fmt.Println("Goodbye!")
			break
		} else { //进入游戏
			usedGameNum := 1          //初始化游戏猜测次数计数
			gameNum := rand.Intn(100) //游戏开始，生成一个随机数
			fmt.Println(gameNum)      //把要猜的数字直接打出来，方便调试
			fmt.Println("[0]:欢迎来带猜数字游戏！")

			for {
				if usedGameNum < maxGameNum { //游戏次数小于最大次数时，才可以继续游戏
					if inNum, err := strconv.Atoi(userStr); err == nil { //用户必须输入数字才可以计算游戏次数
						if inNum == gameNum {
							fmt.Printf("[1]:%d次就猜对了, 你太聪明了!\n", usedGameNum)
							//猜对了之后，询问用户是否继续游戏
							fmt.Print("[2]:再玩一局？yes/no: ")
							fmt.Scan(&userStr)
							if userStr != "y" && userStr != "Y" && userStr != "yes" { //当用户选择退出时，修改退出变量为True
								isQuit = true
							}
							break
						} else if inNum > gameNum {
							fmt.Println("数字大了")

						} else if inNum < gameNum {
							fmt.Println("数字小了")
						}
						userStr = ""  //清空用户的输入信息，让下一次循环顺利进入输入环节
						usedGameNum++ //输入的是数字，才计算有效次数

					} else { //用户输入的不是一个数字，需要重新输入
						fmt.Print("请输入一个整数: ")
						fmt.Scan(&userStr)
					}
				} else {
					fmt.Printf("[1]:%d次都错了，你太笨了！你没有资格重新开始了，游戏结束！\n", usedGameNum)
					isQuit = true
					break
				}
			} //单次游戏逻辑

		}
	}

}
