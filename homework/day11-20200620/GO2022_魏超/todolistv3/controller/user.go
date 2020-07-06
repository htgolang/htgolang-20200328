package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	"todolist/models"
	"todolist/utils"
)

func ListUsers(response http.ResponseWriter, request *http.Request) {
	var (
		users []models.User
		tmpl  *template.Template
	)
	users = models.GetUsers()
	funcMap := template.FuncMap{
		"datetime":  utils.FormatDatetime,
		"headlesex": utils.FormatSex,
	}
	tmpl = template.Must(template.New("list.html").Funcs(funcMap).ParseFiles("views/user/list.html"))
	tmpl.Execute(response, struct {
		Users []models.User
	}{
		Users: users,
	})
}

func AddUser(response http.ResponseWriter, request *http.Request) {
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
				http.Redirect(response, request, "/user/list", http.StatusFound)
			} else {
				errors["submit"] = "提交失败"
			}
		}
	}

	tmpl := template.Must(template.New("add.html").ParseFiles("views/user/add.html"))
	err = tmpl.Execute(response, struct {
		User   models.User
		Errors map[string]string
	}{
		User:   user,
		Errors: errors,
	})
}

func DelUser(response http.ResponseWriter, request *http.Request) {
	var (
		user models.User
	)
	user.ID, _ = strconv.Atoi(request.FormValue("id"))
	user.DeleteUser()
	http.Redirect(response, request, "/user/list/", http.StatusFound)
}

func ModUser(response http.ResponseWriter, request *http.Request) {
	var (
		err    error
		user   models.User
		users  []models.User
		errors = make(map[string]string)
	)
	user.ID, _ = strconv.Atoi(request.FormValue("id"))

	if user.GetUserById() != nil {
		errors["submit"] = "用户不存在"
	}
	if request.Method == http.MethodGet {

	} else if request.Method == http.MethodPost {
		user.Name = request.FormValue("name")
		user.Sex, _ = strconv.Atoi(request.FormValue("sex"))
		account := request.FormValue("account")
		tel := request.FormValue("tel")
		user.Address = request.FormValue("address")
		users = models.GetUsers()

		if user.Account != account {
			var isSameName bool
			for _, ouser := range users {
				if ouser.Account == account {
					errors["account"] = "存在同名账号"
					isSameName = true
					break
				}
			}
			if !isSameName {
				user.Account = account
			}
		}

		if !utils.IsPhone(tel) {
			errors["tel"] = "电话号码格式不正确"
		} else {
			user.Tel = tel
		}

		if len(errors) == 0 {
			err = user.UpdateUser()
			if err == nil {
				http.Redirect(response, request, "/user/list", http.StatusFound)
			} else {
				errors["submit"] = "提交失败"
			}
		}
	}

	funcMap := template.FuncMap{
		"datetime": utils.FormatDatetime,
	}
	tmpl := template.Must(template.New("edit.html").Funcs(funcMap).ParseFiles("views/user/edit.html"))
	err = tmpl.Execute(response, struct {
		User   models.User
		Errors map[string]string
	}{
		User:   user,
		Errors: errors,
	})
	log.Println(err)
}

func PasswdUser(response http.ResponseWriter, request *http.Request) {
	var (
		user   models.User
		errors = make(map[string]string)
	)

	user.ID, _ = strconv.Atoi(request.FormValue("id"))
	if request.Method == http.MethodGet {

	} else if request.Method == http.MethodPost {
		oldPassword := request.FormValue("oldpassword")
		newPassword := request.FormValue("newpassword")
		confirmPasswd := request.FormValue("confirmpassword")

		if user.GetUserById() != nil {
			errors["submit"] = "用户不存在"
		}

		if oldPassword == "" {
			errors["oldpassword"] = "旧密码不可为空"
		}

		if user.Passwd != "" && !utils.ValidPassword(oldPassword, user.Passwd) {
			errors["oldpassword"] = "旧密码错误"
		}

		if len(newPassword) == 0 {
			errors["newpassword"] = "新密码不可为空"
		} else if len(newPassword) < 6 {
			errors["newpassword"] = "密码必须超过6个字符"
		}

		if user.Passwd != "" && newPassword != confirmPasswd {
			errors["submit"] = "密码不一致"
		}

		if len(errors) == 0 {
			user.Passwd = utils.HashPasswd(newPassword, "")
			if user.UpdateUser() == nil {
				http.Redirect(response, request, fmt.Sprintf("/user/mod/?id=%d", user.ID), http.StatusFound)
			} else {
				errors["submit"] = "更改失败"
			}
		}
	}

	tmpl := template.Must(template.New("password.html").ParseFiles("views/user/password.html"))
	err := tmpl.Execute(response, struct {
		User   models.User
		Errors map[string]string
	}{
		User:   user,
		Errors: errors,
	})
	log.Println(err)
}
