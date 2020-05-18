package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	h bool
	f string
)

func init()  {
	flag.BoolVar(&h,"h",false,"Help information")
	flag.StringVar(&f,"f","","Track file context")
}

func main() {
	flag.Parse()
	if flag.NFlag()==0{
		h=true
	}
	if f!=""{
		_,err:=os.Stat(f)
		if os.IsNotExist(err){
			fmt.Println("File doesn't exist!")
			os.Exit(36)
		}
		f1,_:=os.Open(f)
		defer f1.Close()
		bufr:=bufio.NewReader(f1)
		bufw:=bufio.NewWriter(os.Stdout)
		for{
			ctx:=make([]byte,1024)
			n,err:=bufr.Read(ctx)
			if err==io.EOF{
				continue
			}
			_,_=bufw.Write(ctx[:n])
			_=bufw.Flush()
		}
	}
	if h{
		usage()
	}
}

func usage()  {
	fmt.Fprintln(os.Stdout,"Usage of tail:")
	flag.PrintDefaults()
}
