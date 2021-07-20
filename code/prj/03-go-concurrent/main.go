package main

import (
	"fmt"
	"runner/runner"
	"time"
)

func createTask() func(int) {
	return func(id int) {
		time.Sleep(time.Second)
		fmt.Printf("task complate #%d\n", id)
	}
}

func main() {
	r := runner.New(4 * time.Second)
	r.AddTask(createTask(), createTask(), createTask())

	err := r.Start()
	switch err {
	case runner.ErrInterrupt:
		fmt.Println("tasks interrupt")
	case runner.ErrTimeout:
		fmt.Println(" tasks timeout")
	default:
		fmt.Println("all tasks complete")
	}
}
