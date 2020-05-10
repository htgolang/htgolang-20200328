package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("test", os.ModePerm)
	fmt.Println(os.MkdirAll("test/a/b", os.ModePerm))
	fmt.Println(os.Remove("test/a/b"))
	os.RemoveAll("test")
}
