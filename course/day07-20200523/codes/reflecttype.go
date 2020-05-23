package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	var es = []interface{}{1, "test", false, time.Now()}
	for _, e := range es {
		typ := reflect.TypeOf(e)
		fmt.Println(typ.Name(), typ.PkgPath(), int(typ.Kind()), typ.NumMethod())
	}
}
