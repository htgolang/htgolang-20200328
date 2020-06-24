package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)
const (
	dateTimeLayout = "2006-01-02 15:04:05"
)
const (
	dbDriver   = "mysql"
	dbUser     = "root"
	dbPassword = "123456"
	dbName     = "chenzhe"
	dbHost     = "10.0.0.129"
	dbPort     = 3306
)
const (
	sqlTasks = "select task.id, task.name, task.status, task.start_time, task.complete_time, task.deadline_time, user.name as user from task left join user on task.user=user.id"
	sqlUsers = "select id,name from user"
	sqlUserDel = "delete from user where id=?"
	sqlUser = "select name from user where name=?"
	sqlUserId = "select id,name from user where id=? Limit 1"
	sqlUserCheck = "select count(1) from user where name=? Limit 1"
	sqlUserAdd = "insert into user(`name`,`c_time`) values (?,?)"
	sqlUserUpdete="update user set name=? where id=?"
)


var (
	statusMap = map[int]string{
		0: "新建",
		1: "正在进行",
		2: "暂停",
		3: "完成",
	}
)

type User struct {
	Id int
	Username string
}

type Task struct {
	ID           int
	Name         string
	Status       string
	StartTime    *time.Time
	CompleteTime *time.Time
	DeadlineTime *time.Time
	User         *string
}
var (
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=PRC", dbUser, dbPassword, dbHost, dbPort, dbName)
	db,_ = sql.Open(dbDriver, dsn)
)


func Index(w http.ResponseWriter, r *http.Request)  {
	tpl,err := template.ParseFiles("views/index.html")
	if err != nil{
		log.Print(err)
	}
	tpl.ExecuteTemplate(w,"index.html",nil)
}

func ListUser(w http.ResponseWriter, r *http.Request)  {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=PRC", dbUser, dbPassword, dbHost, dbPort, dbName)
	//db,_ := sql.Open(dbDriver, dsn)
	rows,err := db.Query(sqlUsers)
	if err != nil{
		log.Println(err)
	}
	result := make([]User,0)
	for  rows.Next(){
		var name string
		var id int
		rows.Scan(&id,&name)
		result = append(result,User{Id:id,Username:name})
	}
	tpl,_ := template.ParseFiles("views/user.html")
	tpl.ExecuteTemplate(w,"user.html",result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	p :=r.FormValue("id")
	if p == ""{
		log.Println("必须传用户id")
	} else {
		_,err := db.Exec(sqlUserDel,p)
		if err !=nil{
			log.Println(err)
		}

	}
	http.Redirect(w,r,"/user",301)
}
func AddUser(w http.ResponseWriter, r *http.Request) {
	errors := make(map[string]string,20)
	tpl,_ := template.ParseFiles("views/useradd.html")
	if r.Method == http.MethodGet{
		tpl.ExecuteTemplate(w,"useradd.html",nil)
	}else if r.Method == http.MethodPost{
		name := strings.TrimSpace(r.PostFormValue("username"))
		rows,err := db.Query(sqlUser,name)
//名字重复或sql执行有问题
		if err != nil || rows.Next()||name == ""{
			errors["name"]="名称重复或为空"
			tpl.ExecuteTemplate(w,"useradd.html", struct {
				Username string
				Err map[string]string
			}{
				Username:name,
				Err:errors,
			})
			return
		}else {
			//插入用户数据
			db.Exec(sqlUserAdd,name,time.Now().Format(dateTimeLayout))
			http.Redirect(w,r,"/user",301)
		}
		}



}

func EditUser(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	tpl,_ := template.ParseFiles("views/usererr.html","views/useredit.html")
	if id ==""{
		tpl.ExecuteTemplate(w,"usererr.html","id不能为空")
	}else {
		row,err := db.Query(sqlUserId,id)
		if err !=nil{
			log.Println(err)
		}
		if !row.Next(){
			log.Println("没有这个id")
		}

		if r.Method==http.MethodGet{
			var id,name string
			row.Scan(&id,&name)
			tpl.ExecuteTemplate(w,"useredit.html", struct {
				Data map[string]string
				Err map[string]string
			}{
				Data: map[string]string{"id":id,"name":name},
			})

		}else if r.Method==http.MethodPost{
			num :=0
			name := r.FormValue("username")
			rows,err := db.Query(sqlUserCheck,name)
			if err !=nil{
				log.Println(err)
			}
			rows.Next()
			rows.Scan(&num)
			fmt.Println(num)
			if num == 0{
				db.Exec(sqlUserUpdete,name,id)
				http.Redirect(w,r,"/user",301)
			}else {
				tpl.ExecuteTemplate(w,"useredit.html", struct {
					Data map[string]string
					Err map[string]string
				}{
					Data: map[string]string{"id":id,"name":name},
					Err: map[string]string{"name":"名称重复或有误"},
				})
			}
			//fmt.Println(row ,err)
			//sql.ErrNoRows

		}

		//db.Exec(sqlUserUpdete,name,id)
		//http.Redirect(w,r,"/user",301)
	}

}