package main

import (
	"fmt"
	"sync"
)

var (
	putNum   = make(chan int, 80000)
	primeNum = make(chan int, 10000)
	wg       sync.WaitGroup
)

// 放入数据
func insert(intch chan int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		intch <- i
	}
	close(intch)
}

func prime(intch chan int, priCh chan int) {
	defer wg.Done()

	var flag bool
	for num := range intch {
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 { //不是素数
				flag = false
				break
			}
		}
		if flag {
			priCh <- num
		}
	}
}
func main() {
	// 统计1-80000的数字中,哪些是素数 prime
	wg.Add(1)
	go insert(putNum)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go prime(putNum, primeNum)
	}
	wg.Wait()
	fmt.Println(len(primeNum))

	close(primeNum)

	for v := range primeNum {
		fmt.Printf("素数%v\n", v)
	}

}
