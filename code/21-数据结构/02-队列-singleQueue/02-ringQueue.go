package main

import (
	"errors"
	"fmt"
	"os"
)

// 1,(tail+1) % maxSize =head 满了
// 2,tail ==head 为空
// 3,初始化 tail=0  head=0
// 4.（tail+maxSize-head）% maxSize  统计

// circleQueue 环形队列
type circleQueue struct {
	maxSize int
	arr     [4]int
	head    int //0
	tail    int //0
}

// 入队
func (c *circleQueue) push(val int) (err error) {
	if c.isFull() {
		err = errors.New("queue is full")
		return err
	}
	//队尾不包含元素
	c.arr[c.tail] = val
	//c.tail++
	c.tail = (c.tail + 1) % c.maxSize //尾部元素指向下一个空间位置,取模运算保证了索引不越界（余数一定小于除数）
	return nil
}

// 出队
func (c *circleQueue) pop() (res int, err error) {
	if c.isEmpty() {
		err = errors.New("queue is empty")
		return 0, err
	}
	//队尾不包含元素
	//res = c.arr[c.head]
	res = c.arr[c.head]
	//c.head++
	c.head = (c.head + 1) % c.maxSize
	return res, nil
}

// 是否满了
func (c *circleQueue) isFull() bool {
	return (c.tail+1)%c.maxSize == c.head
}

// 是否为空
func (c *circleQueue) isEmpty() bool {
	return c.tail == c.head
}

// 大小
func (c *circleQueue) size() int {
	return (c.tail + c.maxSize - c.head) % c.maxSize
}

// 显示队列
func (c *circleQueue) show() {
	size := c.size()
	if size == 0 {
		fmt.Println("queue is empty")
		return
	}

	tempHead := c.head
	for i := 0; i < size; i++ {
		fmt.Printf("queue[%d]=%v\t", tempHead, c.arr[tempHead])
		tempHead = (tempHead + 1) % c.maxSize
	}
}
func main() {
	var key string
	var val int
	q := &circleQueue{
		maxSize: 4,
		head:    0,
		tail:    0,
	}
	for {
		fmt.Println("------------------------------")
		fmt.Println("1.输入push表示添加数据到队列")
		fmt.Println("2.输入pop表示从队列中获取数据")
		fmt.Println("3.输入show表示显示队列")
		fmt.Println("4.输入exit表示退出")
		fmt.Println("------------------------------")
		fmt.Scanln(&key)
		switch key {
		case "push":
			fmt.Println("请输入要添加的值：")
			fmt.Scanln(&val)
			err := q.push(val)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("添加成功")
				fmt.Println("tail：", q.tail)
			}
		case "pop":
			val, err := q.pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("得到的值为：", val)
				fmt.Println("Front：", q.head)
			}

		case "show":
			q.show()
			fmt.Println()
		case "exit":
			os.Exit(0)
		}
	}
}
