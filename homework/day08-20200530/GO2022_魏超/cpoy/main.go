package main

import (
	"flag"
	"fmt"

	"copy/utils"
)

func main() {
	flag.Usage = func() {
		fmt.Println(`
Usage: copy sourcefile destfile
	`)
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() < 2 {
		fmt.Println("缺少参数.")
		return
	}
	src := flag.Arg(0)
	dst := flag.Arg(1)
	utils.Copy(src, dst)
}
