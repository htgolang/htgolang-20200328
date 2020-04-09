package main

import "fmt"

func main() {
	for aVal := 1; aVal <= 9; aVal++ {
		for bVal := 1; bVal <= 9; bVal++ {
			if aVal < bVal {
				continue
			}
			// fmt.Printf("%d%v%d%v%02d\t",bVal,"x",aVal,"=",(bVal*aVal))
			fmt.Printf("%dx%d=%02d\t", bVal, aVal, (bVal * aVal))
			// fmt.Print(bVal,"x",aVal,"=")
			// fmt.Printf("%02d  ",aVal*bVal)
		}
		fmt.Println("")
	}
	for aVal := 1; aVal <= 9; aVal++ {
		for bVal := 1; bVal <= 9; bVal++ {
			if aVal > bVal {
				continue
			}
			fmt.Print(aVal, "x", bVal, "=")
			fmt.Printf("%02d  ", aVal*bVal)
		}
		fmt.Println("")
	}
}
