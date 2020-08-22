package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "10.0.0.2:6379", redis.DialPassword("golang@2020"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	keys, _ := redis.Strings(conn.Do("KEYS", "*"))
	fmt.Println(keys)

	conn.Do("SET", "cmdb:test:starttime", time.Now().Format("2006-01-02 15:04:05"), "NX")

	starttime, _ := redis.String(conn.Do("GET", "cmdb:test:starttime"))
	fmt.Println(starttime)

	for i := 0; i < 10; i++ {
		conn.Do("LPUSH", "cmdb:test:tasks", i)
	}

	for i := 0; i < 20; i++ {
		fmt.Println(redis.String(conn.Do("RPOP", "cmdb:test:tasks")))
	}
	// for i := 0; i < 5; i++ {
	// 	// key value
	// 	// key value2
	// 	fmt.Println(redis.StringMap(conn.Do("BRPOP", "cmdb:test:tasks", 0)))
	// }

	conn.Do("HMSET", redis.Args{}.Add("cmdb:test:user").Add("name").Add("kk2"))
	user, _ := redis.StringMap(conn.Do("HGETALL", "cmdb:test:user"))
	fmt.Println(user)
}
