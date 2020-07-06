package controller

import "net/http"

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login/", http.StatusFound)
	})
	http.HandleFunc("/login/", Login)
	http.HandleFunc("/register/", Register)
	http.HandleFunc("/menu/", Menu)
	http.HandleFunc("/task/list/", ListTasks)
	http.HandleFunc("/task/add/", AddTask)
	http.HandleFunc("/task/del/", DelTask)
	http.HandleFunc("/task/mod/", ModTask)
	http.HandleFunc("/user/list/", ListUsers)
	http.HandleFunc("/user/add/", AddUser)
	http.HandleFunc("/user/del/", DelUser)
	http.HandleFunc("/user/mod/", ModUser)
	http.HandleFunc("/user/passwd/", PasswdUser)
	// 添加js和css
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("views/css/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("views/js/"))))
}
