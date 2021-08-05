package main

import (
	"fmt"
	"strings"
)

//闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境
//匿名函数内部使用外部变量
func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

// 实际应用：把f2传入f1中执行
func f1(f func()) {
	fmt.Println("this is f1")
	f()
}
func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 解决:定义一个f3 返回值为 f1的参数
func f3(f func(int, int), x, y int) func() {
	return func() {
		f(x, y)
	}
}


// 传入一个后缀,返回带后缀的文件名
func makeSuffix(suffix string) func(string) string {
	return func(filename string) string {
		if strings.HasSuffix(filename, suffix) {
			return filename
		}
		return filename + suffix
	}
}

func main() {
	ret := adder(100) //ret是adder()返回的函数  可以ret()调用
	ret2 := ret(200)
	fmt.Println(ret2)

	res := f3(f2, 100, 200)
	// res()
	f1(res)

	mkJpg := makeSuffix(".jpg")
	fmt.Println(mkJpg("aabb.png"))
}
