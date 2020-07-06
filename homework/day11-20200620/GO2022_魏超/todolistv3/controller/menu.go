package controller

import (
	"html/template"
	"net/http"
)

func Menu(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.New("menu.html").ParseFiles("views/menu/menu.html"))
	tmpl.Execute(response, struct{}{})
}
