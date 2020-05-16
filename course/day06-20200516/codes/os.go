package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println(os.TempDir())
	fmt.Println(os.UserHomeDir())
	fmt.Println(os.UserCacheDir())
	path, _ := os.Executable() // 执行程序的路径
	fmt.Println(path)
	fmt.Println(filepath.Dir(path))
}
