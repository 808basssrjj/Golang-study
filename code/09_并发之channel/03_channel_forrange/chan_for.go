package main

import "fmt"


// 当通道被关闭时，再往该通道发送值会引发panic，
// 从该通道取值的操作会先取完通道中的值，再然后取到的值一直都是对应类型的零值。那如何判断一个通道是否被关闭了呢？

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 启动goroutine将0-100发送到ch1
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 1.通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 2.通道关闭后会退出for range循环
		fmt.Println(i)
	}

	// 从上面的例子中我们看到有两种方式在接收值的时候判断该通道是否被关闭，
	// 不过我们通常使用的是for range的方式。使用for range遍历通道，当通道被关闭的时候就会退出for range。
}
