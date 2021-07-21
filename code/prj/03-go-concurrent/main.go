package main

import (
	// "fmt"
	// "runner/runner"
	// "time"
	"fmt"
	"io"
	"runner/pool"
	"sync"
	"sync/atomic"
	"time"
)

// 定义一个资源
type DBconnection struct {
	id int32
}

func (D DBconnection) Close() error {
	fmt.Println("database closed #" + fmt.Sprint(D.id))
	return nil
}

var counter int32

func Factory() (io.Closer, error) {
	atomic.AddInt32(&counter, 1)
	return &DBconnection{id: counter}, nil
}

var wg sync.WaitGroup

func performQuery(query int, pool *pool.Pool) {
	defer wg.Done()

	resource, err := pool.AcquireResource()
	if err != nil {
		fmt.Println("err:", err)
	}
	defer pool.ReleaseResource(resource)

	time.Sleep(time.Second)
	fmt.Println("finnish query" + fmt.Sprint(query))

}

func main() {
	p, err := pool.New(Factory, 5)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go performQuery(i, p)
	}
	wg.Wait()
	p.Close()
}

// func createTask() func(int) {
// 	return func(id int) {
// 		time.Sleep(time.Second)
// 		fmt.Printf("task complate #%d\n", id)
// 	}
// }
// func main() {
// 	// r := runner.New(4 * time.Second)
// 	// r.AddTask(createTask(), createTask(), createTask())

// 	// err := r.Start()
// 	// switch err {
// 	// case runner.ErrInterrupt:
// 	// 	fmt.Println("tasks interrupt")
// 	// case runner.ErrTimeout:
// 	// 	fmt.Println(" tasks timeout")
// 	// default:
// 	// 	fmt.Println("all tasks complete")
// 	// }

// }
