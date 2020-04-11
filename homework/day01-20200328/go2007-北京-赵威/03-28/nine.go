//打印9*9乘法表
/*
1 * 1 =  1
1 * 2 =  2  2 * 2 = 4
1 * 3 =  3  2 * 3 = 6  3 * 3 = 9
*/
package main
import (
	"fmt"
)
func main() {
for i := 1;i<=9;i++ {
	for j := 1; j<=i;j++ {
		fmt.Printf("%d *  %d = %2d  ",i,j,i * j)
	}
	fmt.Println()
}

}