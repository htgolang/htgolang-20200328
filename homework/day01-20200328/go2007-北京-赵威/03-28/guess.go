//猜数字游戏
package main
import (
	"fmt"
	"time"
	"math/rand"
)
func main() {
	var i int = 0
	rand.Seed(time.Now().Unix())
	for {
		i++
		anum := rand.Intn(100)
		fmt.Println("请输入您猜的数字： ")
		var bnum int
		fmt.Scan(&bnum)
		if i >=5		{
			fmt.Println("您的游戏次数已经用完，是否继续（Y/N）")
			var newsel string
			fmt.Scan(&newsel)
			if newsel == "Y" || newsel == "y" {
				 i = 1
			}else {
				fmt.Println("游戏结束")
				break
			}



		}

		switch {
		case anum == bnum :
			fmt.Printf("这个是第%d次，恭喜你猜对啦",i)
			fmt.Println("是否继续游戏: (Y/N)")
			var selc string
			fmt.Scan(&selc)
			if selc == "Y" || selc == "y" {
				continue
			}
		case anum < bnum :
			fmt.Printf("这个是第%d次，输入数字太大啦，您还有%d次机会\n",i,5-i)
		case anum > bnum :
			fmt.Printf("这个是第%d次，输入数字太小啦，您还有%d次机会\n",i,5-i)
		
		}

	
}
		
}
