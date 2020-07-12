package util

import (
	"strconv"
	"time"
)

func StopTimeFormat(stoptime time.Time) string{
	stop := stoptime.Format("2006-01-02")

	if stop == "0001-01-01" {
		return "无"
	}
	return stop
}

func IndexLeft(index int) string{
	num := strconv.Itoa(index-1)
	return  num
}
func IndexRight(index int) string{
	num := strconv.Itoa(index+1)
	return  num
}
