package ioutils

import "fmt"

func Error(txt string) {
	fmt.Printf("[-] %s\n", txt)
}

func Success(txt string) {
	fmt.Printf("[+] %s\n", txt)
}

func Output(txt string) {
	fmt.Println(txt)
}
