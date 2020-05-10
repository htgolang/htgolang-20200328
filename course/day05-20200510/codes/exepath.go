package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Println(os.Getwd())
	fmt.Println(os.Executable())
}