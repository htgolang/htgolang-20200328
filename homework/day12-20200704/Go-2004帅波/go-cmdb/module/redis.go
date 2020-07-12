package module

import (
	"encoding/json"
	"fmt"
)

func GetKey(email string) *User {
	var user User
	value ,err  := RedisClient.Get(email).Result()
	if err != nil {
		return nil
	}
	err = json.Unmarshal([]byte(value),&user)
	if err != nil {
		return  nil
	}
	return &user
}

func SetKey(user User) error{
	data ,err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("set序列化失败 %v\n",err)
	}
	err = RedisClient.Set(user.Email,data,0).Err()
	if err != nil {
		return fmt.Errorf("set redis key is err %v\n",err)
	}
	return  nil
}

func DelKey(email string) error {
	err := RedisClient.Del(email).Err()
	if err != nil {
		return  fmt.Errorf("redis del is err %v\n",err)
	}
	return nil
}