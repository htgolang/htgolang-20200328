package main

import (
	"fmt"
)

func main() {
	multi := [][]string{}

	fmt.Printf("%T, %#v\n", multi, multi)

	multi = append(multi, []string{"1", "2", "3"})
	multi = append(multi, []string{"1", "2", "3", "5"})
	fmt.Println(multi)

	fmt.Printf("%T, %#v\n", multi[0], multi[0])
	fmt.Printf("%T, %#v\n", multi[0][1], multi[0][1])

	multi[0][1] = "xyz"
	multi[1] = append(multi[1], "xxxx")
	fmt.Println(multi)

}
