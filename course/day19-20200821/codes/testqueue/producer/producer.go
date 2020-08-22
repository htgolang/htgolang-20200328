package producer

import (
	"fmt"
	"math/rand"
	"os"
	"testqueue/conn"
	"time"
)

func Run() {
	// 随机sleep n s 在队列cmdb:test:works 放入当前时间
	rConn := conn.GetConn()
	defer rConn.Close()

	for {
		t := time.Now().Format("2006-01-02 15:04:05")

		rConn.Do("LPUSH", "cmdb:test:works", fmt.Sprintf("%d:%s", os.Getpid(), t))
		time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	}
}
