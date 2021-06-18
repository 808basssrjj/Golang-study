package main

import "fmt"

func main() {
	// 指针只能读，不能修改

	n := 18
	// 1.&取地址
	p := &n
	fmt.Printf("%T\n", p) //*int
	// 2.*根据地址取值
	v := *p
	fmt.Println(v)

	// 3. new()和make()
	// var a1 *int	//nil
	// *a1 = 100	//找不到
	// fmt.Println(*a1)

	//new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值
	var a2 = new(int) // new()申请一个内容地址
	fmt.Printf("%T", a2)
	fmt.Println(a2)
	fmt.Println(*a2)
	*a2 = 100
	fmt.Println(*a2)
	//make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，
	var b map[string]int
	b = make(map[string]int, 10)
	b["沙河娜扎"] = 100
	fmt.Println(b)

}
