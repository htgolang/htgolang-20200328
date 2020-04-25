package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("%X\n", []byte("啦啦啦啦啦"))
	fmt.Println(hex.EncodeToString([]byte("啦啦啦啦啦")))
	txt, _ := hex.DecodeString("e595a6e595a6e595a6e595a6e595a6")
	fmt.Println(string(txt))

}