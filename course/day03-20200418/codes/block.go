package main

import "fmt"

func main() {
	name, desc := "kk", "i'm kk"

	func(name string) {
		fmt.Println(name, desc) // malukang, i'm kk
		name, desc = "烽火", "燕鑫"
		fmt.Println(name, desc) // 烽火, 燕鑫
	}("malukang")
	fmt.Println(name, desc) // kk, 燕鑫

}
