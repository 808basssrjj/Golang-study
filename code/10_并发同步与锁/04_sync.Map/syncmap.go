package main

import (
	"fmt"
	"sync"
)

// var m = make(map[int]int)
var m2 = sync.Map{}

// func get(key int) int {
// 	return m[key]
// }

// func set(key int, value int) {
// 	m[key] = value
// }

// func main() {
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 20; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			set(i, i+100)
// 			fmt.Printf("key:%v value %v\n", i, get(i))
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()

// 	// 报错 原生map不支持并发操作
// 	// sync包中提供了一个开箱即用的并发安全版map–sync.Map
// 	// 同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。
// }
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			m2.Store(i, i+100)
			value, _ := m2.Load(i)
			fmt.Printf("key:%v value %v\n", i, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
