package main

import "fmt"

func main() {
	score := 95
	if score >= 95 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// 特殊写法：可以在 if 表达式之前添加一个执行语句
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
