package main

import (
	"copy/fileutils"
	"flag"
	"fmt"
	"os"
)

var (
	recursive bool
	h         bool
	help      bool
)

func main() {
	flag.BoolVar(&recursive, "r", false, "copy recursive")
	flag.BoolVar(&h, "h", false, "help")
	flag.BoolVar(&help, "help", false, "help")
	flag.Parse()
	flag.Usage = func() {
		fmt.Println("Usage: copy [-r] srcDir destDir")
		flag.PrintDefaults()
	}

	if h || help {
		flag.Usage()
		os.Exit(1)
	}

	if flag.NArg() != 2 {
		fmt.Println("参数错误!")
		os.Exit(1)
	}

	src, dest := flag.Args()[0], flag.Args()[1]

	if recursive && fileutils.IsDir(src) && fileutils.IsDir(dest) {
		fmt.Println(src, dest)
		fileutils.TraverseCopy(src, dest)
	}
}
