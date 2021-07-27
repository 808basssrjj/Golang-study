package main

import (
	"context"
	"fmt"
)

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 0
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func main() {
	//WithCancel返回带有新Done通道的父节点的副本。
	//当调用返回的cancel函数或当关闭父上下文的Done通道时，将关闭返回上下文的Done通道，无论先发生什么情况。
	//取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。
	//func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
