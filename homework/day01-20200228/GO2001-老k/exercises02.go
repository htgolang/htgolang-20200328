package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//实现一个猜数字的游戏
func main() {
	var (
		inputStr string //获取用户输入
		intput  int //用户输入的值
		result int //答案
		err error
	)
	for {
		num := 5
		//获取100以内的随机数为答案
		rand.Seed(time.Now().Unix())
		result = rand.Intn(100)
		for i := 1; i <= num; i++ {
			rnum := num - i
			if _, err = fmt.Scan(&inputStr); err != nil {
				fmt.Println("输入非法")
				continue
			}
			if intput, err = strconv.Atoi(inputStr); err != nil {
				fmt.Println("输入非法")
				continue
			}

			switch {
			case intput > result:
				fmt.Printf("数字大了,你还可以猜%d次\n", rnum)
			case intput < result:
				fmt.Printf("数字小了,你还可以猜%d次\n", rnum)
			case intput == result:
				fmt.Println("猜对了")
				break
			}
		}
		fmt.Println("继续游戏?[Y,y,yes]")
		isQuit := ""
		fmt.Scan(&isQuit)
		if isQuit != "y" && isQuit != "Y" && isQuit != "yes" {
			break
		}
	}
}

