package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 时间
	now := time.Now()
	fmt.Printf("%T, %#v\n", now, now)

	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// unix时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	// => 字符串 2020-04-25 15:50:xx?
	// Printf 格式化字符串 占位符
	// 2006 4位数字的年
	// 01 2位数字的月
	// 02 2位数字的天
	// 03 12进制的小时
	// 15 24进制的小时
	// 04 2位数字的分钟
	// 05 2位数字的秒
	// 年月日 时分秒
	// 2006 01 02 15 04 05

	fmt.Println(now.Format("2006-01-02 03:04:05.000"))
	fmt.Println(now.Format("2006年01月02日 03:04:05"))
	fmt.Println(now.Format("2006年01月02日 15:04:05"))
	fmt.Println(now.Format("15:04:05 2006年01月02日"))
	fmt.Println(now.Format("15:04:05 02/01/2006"))

	// 生成时间
	year, month, day := 1990, time.February, 1
	time1990 := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	fmt.Println(time1990.Format("2006-01-02"))

	// unixtime
	startTime := time.Unix(0, 0)
	fmt.Println(startTime.Format("2006-01-02 15:04:05"))

	// 字符串
	cTime, err := time.Parse("2006-01-02", "2001-12-01")
	fmt.Println(cTime.Format("2006-01-02 15:04:05"), err)

	// cTime, err = time.Parse("2006-01-02", "2001/12/01")
	// fmt.Println(cTime.Format("2006-01-02 15:04:05"), err)

	// 时间区间
	// 生成时间区间
	// now - time
	dura := time.Since(cTime)

	fmt.Printf("%T, %#v\n", dura, dura)
	fmt.Println(dura)

	// time - now
	dura02 := time.Until(cTime)
	fmt.Println(dura02)

	dura03, err := time.ParseDuration("1h1m1s")

	fmt.Println(dura03, err)
	fmt.Println(dura03.Hours())
	fmt.Println(dura03.Minutes())
	fmt.Println(dura03.Seconds())
	fmt.Println(dura03.Nanoseconds())

	dura03, err = time.ParseDuration("1")
	fmt.Println(dura03, err)

	dayInterval, _ := time.ParseDuration("-24h1m")
	fmt.Println(now, now.Add(dayInterval))

	//a > b a < b
	yesterday := now.Add(dayInterval)

	fmt.Println(yesterday.After(now))
	fmt.Println(yesterday.Before(now))

	fmt.Println(yesterday.Sub(now))

	fmt.Println(time.Now())
	// 常量
	time.Sleep(time.Second * 3)
	fmt.Println(time.Now())
}
