package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var testFile = "tasks.txt"

func main() {

	file, _ := os.Open(testFile)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("%T, %#v\n", scanner.Text(), scanner.Text())
		fmt.Println(strings.Repeat("--", 20))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
