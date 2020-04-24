package task

import (
	"encoding/json"
	"io"
	"os"
)

const dbFile = "taskinfo.db"

func ReadDB() string {
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

func WriteDB(tasksInfo map[string]map[string]string) error {
	data, err := json.Marshal(tasksInfo)
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
