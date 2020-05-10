package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	var name string

	// -p
	// -s
	flag.StringVar(&name, "p", "", "path")
	flag.Parse()

	if name == "" {
		return
	}
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	hasher := md5.New()
	ctx := make([]byte, 1024)
	for {
		n, err := file.Read(ctx)
		if err == io.EOF {
			break
		}
		hasher.Write(ctx[:n])
	}
	fmt.Printf("%x", hasher.Sum(nil))
}
