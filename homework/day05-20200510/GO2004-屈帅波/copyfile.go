package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func copyfile(src,desc string) {
	srcfile,err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	} else {
		defer srcfile.Close()
		descfile,err := os.Create(desc)
		if err != nil {
			fmt.Println(err)
		}else {
			defer descfile.Close()
			bytes,err := ioutil.ReadAll(srcfile)
			if err != nil {
				fmt.Println(err)
			}else {
				descfile.Write(bytes)
			}
		}

	}
}


func main() {
	src  := flag.String("s","","src file")
	desc := flag.String("d","","desc file")
	help := flag.Bool("h",false,"help")
	flag.Usage = func() {
		fmt.Println("Usage: copyfile -s xxx  -d xxx")
		flag.PrintDefaults()
	}
	//解析参数
	flag.Parse()
	if *src == "" || *desc == "" || *help  {
		flag.Usage()
	} else {
		copyfile(*src,*desc)
	}
}