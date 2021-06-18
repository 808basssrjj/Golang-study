package main

import "fmt"

//定义
// func 函数名(参数)(返回值){
//     函数体
// }

func f1() {
	fmt.Println("hello")
}
func f2(s string) {
	fmt.Println(s)
}
// 参数类型相同可以简写
func f3(x, y int) int {
	sum := x + y
	return sum
}
// 可变参数
func f4(title string, y ...int) int {
	// y是一个int类型的切片
	fmt.Println(y)
	return 1
}
// 多个返回值
func f5(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}
// 返回值命名
func f6(x, y int) (sum int) {
	sum = x + y
	return
}

func main() {
	f1()
	f1()
	f2("迪迦")
	fmt.Println(f3(1, 2))
	f4("hh", 1, 2, 3, 4)
	fmt.Println(f5(1, 2))
	fmt.Println(f6(10, 20))
}
