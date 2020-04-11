package main
import "fmt"
func main(){
	for x :=1;x<10;x++ {
		for y:=1;y<=x;y++ {
			fmt.Printf("%d * %d = %2d  ",x,y,x*y)
		}
		fmt.Println()
	}
	
}