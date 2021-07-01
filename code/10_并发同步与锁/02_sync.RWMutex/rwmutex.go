package main

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁是完全互斥的，但是有很多实际的场景下是"读多写少"的
// 这种场景下使用读写锁是更好的一种选择。读写锁在Go语言中使用sync包中的RWMutex类型。
// 读写锁分为两种：读锁和写锁。
// 1.当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁
// 2.如果是获取写锁就会等待；当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。

var (
	x  int64
	wg sync.WaitGroup
	// lock   sync.Mutex //互斥锁
	rwlock sync.RWMutex //读写互斥锁
)

func read() {
	rwlock.RLock()               //加读锁
	time.Sleep(time.Millisecond) //假设读操作耗时1毫秒
	rwlock.RUnlock()             //解读锁
	wg.Done()
}

func write() {
	rwlock.Lock() //加写锁
	x = x + 1
	time.Sleep(time.Millisecond * 10) //假设读操作耗时10毫秒
	rwlock.Unlock()                   //解写锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	wg.Wait()
	fmt.Println(time.Since(start))
	fmt.Println(x)
}
