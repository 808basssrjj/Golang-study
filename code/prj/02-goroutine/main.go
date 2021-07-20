package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func player(name string, ch chan int) {
	defer wg.Done()
	for {
		ball, ok := <-ch
		if !ok {
			fmt.Printf("channel is closed! %s win\n", name)
			return
		}

		n := rand.Intn(100)
		if n%10 == 0 {
			// 没接到球
			close(ch)
			fmt.Printf("%s misses th ball %s losee\n", name, name)
			return
		}
		ball++
		fmt.Printf("%s receives the ball %d\n", name, ball)
		ch <- ball //吧球传给对手
	}
}

func main() {
	wg.Add(2)
	ch := make(chan int)

	go player("张三", ch)
	go player("lisi", ch)
	ch <- 0
	wg.Wait()
}
