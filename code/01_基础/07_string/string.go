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
}
