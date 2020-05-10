package main

import (
	"fmt"
	"os"
)

func main() {

	fileInfo, err := os.Stat("password")
	fmt.Println(err)
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())

	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.IsDir())

	fmt.Println(fileInfo.Mode())
}
