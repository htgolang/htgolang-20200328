package main

import (
	"os"
)

func main() {
	file, err := os.OpenFile("password.txt", os.O_TRUNC|os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		return
	}
	defer file.Close()

	file.Write([]byte("abc"))
	file.Sync()
	file.Write([]byte("123"))
	file.WriteString("abcdefg")
}
