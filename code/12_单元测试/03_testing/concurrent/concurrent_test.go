package main

import (
	"fmt"
	"testing"
)

// go test -v -bench=.   //全部 可指定具体

func BenchmarkConcurrentAtomicadd(b *testing.B) {
	b.ResetTimer()
	fmt.Println(b.N) // b.N 循环次数,运行时,基准测试框架将自动分配
	for i := 0; i < b.N; i++ {
		ConcurrentAtomicadd()
	}
}

func BenchmarkConcurrenMutexadd(b *testing.B) {
	b.ResetTimer()
	fmt.Println(b.N) // b.N 循环次数,运行时,基准测试框架将自动分配
	for i := 0; i < b.N; i++ {
		ConcurrenMutexadd()
	}
}