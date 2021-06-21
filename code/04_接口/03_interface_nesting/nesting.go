package main

import "fmt"

// 接口嵌套

type sayer interface {
	say()
}
type mover interface {
	move()
}

// 嵌套
type animal interface {
	sayer
	mover
}

type cat struct {
	name string
}

func (c *cat) say() {
	fmt.Printf("%s会叫\n", c.name)
}
func (c *cat) move() {
	fmt.Printf("%s会跑\n", c.name)
}
func main() {
	var x animal
	cat := &cat{name: "小白"}
	x = cat
	x.move()
	x.say()
}
