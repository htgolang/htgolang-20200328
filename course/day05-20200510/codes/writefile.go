package main

import "os"

func main() {

	file, err := os.Create("name.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.Write([]byte("abc123456"))
	file.Write([]byte("abc123456"))
	file.Write([]byte("abc123456"))
	file.Write([]byte("abc123456"))
	file.Write([]byte("abc123456"))

}
