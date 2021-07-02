package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 使用goroutine和channel实现一个计算int64随机数各位数和的程序。
/*
	1.开启一个goroutine循环生成int64类型的随机数，发送到jobChan
	2.开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	3.主goroutine从resultChan取出结果并打印到终端输出
*/

// job
type job struct {
	value int64
}

// result
type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func inJobChan(j chan<- *job) {
	defer wg.Done()
	// 生成随机数发送到jobChan
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		j <- newJob
		time.Sleep(time.Second)
	}
}
func sumInResult(j <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	// 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	for {
		job := <-j
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}
		resultChan <- newResult
	}
}
func main() {
	wg.Add(1)
	go inJobChan(jobChan)
	// 开启24个goroutine
	for i := 0; i < 24; i++ {
		go sumInResult(jobChan, resultChan)
	}

	// 主goroutine从resultChan取出结果并打印到终端输出
	for res := range resultChan {
		fmt.Printf("value:%d\n sum:%d\n", res.job.value, res.sum)
	}
	wg.Wait()
}
