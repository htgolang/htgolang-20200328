package main

import (
	"log"
	"net/http"
)

func main() {
	addr := ":9999"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("www"))))
	log.Fatal(http.ListenAndServe(addr, nil))
}
