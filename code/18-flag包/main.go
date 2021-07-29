package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// 简单获取
	// 	if len(os.Args) > 0 {
	// 		for index, arg := range os.Args {
	// 			fmt.Printf("[%d]=%v\n", index, arg)
	// 		}

	// 	}

	// 1. flag.Type(flag名, 默认值, 帮助信息) * Type
	// name := flag.String("name", "张三", "姓名")
	// age := flag.Int("age", 18, "年龄")
	// married := flag.Bool("married", false, "婚否")
	// delay := flag.Duration("d", 0, "时间间隔")

	// 2.flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")

	// 3. 通过以上两种方法定义好命令行flag参数后，需要通过调用flag.Parse()来对命令行参数进行解析。
	// 	支持的命令行参数格式有以下几种：
	// -flag xxx （使用空格，一个-符号）
	// --flag xxx （使用空格，两个-符号）
	// -flag=xxx （使用等号，一个-符号）
	// --flag=xxx （使用等号，两个-符号）
	// 其中，布尔类型的参数必须使用等号的方式指定。
	flag.Parse()
	fmt.Println(name, age, married, delay)

	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())

}
