package main

import "fmt"

// 常量在定义的时候必须赋值
// const pi = 3.1415
// const e = 2.7182

// 批量声明  如果省略了值则表示和上面一行的值相同。
const (
	n1 = 100
	n2
	n3
)

// iota是go语言的常量计数器，只能在常量的表达式中使用。
// 在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次
const (
	a1 = iota //0
	a2        //1
	a3        //2
)

//跳过
const (
	b1 = iota //0
	b2        //1
	_
	b4 //3
)

// 插队
const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
)
const c5 = iota //0

// 多个一行
const (
	d1, d2 = iota + 1, iota + 2 //1,2
	d3, d4                      //2,3
	d5, d6                      //3,4
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	// 变量声明必须使用
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)

	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)
	fmt.Println("c5:", c5)

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)
	fmt.Println("d5:", d5)
	fmt.Println("d6:", d6)
}
