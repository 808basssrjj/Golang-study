package main

import "fmt"

func main() {
	n := 100
	fmt.Printf("%T\n", n)

	s := "哈哈哈"
	fmt.Printf("%s\n", s)

	//Printf("格式化字符串"， 值)
	// %T 变量类型
	// %d 十进制数
	// %d 二进制数
	// %o 八进制数
	// %x 十六进制数
	// %c 字符
	// %s 字符串
	// %p 指针
	// %v 值
	// %f 浮点数

	// 获取输入 fmt.Scan
	// var s1 string
	// fmt.Scan(&s1)
	// fmt.Println("用户输入的内容时：", s1)

	// Scanf
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
	fmt.Scanln("1:%s 2:%d 3:%t", &name, &age, &married)

}
