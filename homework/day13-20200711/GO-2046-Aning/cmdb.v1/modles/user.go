package modles

import (
	"cmdb/utils"
	gosql "database/sql"
	"time"
)

const (
	sqlQueryByName = "select id,name,password from user where name=? "
	sqlQuery       = "select id,staff_id,name,nickname,tel,email,addr,status,create_at,updated_at,deleted_at from user"
)

type User struct {
	ID         int
	Name       string
	StaffId    string
	NickName   string
	Password   string
	Gender     int
	Tel        string
	Addr       string
	Email      string
	Department string
	Status     int
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
}

func GetUserByName(name string) *User {
	user := &User{}
	// fmt.Println(name)
	err := db.QueryRow(sqlQueryByName, name).Scan(&user.ID, &user.Name, &user.Password)
	// fmt.Println(err)
	if err == nil {
		// fmt.Println("modles GetUserByName")
		return user
	}
	return nil
}
func (u *User) ValidPassword(password string) bool {
	// fmt.Println(password, u.Password)
	return u.Password == utils.Md5Text(password)
}

func QueryUser(q string) []*User {
	users := make([]*User, 0, 10)
	sql := sqlQuery
	var (
		rows *gosql.Rows
		err  error
	)
	params := []interface{}{}
	// q = utils.Like(q)
	if q != "" {
		sql += " where staff_id like ? ESCAPE '/' OR name like ? ESCAPE '/' OR nickname like ? ESCAPE '/' OR tel like ? ESCAPE '/' OR email like ? ESCAPE '/' OR addr like ? ESCAPE '/' OR department like ? ESCAPE '/' "
		params = append(params, q, q, q, q, q, q, q)
	}
	rows, err = db.Query(sql, params...)

	// if q != "" {
	// 	sql += " where staff_id like ? ESCAPE '/' OR name like ? ESCAPE '/' OR nickname like ? ESCAPE '/' OR tel like ? ESCAPE '/' OR email like ? ESCAPE '/' OR addr like ? ESCAPE '/' OR department like ? ESCAPE '/' "
	// 	fmt.Println(sql)
	// 	rows, err = db.Query(sql, q, q, q, q, q, q, q)
	// 	fmt.Println(err)
	// } else {
	// 	rows, err = db.Query(sql)
	// }

	if err != nil {
		return nil
	}
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.StaffId, &user.Name, &user.NickName, &user.Tel, &user.Email, &user.Addr, &user.Status, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		// fmt.Println(err)
		if err == nil {
			users = append(users, user)
		}
	}
	return users
}

func (u *User) GenderText() string {
	if u.Gender == 0 {
		return "女"
	}
	return "男"
}
func (u *User) StatusText() string {
	switch u.Status {
	case 0:
		return "正常"
	case 1:
		return "锁定"
	case 2:
		return "离职"
	}
	return "未知"
}
