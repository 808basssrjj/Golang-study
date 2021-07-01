package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func rd() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 5; i++ {
		r1 := rand.Int()    //int64
		r2 := rand.Intn(10) //0<=x<10
		fmt.Println(0-r1, 0-r2)
	}
}

// 使用了sync.WaitGroup来实现goroutine的同步
var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}
func main() {
	rd()
	for i := 0; i < 100; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束(计数器为0)

	// 每次打印的数字的顺序都不一致。这是因为10个goroutine是并发执行的，而goroutine的调度是随机的。
}

