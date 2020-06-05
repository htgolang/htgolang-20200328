package main

import (
	"copy/utils"
	"flag"
	"fmt"
	"os"
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
	srcFileInfo, err := os.Stat(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	if srcFileInfo.IsDir() {
		err = utils.CopyDir(src, dst)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		err = utils.CopyFile(src, dst)
		if err != nil {
			fmt.Println(err)
		}
	}
}
