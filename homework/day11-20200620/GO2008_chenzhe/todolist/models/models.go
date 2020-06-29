package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	//"todolist/until/task"
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
	sqlTasks = "select task.id, task.task_name,task.content, task.status, task.starttime,user.id as user_id, user.name as user_name, task.completetime from task left join user on task.user_id=user.id"
	sqlTaskDel = "delete from user where id=?"
	sqlTaskAdd = "insert into task(`user_id`,`task_name`,`content`,`status`,`starttime`,`completetime`) values (?,?,?,?,?,?)"
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
	Content		string
	StatusId	int
	Status       string
	StartTime    string
	CompleteTime string
	//User         *string
	User User
}

type TaskData struct {
	//Id string
	Id int
	Name string
	Content string
	Status string
	Starttime string
	Completetime string
	Err map[string]string
	User []User
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

func ListTask(w http.ResponseWriter, r *http.Request){

	tpl,_ := template.ParseFiles("views/task.html","views/task.html")
	data := make([]Task,0)
	rows,err :=db.Query(sqlTasks)
	if err !=nil{
		log.Println(err)
	}
	var task_name,task_starttime,task_completetime,user_name,task_content string
	var task_id,task_status,user_id int
	for rows.Next(){
		rows.Scan(&task_id,&task_name,&task_content,&task_status,&task_starttime,&user_id,&user_name,&task_completetime)
		if task_completetime == ""{
			task_completetime = "未结束"
		}

		data = append(data,Task{
			ID:task_id,
			Name:task_name,
			Content:task_content,
			StatusId:task_status,
			Status:statusMap[task_status],
			StartTime:task_starttime,
			CompleteTime:task_completetime,
			User:User{
				Id:user_id,
				Username:user_name,
			},
		})

	}
	tpl.ExecuteTemplate(w,"task.html",data)

}

func DeleteTask(w http.ResponseWriter, r *http.Request){
	id := r.FormValue("id")
	db.Exec(sqlTaskDel,id)
	http.Redirect(w,r,"/task",http.StatusFound)
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	tpl,_ := template.ParseFiles("views/taskadd.html","views/taskadd.html")
	rows,err := db.Query(sqlUsers)
	if err != nil{
		fmt.Println(err)
	}
	Users := make([]User,0)
	var id int
	var name string
	for rows.Next() {
		rows.Scan(&id, &name)
		Users = append(Users,User{
			Id:id,
			Username:name,
		})
	}
	data := TaskData{
		User:Users,
	}
	if r.Method == http.MethodGet{

		tpl.ExecuteTemplate(w,"taskadd.html",data)

	}
	if r.Method == http.MethodPost{
		r.ParseForm()
		errs := make(map[string]string)
		id_str := r.PostForm.Get("id")
		id,_ :=strconv.Atoi(id_str)
		name := r.PostForm.Get("name")
		if strings.TrimSpace(name)=="" {
			errs["name"] = "名称有误或不能为空"
		}
		content := r.PostForm.Get("content")
		if strings.TrimSpace(content)=="" {
			errs["content"] = "内容有误或不能为空"
		}
		status := r.PostForm.Get("status")
		starttime := r.PostForm.Get("starttime")
		if strings.TrimSpace(starttime)=="" {
			errs["starttime"] = "时间有误或不能为空"
		}
		completetime := r.PostForm.Get("completetime")
		if status == "3"{
			if strings.TrimSpace(completetime)=="" {
				errs["completetime"] = "时间有误或不能为空"
			}
		}
		if len(errs) !=0{

			tpl.ExecuteTemplate(w,"taskadd.html",TaskData{
				Id:id,
				Name:name,
				Content:content,
				Status:status,
				Starttime:starttime,
				Completetime:completetime,
				Err:errs,
				User:Users,

			})
		}else {
			res, err := db.Exec(sqlTaskAdd, id, name, content, status, starttime, completetime)
			fmt.Println(res)
			fmt.Println(err)
			http.Redirect(w, r, "/task", http.StatusFound)
		}
	}

}