package main


import "fmt"


func main()  {
	for i:=1;i<10;i++{
		for j:=1;j<10;j++{
			if i>=j{
				fmt.Printf("%#v*%#v=%#v \t",j,i,i*j)
			}
		}
		fmt.Printf("\n")
	}
}