package main

import "fmt"

// 递归:函数自己调用自己
// 要有一个明确的退出条件

//3! = 3*2*1 	= 3*2!
//4! = 4*3*2*!	= 4*3!
func f1(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f1(n-1)
}

func main() {
	ret := f1(5)
	fmt.Println(ret)
}
