package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// 1.channel是一种类型，一种引用类型。声明通道类型的格式如下：
	// var 变量 chan 元素类型
	var ch1 chan int
	var ch2 chan string
	// 通道是引用类型，通道类型的空值是nil。
	fmt.Println(ch1, ch2)

	// 2.声明的通道后需要使用make函数初始化之后才能使用。
	// make(chan 元素类型, [缓冲大小])
	// ch4 := make(chan int, 10) //带缓冲区
	// ch5 := make(chan bool)    //不带缓冲区
	// fmt.Println(ch4, ch5)

	// 3.channel操作
	// 通道有发送（send）、接收(receive）和关闭（close）三种操作。
	// 发送和接收都使用<-符号。
	ch := make(chan int)
	// ch <- 10 //发送10
	// x := <-ch // 从ch中接收值并赋值给变量x
	// <-ch      // 从ch中接收值，忽略结果
	// close(ch) // 关闭

	// ch <- 10 时会死锁  因为无缓冲的通道只有在有人接收值的时候才能发送值。
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-ch
		fmt.Println("后台goroutine从通道ch中取到了", x)
	}()
	ch <- 10 //发送10
	fmt.Println("10发送到通道ch中了")
	fmt.Println(ch)
	wg.Wait()
}
