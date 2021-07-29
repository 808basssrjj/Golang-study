package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("这是一条普通日志")
	v := "普通"
	log.Printf("这也是一条%v日志", v)
	//log.Fatalln("这是一条会fatal的日志")
	//log.Panicln("这是一条会panic的日志")
	//Fatal系列函数会在写入日志信息后调用os.Exit(1)。Panic系列函数会在写入日志信息后panic。

	// 2.flag选项
	//const (
	//	// 控制输出日志信息的细节，不能控制输出的顺序和格式。
	//	// 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	//	Ldate         = 1 << iota     // 日期：2009/01/23
	//	Ltime                         // 时间：01:23:23
	//	Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
	//	Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
	//	Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
	//	LUTC                          // 使用UTC时间
	//	LstdFlags     = Ldate | Ltime // 标准logger的初始值
	//)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")

	// 3.配置日志前缀
	//func Prefix() string
	//func SetPrefix(prefix string)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("日志日志")
	log.SetPrefix("[迪迦]")
	log.Println("奥特曼奥特曼")

	// 4.配置日志输出位置
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file failed err:", err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("普通日志")
	log.SetPrefix("[gogogo]")
	log.Println("langlanglang")
}
