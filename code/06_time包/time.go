package main

import (
	"fmt"
	"time"
)

func main() {
	// 1.时间类型
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// 2.时间戳  自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数。
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
	// time.Unix()函数可以将时间戳转为时间格式
	ret := time.Unix(1624341878, 0)
	fmt.Println(ret)

	// 3.时间间隔
	// const (
	// 	Nanosecond  Duration = 1
	// 	Microsecond          = 1000 * Nanosecond
	// 	Millisecond          = 1000 * Microsecond
	// 	Second               = 1000 * Millisecond
	// 	Minute               = 60 * Second
	// 	Hour                 = 60 * Minute
	// )
	fmt.Println(time.Second)

	// 4.事件操作
	// func (t Time) Add(d Duration) Time
	// func (t Time) Sub(u Time) Duration
	// func (t Time) Equal(u Time) bool
	// func (t Time) Before(u Time) bool
	// func (t Time) After(u Time) bool
	fmt.Println(now.Add(24 * time.Hour)) //加一天

	// 5.定时器
	// ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	// for i := range ticker {
	// 	fmt.Println(i)//每秒都会执行的任务
	// }

	// 6.时间格式化
	// 时间类型有一个自带的方法Format进行格式化，
	// 需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S
	// 而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）。
	fmt.Println(now.Format("2006-01-02"))                         // Y-m-d
	fmt.Println(now.Format("2006-01-02 15:04:05"))                // Y-m-d H:M:S  24小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan")) // 12小时制

	// 7.解析字符串格式的时间
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))

	n := 5
	fmt.Println("开始sleep")
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("五秒过去了")
}
