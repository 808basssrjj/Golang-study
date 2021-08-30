package main

import (
	"fmt"
)

// 队列FIFO 先进先出
// 可以使用数组或链表模拟

// 非环形 数组实现
// maxSize front会随着数据输入=出改变  rear随着数据输入改变
type queue struct {
	arr     [5]int
	maxSize int
	front   int // 头指针
	rear    int // 尾指针
}

// addQueue
// 1.将尾指针往后移：  rear+1, front == rear[kong]
// 2.若尾指针rear小于等于maxSize-1,则数据存入，否则队列满了
func (q *queue) addQueue(val int) {
	if q.rear == q.maxSize-1 { //rear指向最后一个元素，并且包含
		fmt.Println("queue full")
		return
	}

	q.rear++
	q.arr[q.rear] = val
}

func (q *queue) getQueue() (val int) {
	if q.rear == q.front { //rear指向最后一个元素，并且包含
		fmt.Println("queue empty")
		return
	}

	q.front++
	val = q.arr[q.front]
	return
}

func (q *queue) showQueue() {
	// front不包含首个元素
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("array[%d]:%d\t", i, q.arr[i])
		fmt.Println()
	}
}

func main() {
	q := &queue{maxSize: 5, front: -1, rear: -1}
	q.addQueue(1)
	q.addQueue(2)
	q.addQueue(3)
	q.addQueue(4)
	q.addQueue(5)
	q.addQueue(5)
	q.showQueue()
	fmt.Println(q.getQueue())
	fmt.Println(q.getQueue())
	fmt.Println(q.getQueue())
}
