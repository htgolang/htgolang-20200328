package  main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const max  = 3

func num() bool {
	fmt.Println("太笨了")
	fmt.Printf("请问是否重新来过(Y/N)")
	scnner := bufio.NewScanner(os.Stdin)
	scnner.Scan()
	//当需要重新开始返回true
	if scnner.Text() == "y" || scnner.Text() == "Y" {
		return true
	}
	//当不需要重新开始 false
	return false
}


func main() {
	rand.Seed(time.Now().Unix())
STARTONE:
	rand := rand.Intn(100)
	scnner := bufio.NewScanner(os.Stdin)
	for i:=1;i<=max;i++{
	STARTTWO:
		fmt.Printf("请输入您猜的数字:")
		scnner.Scan()
		usernum,err := strconv.Atoi(scnner.Text())
		if err != nil{
			fmt.Printf("输入有误请重新输入,还有%d次机会\n",max-i)
			//如果用户输入的不是int类型 那么 这里做判断

			//第一种情况 次数到达了最大限制  判断用户是否继续
			if i  == max {
				if num() {
					goto  STARTONE
				}
				break
				//第二种情况 次数没有到达最大限制 但是用户输入的不是int 那么这里有效次数-1  并且让用户重新输入
			}else {
				i++
				goto STARTTWO
			}
		}
		if usernum == rand {
			fmt.Printf("恭喜您用了%d次就猜对了\n",i)
			break
		}else if usernum < rand  {
			fmt.Printf("猜小了,还有%d次机会\n",max-i)
		}else if usernum > rand {
			fmt.Printf("猜大了,还有%d次机会\n",max-i)
		}
		//如果次数到达最大限制 那么开始判断用户是否继续
		if i == max {
			if num() {
				goto STARTONE
			}
			break
		}
	}
}
