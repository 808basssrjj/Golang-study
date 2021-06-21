package main

import "fmt"

// 用来资源清理、文件关闭、解锁及记录时间等
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("defer1") // 函数即将返回时执行defer后的语句
	defer fmt.Println("defer2") // 多个先进后出
	fmt.Println("end")
}

// defer执行时机
// go中的return不是原子操作，它分为给返回值赋值和RET指令两步
// defer语句执行的时机就在返回值赋值操作后，RET指令执行前
func f1() int {
	x := 5
	defer func() {
		x++ //修改的是x不是返回值
	}()
	return x //1. res=x=5 //2.x=6 3.return res
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值=x(指向x) defer改为了x=6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 返回值=y=x=5(指向y)  defer改为了x=6
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 返回值=x=5(指向x)   defer修改的是匿名函数的x副本
}

// 面试题   defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	deferDemo()

	// 执行事件
	fmt.Println(f1()) //5
	fmt.Println(f2()) //6
	fmt.Println(f3()) //5
	fmt.Println(f4()) //5

	// 面试题
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
	// 相当于
	// calc("A", 1, 2)  defer calc("AA", 1, 3)
	// calc("B", 10, 2) defer calc("BB", 10, 12)
	// defer calc("BB", 10, 12)
	// defer calc("AA", 1, 3)
}
