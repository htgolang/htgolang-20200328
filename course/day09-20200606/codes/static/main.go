package main

import "net/http"

func main() {
	http.Handle("/static/", http.FileServer(http.Dir("./")))
	http.Handle("/static2/", http.StripPrefix("/static2/", http.FileServer(http.Dir("./www"))))

	http.ListenAndServe(":8888", nil)
}
