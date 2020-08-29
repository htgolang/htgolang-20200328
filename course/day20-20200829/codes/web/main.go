package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	addr := "0.0.0.0:80"

	redisAddr, ok := os.LookupEnv("MY_REDIS_ADDR")
	if !ok {
		redisAddr = "127.0.0.1"
	}
	redisPort, ok := os.LookupEnv("MY_REDIS_PORT")
	if !ok {
		redisPort = "6379"
	}

	redisPassword, ok := os.LookupEnv("MY_REDIS_PASSWORD")
	if !ok {
		redisPassword = ""
	}

	redisClient, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisAddr, redisPort), redis.DialPassword(redisPassword))
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		visit, _ := redis.Int(redisClient.Do("INCR", "visit"))
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, "%s: %s: %d", hostname, time.Now().Format("2006-01-02 15:04:05"), visit)
	})
	http.ListenAndServe(addr, nil)
}
