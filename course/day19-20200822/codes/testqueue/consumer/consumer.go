package consumer

import (
	"fmt"
	"log"
	"testqueue/conn"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

func Run() {
	rConn := conn.GetConn()
	defer rConn.Close()

	lockKey := "consumer:master:uid"
	uid := uuid.New().String()
	log.Println("uid", uid)
	go func() {
		rConn := conn.GetConn()
		defer rConn.Close()
		ticker := time.NewTicker(time.Second * 10)
		defer ticker.Stop()
		for {
			rt, err := rConn.Do("SET", lockKey, uid, "EX", 30, "NX")
			if err != nil || rt != "OK" {
				rt, err := redis.String(rConn.Do("GET", lockKey))
				if err == nil {
					if rt == uid {
						log.Printf("is Master 2")
						// 续时间
						rConn.Do("Expire", lockKey, 30)
					} else {
						log.Printf("not Master")
					}
				}
			} else {
				log.Printf("is Master 1")
			}
			<-ticker.C
		}
	}()
	for {
		if rt, err := redis.String(rConn.Do("Get", lockKey)); err != nil || rt != uid {
			log.Printf("is not master, no worker")
			time.Sleep(time.Second * 10)
			continue
		}

		values, err := redis.Strings(rConn.Do("BRPOP", "cmdb:test:works", 3))
		if err != nil {
			continue
		}
		fmt.Println(values[1])
	}
}
