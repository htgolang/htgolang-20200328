package main

import (
	"io/ioutil"
	"os"
)

func main() {
	ioutil.WriteFile("test/user.txt", []byte("我是KK"), os.ModePerm)
}
