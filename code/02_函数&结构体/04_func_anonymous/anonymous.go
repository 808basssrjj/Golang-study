package main

import "fmt"

func main() {
	// 匿名函数（多用在函数中）
	// 多用来实现回调函数和闭包
	f1 := func(x, y int) {
		fmt.Println(x, y)
	}
	f1(10, 20)

	// 只调用一次，可写成立即执行
	func(x, y int) {
		fmt.Println(x, y)
	}(100, 200)
}
