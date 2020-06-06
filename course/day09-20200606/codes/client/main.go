package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {

	// response, err := http.Get("http://localhost:8888/?a=1&b=2")

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(response.Proto, response.StatusCode, response.Status)
	// 	fmt.Println(response.Header)
	// 	io.Copy(os.Stdout, response.Body)

	// }

	// buffer := bytes.NewBufferString(`{"abbbb": 1111}`)
	// response, err := http.Post("http://localhost:8888", "application/json", buffer)
	// fmt.Println(response, err)

	params := url.Values{}
	params.Add("a", "1")
	params.Add("a", "2")
	params.Add("b", "3")

	response, err := http.PostForm("http://localhost:8888", params)
	fmt.Println(response, err)
}
