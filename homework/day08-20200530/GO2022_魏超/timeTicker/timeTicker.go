package timeTicker

import (
	"time"
)

// TimeTicker 间隔interval时间触发写入管道的数据，如果interval小于等于零，则返回一个nil的管道；deadline 设置截止时间，当deadline小于零，将不设置截止时间
func TimeTicker(interval time.Duration, deadline time.Duration) <-chan struct{} {
	ticker := make(chan struct{})
	limit := true
	if interval <= 0 {
		return nil
	}
	if deadline <= 0 {
		limit = false
	}
	deadlineTime := time.Now().Add(deadline)
	go func() {
		defer close(ticker)
		for {
			ticker <- struct{}{}
			time.Sleep(interval)
			if limit && time.Now().After(deadlineTime) {
				break
			}
		}
	}()
	return ticker
}
