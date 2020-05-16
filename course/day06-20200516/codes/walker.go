package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, file os.FileInfo, err error) error {
		fmt.Println(path, file.Name(), file.IsDir())
		return nil
	})
}
