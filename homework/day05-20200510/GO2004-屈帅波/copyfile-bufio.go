package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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
			if err != nil {
				fmt.Println(err)
			}else {
				bytes := make([]byte,1024*1024)
				reader := bufio.NewReader(srcfile)
				writer := bufio.NewWriter(descfile)
				for {

					n, err := reader.Read(bytes)
					if err != nil {
						if err != io.EOF {
							fmt.Println(err)
						}
						break
					}
					writer.Write(bytes[:n])
					writer.Flush()
				}

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