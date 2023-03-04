package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

// 通过SingleFlight可以将对同一个Key的并发请求进行合并，
// 只让其中一个请求进行查询，其他请求共享同一个结果，可以很大程度提升并发能力。

type (
	call struct {
		val interface{}    // 返回的结果
		err error          //操作发生的错误
		wg  sync.WaitGroup // 用于实现通过1个 call，其他 call 阻塞
	}

	SingleFlight struct {
		calls map[string]*call // 不同的 call 对应不同的 key
		lock  sync.Mutex       // 利用锁控制请求
	}
)

func NewSingleFlight() *SingleFlight {
	return &SingleFlight{
		calls: make(map[string]*call),
	}
}

func (s *SingleFlight) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	c, done := s.createCall(key)
	if done { // 有其他协程发起call请求, 等待结果直接返回
		return c.val, c.err
	}

	// 是第一个发起请求的协程
	s.makeCall(c, key, fn)
	return c.val, c.err
}

func (s *SingleFlight) createCall(key string) (*call, bool) {
	s.lock.Lock() // map不安全要上锁
	// 有其他协程在进行请求, 当前协程只需要等待
	if c, ok := s.calls[key]; ok {
		s.lock.Unlock()
		c.wg.Wait()
		return c, true
	}

	// 第一个发起请求的协程, 初始化call, 并执行add阻塞其他协程
	c := new(call)
	c.wg.Add(1)
	s.calls[key] = c
	s.lock.Unlock()

	return c, false
}

func (s *SingleFlight) makeCall(c *call, key string, fn func() (interface{}, error)) {
	// 真正call请求要做的事情
	c.val, c.err = fn()

	// 删除 map 中的 key，使得下次发起请求可以获取新的值。
	// 调用wg.Done()，让之前阻塞的协程全部获得结果并返回。
	s.lock.Lock()
	delete(s.calls, key)
	s.lock.Unlock()
	c.wg.Done()
}

func main() {
	var wg sync.WaitGroup
	sl := NewSingleFlight()

	round := 10
	wg.Add(round)

	for i := 0; i < round; i++ {
		go func() {
			defer wg.Done()

			val, err := sl.Do("get_rand_int", func() (interface{}, error) {
				fmt.Println("i am doing")
				time.Sleep(time.Second)
				return rand.Int(rand.Reader, big.NewInt(100))
			})

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		}()
	}
	wg.Wait()
}
