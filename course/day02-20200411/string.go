package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	ascii := "abc我爱中华人民共和国"
	fmt.Println([]byte(ascii))
	fmt.Println([]rune(ascii))

	fmt.Println(len(ascii))
	fmt.Println(len([]rune(ascii)))

	fmt.Println(utf8.RuneCountInString(ascii))

	fmt.Println(string([]byte(ascii)))
	fmt.Println(string([]rune(ascii)))

	// int, float, bool
	fmt.Println(strconv.Itoa('a'))
	ch, err := strconv.Atoi("97")
	fmt.Println(ch, err)

	fmt.Println(strconv.FormatFloat(3.1415926, 'f', 10, 64))
	pi, err := strconv.ParseFloat("3.1415924", 64)
	fmt.Println(pi, err)

	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.ParseBool("true"))

	b, err := strconv.ParseInt("5", 10, 8)
	fmt.Println(b, err)

	fmt.Println(strconv.FormatInt(15, 2))
}
