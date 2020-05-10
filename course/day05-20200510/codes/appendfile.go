package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("name.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0700)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Write([]byte("123\t"))
	file.Write([]byte("abc\n"))
	file.Write([]byte("abc"))
	file.Write([]byte("123"))
}
