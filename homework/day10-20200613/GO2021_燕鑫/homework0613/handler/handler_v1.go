package handler

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"homework0613/todolist"
	"homework0613/tools"
	"net/http"
	"html/template"
)

func ListTasksById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	dburl := tools.DBFILE
	content := make(map[string]interface{})
	user := p.ByName("user")
	id := p.ByName("id")
	tsrv := todolist.NewTaskService(dburl, user)
	filter := ""
	if user == "root" {
		if id != "" {
			filter = fmt.Sprintf("id=%s", id)
		}
	} else {
		if id != "" {
			filter = fmt.Sprintf("id=%s user=%s", id, user)
		} else {
			filter = fmt.Sprintf("user=%s", user)
		}
	}
	result, resultcnt, err, sortkey, desc := tsrv.GetByFilter(filter)
	if err != nil {
		content["err"] = true
		content["errstr"] = err
	} else {
		content["err"] = false
		content["resultcnt"] = resultcnt
		theader, tbody := tsrv.PrintLines(result, sortkey, desc)
		content["theader"] = theader
		content["tbody"] = tbody
	}

	t := template.New("list.html")
	t, errt := t.ParseFiles("www/list.html")
	if errt != nil {
		fmt.Println(errt)
	}
	erre := t.Execute(w, content)
	if erre != nil {
		fmt.Println(erre)
	}
}
