package main

import "fmt"

func main() {
	f1 := 1.23234
	fmt.Printf("%T\n", f1)
	fmt.Printf("%f\n", f1)
	f2 := float32(1.232345)
	fmt.Printf("%T\n", f2)
	fmt.Printf("%f\n", f2)

	// 布尔类型变量的默认值为false。
	// Go 语言中不允许将整型强制转换为布尔型.
	// 布尔型无法参与数值运算，也无法与其他类型进行转换。
}
