package controller

import (
	"html/template"
	"net/http"

	"todolist/models"
	"todolist/utils"
)

func Login(response http.ResponseWriter, request *http.Request) {
	var (
		user   models.User
		errors = make(map[string]string, 2)
	)

	if request.Method == http.MethodGet {

	} else if request.Method == http.MethodPost {
		user.Account = request.FormValue("account")
		inputpasswd := request.FormValue("password")

		if user.Account == "" {
			errors["account"] = "用户名不可为空"
		}

		if inputpasswd == "" {
			errors["password"] = "密码不可为空"
		}

		if user.GetUserByAccount() != nil || user.ID == 0 {
			errors["account"] = "用户不存在"
		}

		if len(errors) == 0 {
			if utils.ValidPassword(inputpasswd, user.Passwd) {
				http.Redirect(response, request, "/menu/", http.StatusFound)
			} else {
				errors["login"] = "密码错误"
			}
		}
	}

	tmpl := template.Must(template.New("login.html").ParseFiles("views/login/login.html"))
	tmpl.Execute(response, struct {
		Errors map[string]string
	}{
		Errors: errors,
	})
}
