package main

import (
	"fmt"
	"os"
)

// 结构体版
/*
学生有id、姓名
程序提供展示学生列表、添加学生、修改学生、删除学生等功能
*/

var sys system // 全局变量
func showMeun() {
	// 1.打印菜单
	fmt.Println("-----------------欢迎使用学生管理系统----------------------")
	fmt.Println(`
		1. 查看所有学生
		2. 新增学生
		3. 修改学生
		4. 删除学生
		5. 退出
	`)
}
func main() {
	sys = system{
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMeun()
		// 等待用户输入
		fmt.Printf("请输入要执行的操作:")
		var chioce int
		fmt.Scanln(&chioce)
		fmt.Println("你输入的是:", chioce)

		switch chioce {
		case 1:
			sys.showStu()
		case 2:
			sys.addStu()
		case 3:
			sys.editStu()
		case 4:
			sys.deleteStu()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("滚!")
		}
	}
}
