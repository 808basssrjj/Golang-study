package main

import (
	"sync"
	"sync/atomic"
)

func atomicCounter(counter *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		atomic.AddInt64(counter, 1)
	}
}

func mutexCounter(counter *int64, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		*counter++
		lock.Unlock()
	}
}

func ConcurrentAtomicadd() int64 {
	wg := sync.WaitGroup{}
	wg.Add(2)
	var counter int64 = 0
	go atomicCounter(&counter, &wg)
	go atomicCounter(&counter, &wg)
	wg.Wait()
	return counter
}

func ConcurrenMutexadd() int64 {
	wg := sync.WaitGroup{}
	wg.Add(2)
	var lock sync.Mutex
	var counter int64 = 0
	go mutexCounter(&counter, &wg, &lock)
	go mutexCounter(&counter, &wg, &lock)
	wg.Wait()
	return counter
}

func main() {
	println(ConcurrentAtomicadd())
	println(ConcurrenMutexadd())
}