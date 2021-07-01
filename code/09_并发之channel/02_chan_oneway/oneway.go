package main

import "fmt"

// 有的时候我们会将通道作为参数在多个任务函数间传递，
// 很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或只能接收。
// 可以使用单向通道 多用在函数参数中

func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	// chan<- int是一个只写单向通道（只能对其写入int类型值），可以对其执行发送操作但是不能执行接收操作；
	// <-chan int是一个只读单向通道（只能从其读取int类型值），可以对其执行接收操作但是不能执行发送操作。
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
