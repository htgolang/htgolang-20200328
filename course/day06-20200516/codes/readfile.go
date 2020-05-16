package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	ctx, err := ioutil.ReadFile("multiwriter.go")
	fmt.Println(string(ctx), err)
}
