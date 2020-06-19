package handler

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"homework0613/todolist"
	"html/template"
	"net/http"
)

type ServiceHandler struct {
	taskSrv *todolist.TaskService
}

func NewServiceHandler(dburl string) *ServiceHandler {
	return &ServiceHandler{taskSrv: todolist.NewTaskService(dburl)}
}

func (s *ServiceHandler) ListTasksById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	content := make(map[string]interface{})
	user := p.ByName("user")
	id := p.ByName("id")
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
	result, resultcnt, err, sortkey, desc := s.taskSrv.GetByFilter(filter)
	if err != nil {
		content["err"] = true
		content["errstr"] = err
	} else {
		content["err"] = false
		content["resultcnt"] = resultcnt
		theader, tbody := s.taskSrv.PrintLines(result, sortkey, desc)
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

func (s *ServiceHandler) UpdateTaskById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//content := make(map[string]interface{})
	username := p.ByName("user")
	s.taskSrv.SetUser(username)
	taskid := p.ByName("id")
	resultItem, _, err, _, _ := s.taskSrv.GetByFilter(fmt.Sprintf("id=%s", taskid))
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}

	t := template.New("update.html")
	t, errt := t.ParseFiles("www/update.html")
	if errt != nil {
		fmt.Println(errt)
	}
	fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@")
		status := r.PostFormValue("Status")
		user := r.PostFormValue("User")
		_, err := s.taskSrv.UpdateTask(resultItem[0], status, user)
		if err != nil {
			fmt.Println("\n", err, "\n")
		}
		http.Redirect(w, r, fmt.Sprintf("/listtask/%s/", username), 302)
	} else {
		erre := t.Execute(w, resultItem[0])
		if erre != nil {
			fmt.Println(erre)
		}
	}
}
