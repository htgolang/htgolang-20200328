package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	var es = []interface{}{1, "test", false, time.Now()}
	for _, e := range es {
		val := reflect.ValueOf(e)
		fmt.Println()
	}
}
