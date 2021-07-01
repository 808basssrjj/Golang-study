package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex //互斥锁

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() //加锁
		x = x + 1
		lock.Unlock() //解锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

// 开启了两个goroutine去累加变量x的值，这两个goroutine在访问和修改x变量的时候就会存在数据竞争，导致最后的结果与期待的不符。
// Go语言中使用sync包的Mutex类型来实现互斥锁
