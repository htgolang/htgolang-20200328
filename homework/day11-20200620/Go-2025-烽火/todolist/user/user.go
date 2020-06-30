package user

import (
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"todolist/config"
	"todolist/db"
	"todolist/utils"
)

type User struct {
	ID       int
	Name     string
	Status   int
	Password string
	Flag     int
}

func NewUser() *User {
	return &User{}
}

func init() {
	// 设置log
	logfile, _ := os.OpenFile("user.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	defer logfile.Close()
	log.SetOutput(logfile)

}

func GetAccounts() []User {
	accounts := make([]User, 0)
	rows, err := db.Config.DB.Query(config.SqlQueryAllUser)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Status)
		if err != nil {
			log.Println(err)
		} else {
			accounts = append(accounts, user)
		}
	}

	return accounts
}

//新增用户
func (u *User) AddUser(w http.ResponseWriter, r *http.Request) {
	var (
		errMsg = make(map[string]string)
		user   User
	)

	if r.Method == http.MethodPost {
		name := strings.TrimSpace(r.PostFormValue("name"))
		password1 := strings.TrimSpace(r.PostFormValue("password1"))
		password2 := strings.TrimSpace(r.PostFormValue("password2"))

		user.Name = name
		//检查用户名
		if err := utils.CheckUserName(name); err != nil {
			errMsg["username"] = err.Error()
		}
		//检查密码
		md5password, err := utils.CheckUserPassword(password1, password2)

		if err == nil {
			user.Password = md5password
		} else {
			errMsg["password"] = err.Error()
		}

		if len(errMsg) == 0 {
			db.Config.DB.Exec(config.SqlCreateUser, user.Name, user.Password)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}

	tpl := template.Must(template.ParseFiles("views/user/adduser.html"))
	tpl.ExecuteTemplate(w, "adduser.html", struct {
		User
		Error map[string]string
	}{user, errMsg})

}

// 查询用户
func (u *User) QueryUser(w http.ResponseWriter, r *http.Request) {
	var (
		user  User
		error string
		rt    bool
	)
	uid := r.FormValue("uid")
	row := db.Config.DB.QueryRow(config.SqlQueryUserInfo, uid)
	err := row.Scan(&user.ID, &user.Name, &user.Status, &user.Password)

	if err != nil {
		log.Println(err)
		error = "用户不存在"
	} else {
		rt = true
	}

	funcs := template.FuncMap{
		"status": func(index int) string {
			return config.UserStatusMap[index]
		},
	}

	tpl := template.Must(template.New("tpl").Funcs(funcs).ParseFiles("views/user/queryuser.html"))
	tpl.ExecuteTemplate(w, "queryuser.html", struct {
		User
		Error string
		OK    bool
	}{user, error, rt})
}
