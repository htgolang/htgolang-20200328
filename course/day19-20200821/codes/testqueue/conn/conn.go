package conn

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func GetConn() redis.Conn {
	conn, err := redis.Dial("tcp", "10.0.0.2:6379", redis.DialPassword("golang@2020"))
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
