package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Printf("%q\n", strings.Join([]string{}, ":"))
	fmt.Printf("%q\n", strings.Join([]string{"kk"}, ":"))
}
