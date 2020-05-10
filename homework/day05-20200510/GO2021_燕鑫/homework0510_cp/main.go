package main

import (
	"fmt"
	"homework0510_cp/handlefile"
	"os"
)

func main() {
	fcp:=handlefile.NewFileCp(os.Args[1],os.Args[2])
	err:=fcp.CP()
	if err != nil {
		fmt.Println(err)
		os.Exit(12)
	}
}
