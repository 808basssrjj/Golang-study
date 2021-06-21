package main

import "fmt"

// 值接收者和指针接收者实现接口的区别

type Mover interface {
	move()
}
type dog struct{}
type cat struct{}


// 1.值接收者实现接口
func (d dog) move() {
	fmt.Println("狗会跑")
}

// 1.指针接收者实现接口
func (c *cat) move() {
	fmt.Println("猫会跑")
}

func main() {
	// 1.
	var x Mover
	var wangcai = dog{} // 旺财是dog类型
	x = wangcai         // x可以接收dog类型
	var fugui = &dog{}  // 富贵是*dog类型
	x = fugui           // x可以接收*dog类型
	x.move()

	// 2.
	var x1 Mover
	// var cat1 = cat{} // cat1是cat类型
	// x1 = cat1         // x1不可以接收dog类型
	var cat2 = &cat{}  // cat2是*cat类型
	x1 = cat2           // x1可以接收*dog类型
	x1.move()
}
