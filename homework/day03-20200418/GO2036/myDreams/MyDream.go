package main

import (
	"fmt"
	"sort"
)

func main() {
	var charMap map[string]int = map[string]int{}
	char := `But one hundred years later, the Negro still is not free. One hundred years later, the life of the Negro is still sadly crippled by the manacles of segregation and the chains of discrimination. One hundred years later, the Negro lives on a lonely island of poverty in the midst of a vast ocean of material prosperity. One hundred years later, the Negro is still languished in the corners of American society and finds himself an exile in his own land. And so we've come here today to dramatize a shameful condition.
In a sense we've come to our nation's capital to cash a check. When the architects of our republic wrote the magnificent words of the Constitution and the Declaration of Independence, they were signing a promissory note to which every American was to fall heir. This note was a promise that all men, yes, black men as well as white men, would be guaranteed the "unalienable Rights" of "Life, Liberty and the pursuit of Happiness." It is obvious today that America has defaulted on this promissory note, insofar as her citizens of color are concerned. Instead of honoring this sacred obligation, America has given the Negro people a bad check, a check which has come back marked "insufficient funds."
But we refuse to believe that the bank of justice is bankrupt. We refuse to believe that there are insufficient funds in the great vaults of opportunity of this nation. And so, we've come to cash this check, a check that will give us upon demand the riches of freedom and the security of justice.`
	var charSort map[int][]string = map[int][]string{}
	for _,s := range char {
		if s >= 'a' && s <= 'z' || s >= 'A' && s <= 'Z' {
			charMap[string(s)] = charMap[string(s)] + 1
		}

	}
	for k,v := range charMap {
		if _,ok := charSort[v]; ok {
			charSort[v] = append(charSort[v],k)
		} else {
			charSort[v] = []string{k}
		}
	}

	fmt.Println(charSort)

	SortCh := []string{}
	SortInt := []int{}
	for k,_ := range charSort {
		SortInt = append(SortInt,k)
	}
	fmt.Println(SortInt)
	sort.Sort(sort.Reverse(sort.IntSlice(SortInt)))
	fmt.Println(SortInt)

	if len(SortInt) >= 10 {
		for _,v := range SortInt[:10] {
			SortCh = append(SortCh,charSort[v]...)
		}
		} else {
		for _,v := range SortInt {
			SortCh = append(SortCh,charSort[v]...)
		}
	}

	for _,v := range SortCh {
		fmt.Printf("%s,%d\n",v,charMap[v])
	}

}