package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello go")
}

// 在程序启动时，Go程序就会为main()函数创建一个默认的goroutine
func main() {
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main")
	// 当main()函数返回的时候该goroutine就结束了，所有在main()函数中启动的goroutine会一同结束，
	time.Sleep(time.Second)
}
