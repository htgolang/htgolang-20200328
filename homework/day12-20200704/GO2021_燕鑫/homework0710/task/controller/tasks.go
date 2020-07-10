package controller

import (
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
	"task/base/errors"
	"task/todolist"
)

type TasksController struct {
	beego.Controller
}


func (t *TasksController) ListTasks() {
	taskSrv:=todolist.NewTaskService()
	content := make(map[interface{}]interface{})
	username := t.Ctx.Input.Params()["0"]
	fmt.Println(username)
	//if !t.logincheck(username) {
	//	fmt.Println("must login")
	//	return
	//}
	id := t.Ctx.Input.Params()["1"]
	fmt.Println(id)
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
	fmt.Println(filter)
	result, resultcnt, err := taskSrv.GetByFilter(filter, "")
	if err != nil {
		content["err"] = true
		content["errstr"] = err
	} else {
		content["err"] = false
		content["resultcnt"] = resultcnt
		content["theader"] = []string{"id", "name", "starttime", "endtime", "status", "user"}
		content["tbody"] = result
		content["currentuser"] = username
	}

	t.Data = content

	t.TplName = "list.html"
}
//
//func (t *TasksController) CreateTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	username := p.ByName("user")
//	if !t.logincheck(w, r, username) {
//		return
//	}
//	if username != t.username && t.username != "root" {
//		fmt.Println("You have no rights!")
//		return
//	}
//	t := template.New("create.html")
//	t, errt := t.ParseFiles("www/create.html")
//	if errt != nil {
//		fmt.Println(errt)
//	}
//	if r.Method == http.MethodPost {
//		taskname := r.PostFormValue("Taskname")
//		err := t.taskSrv.CreateNewTask(taskname)
//		if err != nil {
//			fmt.Println("\n", err, "\n")
//		}
//		http.Redirect(w, r, fmt.Sprintf("/listtask/%s/", username), http.StatusFound)
//	} else {
//		erre := t.Execute(w, map[string]string{"User": username})
//		if erre != nil {
//			fmt.Println(erre)
//		}
//	}
//}
//
//func (t *TasksController) UpdateTaskById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	username := p.ByName("user")
//	if !t.logincheck(w, r, username) {
//		return
//	}
//	if username != t.username && t.username != "root" {
//		fmt.Println("You have no rights!")
//		return
//	}
//	taskid := p.ByName("id")
//	resultItem, _, err := t.taskSrv.GetByFilter(fmt.Sprintf("id=%s", taskid), "")
//	if err != nil {
//		fmt.Println("\n", err, "\n")
//		return
//	}
//
//	t := template.New("update.html")
//	t, errt := t.ParseFiles("www/update.html")
//	if errt != nil {
//		fmt.Println(errt)
//	}
//	if r.Method == http.MethodPost {
//		status := r.PostFormValue("Status")
//		user := r.PostFormValue("User")
//		_, err := t.taskSrv.UpdateTask(&resultItem[0], status, user)
//		if err != nil {
//			fmt.Println("\n", err, "\n")
//		}
//		http.Redirect(w, r, fmt.Sprintf("/listtask/%s/", username), http.StatusFound)
//	} else {
//		erre := t.Execute(w, resultItem[0])
//		if erre != nil {
//			fmt.Println(erre)
//		}
//	}
//}
//
//func (t *TasksController) DeleteTaskById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	username := p.ByName("user")
//	if !t.logincheck(w, r, username) {
//		return
//	}
//	if username != t.username && t.username != "root" {
//		fmt.Println("You have no rights!")
//		return
//	}
//	taskid := p.ByName("id")
//	resultItem, _, err := t.taskSrv.GetByFilter(fmt.Sprintf("id=%s", taskid), "")
//	if err != nil {
//		fmt.Println("\n", err, "\n")
//		return
//	}
//
//	_, err = t.taskSrv.DeleteTask(&resultItem[0])
//	if err != nil {
//		fmt.Println("\n", err, "\n")
//		return
//	}
//	http.Redirect(w, r, fmt.Sprintf("/listtask/%s/", username), http.StatusFound)
//}

func (t *TasksController) Login() {
	usersrv:=todolist.NewUserService()
	errs := errors.NewErrors()
	username := ""
	password := ""
	if t.Ctx.Input.IsPost() {
		username = t.GetString("Username")
		password = t.GetString("Password")
		_, err := usersrv.VerifyUser(username, password)
		if err != nil {
			t.Redirect("/tasks/login/", http.StatusFound)
			errs.Add("login", err.Error())
		} else {
			t.Redirect(fmt.Sprintf("/tasks/listtasks/%s/", username), http.StatusFound)
		}
	}
	fmt.Println(errs)
	t.Data["Username"] = username
	t.Data["Password"] = password
	t.Data["Errors"] = errs
	t.TplName = "login.html"
}

func (t *TasksController) Logout() {
	t.Redirect("/tasks/login/", http.StatusFound)
}

func (t *TasksController) SignUp() {
	usersrv:=todolist.NewUserService()
	if t.Ctx.Input.IsPost() {
		username := t.GetString("Username")
		password := t.GetString("Password")
		err := usersrv.CreateUser(username, password)
		if err != nil {
			fmt.Println(err)
		}
		t.Redirect("/tasks/login/", http.StatusFound)
	} else {
		t.Redirect("/tasks/login/", http.StatusFound)
	}
}

//func (t *TasksController) logincheck(username string) bool {
//	if t.username == "" {
//		fmt.Println("must login")
//		t.Redirect("/tasks/login/", http.StatusFound)
//		return false
//	}
//	if username != t.username && t.username != "root" {
//		fmt.Println("You have no rights!")
//		return false
//	}
//	return true
//}
