package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "'D:\\Go\\src\\day01'"
	fmt.Println(path)

	// 多行字符串
	s1 := `
		a
		b
		c
	`
	fmt.Println(s1)

	// 字符串相关操作
	// 1.长度
	fmt.Println(len(path))

	// 2.字符串拼接
	name := "张三"
	world := "大帅比"
	ss := name + world
	fmt.Println(ss)
	fmt.Printf("%s%s", name, world)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)

	// 3.分割
	ret := strings.Split(path, "\\")
	fmt.Println(ret)

	// 4.包含
	fmt.Println(strings.Contains(ss, "张三"))

	// 5.前后缀
	fmt.Println(strings.HasPrefix(ss, "张三"))
	fmt.Println(strings.HasSuffix(ss, "张三"))

	// 6.字串出现的位置
	s2 := "abcdeb"
	fmt.Println(strings.Index(s2, "c"))
	fmt.Println(strings.LastIndex(s1, "b"))

	// 7.join操作
	fmt.Println(strings.Join(ret, "+"))

	// 8.统计字串出现的次数
	fmt.Println(strings.Count("seafood", "o"))

	// 9.不区分大小写比较
	fmt.Println(strings.EqualFold("golang", "GOLANg"))

	// 10.替换  n=-1代表全部替换
	s3 := "go go hhh"
	fmt.Println(strings.Replace(s3, "go", "golang", -1))

	// 11.大小写转换
	fmt.Println(strings.ToLower("GOLANG"))
	fmt.Println(strings.ToUpper("golang"))

	// 12.去除左右两边空格
	fmt.Println(strings.TrimSpace(" h sd hada "))
	// 13.去掉左右/左/右指定字符
	fmt.Println(strings.Trim(",golang,", ","))
	fmt.Println(strings.TrimLeft(",golang,", ","))
	fmt.Println(strings.TrimRight(",golang,", ","))
}
