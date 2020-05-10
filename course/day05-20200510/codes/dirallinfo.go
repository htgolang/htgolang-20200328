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
	fileInfos, err := file.Readdir(-1)
	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name(), fileInfo.IsDir())
	}
}
