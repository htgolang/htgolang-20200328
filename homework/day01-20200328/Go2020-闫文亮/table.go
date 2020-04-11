package main

import "fmt"

func main(){
	for i:=1;i<=9;i++{
		for j:=1;j<i+1;j++{
			fmt.Printf("%-2d* %-2d= %-3d ", i,j,i*j)
		}
		fmt.Println()
	}
}
