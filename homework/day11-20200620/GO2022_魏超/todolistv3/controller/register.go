package controller

import (
	"html/template"
	"net/http"
	"strconv"
	"time"
	"todolist/models"
	"todolist/utils"
)

func Register(response http.ResponseWriter, request *http.Request) {
	var (
		err    error
		user   models.User
		users  []models.User
		errors = make(map[string]string)
		now    = time.Now()
	)

	if request.Method == http.MethodGet {

	} else if request.Method == http.MethodPost {
		users = models.GetUsers()
		user.Name = request.FormValue("name")
		user.Account = request.FormValue("account")
		user.Sex, _ = strconv.Atoi(request.FormValue("sex"))
		user.Tel = request.FormValue("tel")
		user.Address = request.FormValue("address")
		user.RegisterTime = &now
		passwd := request.FormValue("password")
		confirmPasswd := request.FormValue("confirmpassword")

		for _, cuser := range users {
			if cuser.Account == user.Account {
				errors["name"] = "存在同名用户"
				break
			}
		}

		if !utils.IsPhone(user.Tel) {
			errors["tel"] = "电话号码格式不正确"
		}

		if len(passwd) == 0 {
			errors["password"] = "密码不可以为空"
		} else if len(passwd) < 6 {
			errors["password"] = "密码必须是超过6个字符"
		}

		if passwd != confirmPasswd {
			errors["password"] = "两次输入密码不相同"
		} else {
			user.Passwd = utils.HashPasswd(passwd, "")
		}

		if len(errors) == 0 {
			err = user.CreateUser()
			if err == nil {
				http.Redirect(response, request, "/login/", http.StatusFound)
			} else {
				errors["submit"] = "提交失败"
			}
		}
	}

	tmpl := template.Must(template.New("register.html").ParseFiles("views/login/register.html"))
	tmpl.Execute(response, struct {
		User   models.User
		Errors map[string]string
	}{
		User:   user,
		Errors: errors,
	})
}
