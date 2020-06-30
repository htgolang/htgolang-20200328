package config

const (
	DbDriver   = "mysql"
	DbUser     = "devon"
	DbPassword = "golang@2020"
	DbName     = "todolist"
	DbHost     = "127.0.0.1"
	DbPort     = 32769
)

const (
	SqlCreateTask            = `insert into task(name, content, start_time, deadline_time, user) values(?,?,?,?,?)`
	SqlCreateTaskWithCt      = `insert into task(name, content, start_time, end_time, deadline_time, user) values(?,?,?,?,?,?)`
	SqlQueryAllTask          = "select task.id, task.name, task.status, start_time, end_time, deadline_time, user.name, content from task left join user on task.user=user.id"
	SqlDeleteTask            = `delete from task where id=?`
	SqlQueryTask             = `select id, name, status, start_time, deadline_time, content, user from task where id=?`
	SqlQueryTaskWithUserName = `select task.name, task.status, start_time, end_time, deadline_time, content, user.name from task  left join user on task.user=user.id where task.id=?`
	SqlUpdateTaskIncludeTime = `update task set name=?, status=?, start_time=?, deadline_time=?, end_time=?, content=?, user=? where id=?`
	SqlUpdateTask            = `update task set name=?, status=?, start_time=?, deadline_time=?, content=?,user=? where id=?`
	SqlQueryAllUser          = `select id, name, status from user`

	SqlCreateUser    = `insert into user(name, password) values(?,?)`
	SqlQueryUser     = `select name from user where name=?`
	SqlQueryUserInfo = `select * from user where id=?`
)

const (
	TimeLayout = "2006-01-02 15:04"
	DateLayout = "2006-01-02T15:04:05Z"
)

var (
	StatusMap = map[int]string{
		0: "新建",
		1: "正在进行",
		2: "暂停",
		3: "完成",
	}
	UserStatusMap = map[int]string{
		0: "在职",
		1: "离职",
	}
)
