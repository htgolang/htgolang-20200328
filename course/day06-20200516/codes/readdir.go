package main

import (
	"fmt"
	"io/ioutil"
)

// B
// FormatSize(int64) // 字节大小B => > KB > MB > GB > TB > PB

func FormatSize(size int64) string {
	// size int64 => float64
	// B
	// size >= 1024 => size / 1024 KB
	// size / 1024 >= 1024 =>  size / 1024 / 1024 MB
	// size / 1024 / 1024 >= 1024 =>  size / 1024 / 1024 / 1023 GB

	// units := map[int]string{0: "B", 1: "KB", 2: "MB", 3: "GB", 4: "TB", 5: "PB"}
	units := [6]string{"B", "KB", "MB", "GB", "TB", "PB"}

	fsize := float64(size)
	unit := float64(1024)

	index := 0

	for fsize >= unit && index < len(units)-1 {
		fsize /= unit
		index++
	}
	return fmt.Sprintf("%.2f%s", fsize, units[index])
}

func main() {
	files, err := ioutil.ReadDir(".")
	fmt.Println(err)
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("D", file.Name(), file.ModTime().Format("2006-01-02 15:03:04"))
		} else {
			fmt.Println("F", file.Name(), file.Size(), file.ModTime().Format("2006-01-02 15:03:04"))
		}
	}

	fmt.Println(FormatSize(10))
	fmt.Println(FormatSize(1024))
	fmt.Println(FormatSize(1030))
	fmt.Println(FormatSize(1024 * 1024))
	fmt.Println(FormatSize(1024 * 1025))
	fmt.Println(FormatSize(1024 * 1025 * 1024))
	fmt.Println(FormatSize(1024 * 1025 * 1024 * 1024))
	fmt.Println(FormatSize(1024 * 1025 * 1024 * 1024 * 1024))
	fmt.Println(FormatSize(1024 * 1025 * 1024 * 1024 * 1024 * 1024))

}
