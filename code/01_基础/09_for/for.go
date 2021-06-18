package main

import "fmt"

func main() {
	// 1. 基本格式
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 2.for循环的初始语句可以被忽略，但是初始语句后的分号必须要写
	// i := 5
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 3.for循环的初始语句和结束语句都可以省略 类似while
	// i := 0
	// for i < 3 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// 4.死循环
	// for {

	// }

	// 5.for range
	s := "Hello沙河"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

	// 打印9*9
	for i := 1; i < 10; i++ {
		for j := 1; j <=i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
