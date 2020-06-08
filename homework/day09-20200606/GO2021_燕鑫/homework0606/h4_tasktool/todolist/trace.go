package todolist

import (
	"fmt"
	"time"
)

func trace(t1 time.Time) {
	fmt.Printf("(Time elapsed :%v)\n", time.Now().Sub(t1))
}

func begin() time.Time {
	return time.Now()
}

