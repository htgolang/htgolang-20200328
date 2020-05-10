package main

import (
	"os"
	"time"
)

func main() {
	os.Remove("name2.txt")
	os.Chmod("name2.txt", 0700)
	os.Chown("name2.txt", 1, 1)
	os.Chtimes("name2.txt", time.Now(), time.Now())
	// os.Link()
}
