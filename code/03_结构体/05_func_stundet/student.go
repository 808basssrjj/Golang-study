package main

import (
	"fmt"
	"os"
)

// 函数版
/*
学生有id、姓名
程序提供展示学生列表、添加学生、删除学生等功能
*/

type student struct {
	id   int64
	name string
}

// 构造函数
func newStundet(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

//变量声明
var (
	allStudent map[int64]*student
)

func showAllStud() {
	for k, v := range allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", k, v.name)
	}

}
func addStu() {
	var (
		id   int64
		name string
	)
	// 1.创建一个学生
	fmt.Printf("请输入学号:")
	fmt.Scanln(&id)
	fmt.Printf("请输入姓名:")
	fmt.Scanln(&name)
	newStu := newStundet(id, name)
	// 2.添加到allStudent
	allStudent[id] = newStu

}
func deleteStu() {
	var (
		deleteId int64
	)
	// 1.输入删除学生的学号
	fmt.Printf("请输入删除学生的学号:")
	fmt.Scanln(&deleteId)
	// 2.删除map中对应的信息
	delete(allStudent, deleteId)
}

func main() {
	allStudent = make(map[int64]*student, 20) //初始化(申请内存)
	for {
		// 1.打印菜单
		fmt.Println("欢迎使用学生管理系统")
		fmt.Println(`
		1. 查看所有学生
		2. 新增学生
		3. 删除学生
		4. 退出
	`)
		fmt.Printf("请输入你想执行的操作:")
		// 2.等待输入
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你输入的是%d操作\n", choice)
		// 3.执行对于函数
		switch choice {
		case 1:
			showAllStud()
		case 2:
			addStu()
		case 3:
			deleteStu()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("滚!")
		}
	}
}
