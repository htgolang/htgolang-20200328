package main

import "fmt"

func main() {

	letters := "abcdefghi"
	//0. 1 2 3 4 ... len(letters) - 1
	for i := 0; i < len(letters); i++ {
		fmt.Printf("%c\n", letters[i])
	}

	letters = "我爱中华人民共和国"
	for _, v := range letters {
		fmt.Printf("%q\n", v)
	}
}
