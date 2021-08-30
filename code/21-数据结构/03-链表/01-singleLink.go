package main

import "fmt"

type heroNode struct {
	no   int
	name string
	nick string
	next *heroNode //指向下一个节点
}

// 1.尾插法 新增
func insertHero(head *heroNode, newHero *heroNode) {
	// 1.找到最后一个节点
	temp := head
	for {
		if temp.next == nil { //找到了
			break
		}
		temp = temp.next
	}
	// 2.加到最后
	temp.next = newHero
}

// 2.显示链表所有节点
func listLink(head *heroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("link is empty")
		return
	}
	for {
		fmt.Printf("[%d, %s, %s]=> ", temp.next.no, temp.next.name, temp.next.nick)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
	fmt.Println()
}

// 3.有序插入  根据no
func insertHero2(head *heroNode, newHero *heroNode) {
	// 1.找到最后一个节点
	temp := head
	for {
		if temp.next == nil {
			// 到最后
			break
		}
		if temp.next.no > newHero.no {
			// 可以插入到temp后面
			break
		} else if temp.next.no == newHero.no {
			fmt.Println("当前序号已存在")
			return
		}
		temp = temp.next
	}
	newHero.next = temp.next //先将新元素指向后面的元素
	temp.next = newHero
}

// 4.删除节点
func deleteNode(head *heroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			// 到最后
			break
		}
		if temp.next.no == id {
			// 找到要删除的节点
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("当前id不存在")
	}
}

func main() {
	// 1.创建头节点
	head := &heroNode{}
	// 2.创建一个新的heroNode
	hero1 := &heroNode{name: "宋江", nick: "及时雨", no: 1}
	hero2 := &heroNode{name: "卢俊义", nick: "玉麒麟", no: 2}
	hero3 := &heroNode{name: "吴用", nick: "智多星", no: 3}
	//hero4 := &heroNode{name: "林冲", nick: "豹子头", no: 3}

	//insertHero(head, hero2)
	//insertHero(head, hero1)
	//insertHero(head, hero3)

	insertHero2(head, hero2)
	insertHero2(head, hero1)
	insertHero2(head, hero3)
	//insertHero2(head, hero4)

	listLink(head)
	deleteNode(head, 2)
	deleteNode(head, 5)
	listLink(head)
}
