package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("password")

	if err != nil {
		return
	}
	defer file.Close()

	names, err := file.Readdirnames(-1)
	fmt.Println(names)
}
