package handler

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/julienschmidt/httprouter"
	"homework0620/todolist"
	"html/template"
	"net/http"
)

type ServiceHandler struct {
	db       *gorm.DB
	taskSrv  *todolist.TaskService
	usersrv  *todolist.UserService
	username string
}

func NewServiceHandler(db *gorm.DB) *ServiceHandler {
	return &ServiceHandler{db: db, taskSrv: todolist.NewTaskService(db), usersrv: todolist.NewUserService(db)}
}

func (s *ServiceHandler) ListTasksById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	content := make(map[string]interface{})
	username := p.ByName("user")
	if !s.logincheck(w, r, username) {
		fmt.Println("must login")
		return
	}
	id := p.ByName("id")
	filter := ""
	if username == "root" {
		if id != "" {
			filter = fmt.Sprintf("id=%s", id)
		}
	} else {
		if id != "" {
			filter = fmt.Sprintf("id=%s user=%s", id, username)
		} else {
			filter = fmt.Sprintf("user=%s", username)
		}
	}
	result, resultcnt, err := s.taskSrv.GetByFilter(filter, "")
	if err != nil {
		content["err"] = true
		content["errstr"] = err
	} else {
		content["err"] = false
		content["resultcnt"] = resultcnt
		content["theader"] = []string{"id", "name", "starttime", "endtime", "status", "user"}
		content["tbody"] = result
		content["currentuser"] = s.username
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

func (s *ServiceHandler) CreateTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("user")
	if !s.logincheck(w, r, username) {
		return
	}
	if username != s.username && s.username != "root" {
		fmt.Println("You have no rights!")
		return
	}
	t := template.New("create.html")
	t, errt := t.ParseFiles("www/create.html")
	if errt != nil {
		fmt.Println(errt)
	}
	if r.Method == http.MethodPost {
		taskname := r.PostFormValue("Taskname")
		err := s.taskSrv.CreateNewTask(taskname)
		if err != nil {
			fmt.Println("\n", err, "\n")
		}
		http.Redirect(w, r, fmt.Sprintf("/listtask/%s/", username), 302)
	} else {
		erre := t.Execute(w, map[string]string{"User": username})
		if erre != nil {
			fmt.Println(erre)
		}
	}
}

func (s *ServiceHandler) UpdateTaskById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("user")
	if !s.logincheck(w, r, username) {
		return
	}
	if username != s.username && s.username != "root" {
		fmt.Println("You have no rights!")
		return
	}
	taskid := p.ByName("id")
	resultItem, _, err := s.taskSrv.GetByFilter(fmt.Sprintf("id=%s", taskid), "")
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}

	t := template.New("update.html")
	t, errt := t.ParseFiles("www/update.html")
	if errt != nil {
		fmt.Println(errt)
	}
	if r.Method == http.MethodPost {
		status := r.PostFormValue("Status")
		user := r.PostFormValue("User")
		_, err := s.taskSrv.UpdateTask(&resultItem[0], status, user)
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

func (s *ServiceHandler) DeleteTaskById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("user")
	if !s.logincheck(w, r, username) {
		return
	}
	if username != s.username && s.username != "root" {
		fmt.Println("You have no rights!")
		return
	}
	taskid := p.ByName("id")
	resultItem, _, err := s.taskSrv.GetByFilter(fmt.Sprintf("id=%s", taskid), "")
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}

	_, err = s.taskSrv.DeleteTask(&resultItem[0])
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/listtask/%s/", username), 302)
}

func (s *ServiceHandler) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t := template.New("login.html")
	t, errt := t.ParseFiles("www/login.html")
	if errt != nil {
		fmt.Println(errt)
	}
	if r.Method == http.MethodPost {
		username := r.PostFormValue("Username")
		password := r.PostFormValue("Password")
		_, err := s.usersrv.VerifyUser(username, password)
		if err != nil {
			http.Redirect(w, r, "/login/", 302)
			fmt.Println(err)
		} else {
			s.username = username
			s.taskSrv.SetUser(username)
			http.Redirect(w, r, fmt.Sprintf("/listtask/%s/", username), 302)
		}
	} else {
		erre := t.Execute(w, "")
		if erre != nil {
			fmt.Println(erre)
		}
	}
}

func (s *ServiceHandler) Logout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	s.username = ""
	http.Redirect(w, r, "/login/", 302)
}

func (s *ServiceHandler) SignUp(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Method == http.MethodPost {
		username := r.PostFormValue("Username")
		password := r.PostFormValue("Password")
		err := s.usersrv.CreateUser(username, password)
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "/login/", 302)
	} else {
		http.Redirect(w, r, "/login/", 302)
	}
}

func (s *ServiceHandler) logincheck(w http.ResponseWriter, r *http.Request, username string) bool {
	if s.username == "" {
		fmt.Println("must login")
		http.Redirect(w, r, "/login/", 302)
		return false
	}
	if username != s.username && s.username != "root" {
		fmt.Println("You have no rights!")
		return false
	}
	return true
}
