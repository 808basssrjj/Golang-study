package main

import "fmt"

// 函数类型
func f1() {
	fmt.Println("aaa")
}
func f2() int {
	return 10
}
func f3(x, y int) int {
	return 10
}

// 函数可作为参数的类型
func f4(x func() int) {
	ret := x()
	fmt.Println(ret)
}

// 函数可作为返回值
func f5(x func() int) func(int, int) int {
	return f3
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)
	c := f3
	fmt.Printf("%T\n", c)

	f4(f2)
	f4(b)
}
