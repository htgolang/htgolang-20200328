package user

import (
	"encoding/json"
	"io"
	"os"
)

// ReadDB 读取存储的数据
func ReadDB(dbFile string) string {
	file, err := os.Open(dbFile)
	if err != nil {
		return ""
	}
	defer file.Close()
	var data []byte
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		data = append(data, buf[:n]...)
		if err == io.EOF {
			break
		}
	}
	return string(data)
}

// WriteDB 写入数据
func WriteDB(usersInfo map[int]*UserType, dbFile string) error {
	data, err := json.Marshal(usersInfo)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(dbFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
