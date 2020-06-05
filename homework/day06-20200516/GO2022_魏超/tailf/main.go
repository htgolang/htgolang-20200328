package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func tailf(files ...string) {
	var filehandle = make([]io.Reader, 0)
	for _, file := range files {
		fileinfo, err := os.Stat(file)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("%s No such file.", file)
			} else if os.IsPermission(err) {
				fmt.Printf("%s Permission denied.", file)
			} else {
				fmt.Printf("%s %s", file, err)
			}
			return
		}
		if fileinfo.IsDir() {
			fmt.Printf("%s is a directory", file)
		}
	}
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("%s %s", file, err)
			return
		}
		defer f.Close()
		f.Seek(0, os.SEEK_END)
		filehandle = append(filehandle, f)
	}
	for {
		reader := io.MultiReader(filehandle...)
		_, err := io.Copy(os.Stdout, reader)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	var h, help bool
	flag.BoolVar(&h, "h", false, "help")
	flag.BoolVar(&help, "help", false, "help")
	flag.Usage = func() {
		fmt.Println(`
Usage: tailf file...
	`)
		flag.PrintDefaults()
	}
	flag.Parse()
	if h || help {
		flag.Usage()
	}
	if flag.NArg() < 1 {
		fmt.Println("missing file operand.")
	}
	tailf(flag.Args()...)
}
