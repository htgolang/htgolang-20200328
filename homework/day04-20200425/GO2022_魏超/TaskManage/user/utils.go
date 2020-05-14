package user

import "fmt"

// SaveDB 保存数据
func SaveDB(usersInfo map[int]*UserType, dbFile string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("新建任务失败,%s", err)
		}
	}()
	defer func() {
		err := WriteDB(usersInfo, dbFile)
		if err != nil {
			panic(err)
		}
	}()
}

// GenerateUserID 生成用户信息的ID
func GenerateUserID(usersInfo map[int]*UserType) int {
	var rt int
	for _, user := range usersInfo {
		id := user.ID
		if rt < id {
			rt = id
		}
	}
	return rt + 1
}
