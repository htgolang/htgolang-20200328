package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println(json.Valid([]byte("1")))
	fmt.Println(json.Valid([]byte("true")))
	fmt.Println(json.Valid([]byte("[true]")))
	fmt.Println(json.Valid([]byte("[{'name' : 'kk'}]")))
	fmt.Println(json.Valid([]byte("[{\"name\" : \"kk\"}]")))
	fmt.Println(json.Valid([]byte(`[{"name" : "kk"}]`)))
}
