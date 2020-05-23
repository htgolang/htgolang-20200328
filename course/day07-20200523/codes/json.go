package main

import (
	"encoding/json"
	"fmt"
	"time"
)

const TimeLayout = "2006-01-02 15:04:05"

type Birthday time.Time

func (b *Birthday) UnmarshalText(txt []byte) error {
	// 反序列化 => text => 数据类型
	if txt == nil {
		*b = Birthday(time.Unix(0, 0))
	}
	if t, err := time.Parse(TimeLayout, string(txt)); err != nil {
		return err
	} else {
		*b = Birthday(t)
		return nil
	}
}

func (b Birthday) MarshalText() ([]byte, error) {
	// 序列化 => 数据类型 => text
	return []byte(time.Time(b).Format(TimeLayout)), nil
}

func (b Birthday) String() string {
	return time.Time(b).Format(TimeLayout)
}

type User struct {
	ID       int `json:"id"`
	Name     string
	Password string `json:"-"`
	Sex      bool   `json:"sex,omitempty"` //name,omitempty,string //name,type
	Birthday *Birthday
}

func main() {
	now := Birthday(time.Now())

	users := []User{
		{1, "kk", "123", true, &now},
		{2, "魏超", "1234", true, &now},
		{3, "燕鑫", "456787", false, &now},
	}

	bytes, _ := json.MarshalIndent(users, "", "\t")
	fmt.Println(string(bytes))

	var birthday Birthday = Birthday(now)
	ctx, _ := json.Marshal(birthday)
	fmt.Println(string(ctx))

	var bt Birthday
	fmt.Println(json.Unmarshal([]byte(`"2020-01-05 10:20:10"`), &bt))
	fmt.Println(time.Time(bt))
	fmt.Println(bt)
}
