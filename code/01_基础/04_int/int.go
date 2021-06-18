package main

import "fmt"

func main() {
	i1 := 101
	fmt.Printf("%d\n", i1)

	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	// 十六进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)

	// 查看类型
	fmt.Printf("%T\n", i3)
}
