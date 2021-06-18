package main

import "fmt"

// 声明变量 变量名推荐驼峰
// var name string
// var age int
// var isOK bool

// 批量声明  全局变量推荐
var (
	name string
	age  int
	isOK bool
)

func main() {
	name = "张三"
	age = 16
	isOK = true
	// 变量声明必须使用
	fmt.Print(isOK)
	fmt.Println()               // 空行
	fmt.Printf("name:%s", name) //占位符
	fmt.Println()
	fmt.Println(age) // 换行

	// 声明并赋值
	var s1 string = "who"
	fmt.Println(s1)
	// 变量推导
	var s2 = "zhangsan"
	fmt.Println(s2)
	// 简短变量声明(只能在函数中)  局部变量推荐
	s3 := "hhh"
	fmt.Println(s3)

	// 在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量（anonymous variable）。 匿名变量用一个下划线_表示，
	//_多用于占位，表示忽略值。
}
