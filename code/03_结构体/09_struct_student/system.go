package main

import "fmt"

type student struct {
	id   int64
	name string
}
type system struct {
	allStudent map[int64]student
}

// 列表
func (s system) showStu() {
	for _, stu := range s.allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", stu.id, stu.name)
	}
}

// 添加
func (s system) addStu() {
	var (
		stuId   int64
		stuName string
	)
	// 1.获取用户输入
	fmt.Printf("请输入学号:")
	fmt.Scanln(&stuId)
	fmt.Printf("请输入姓名:")
	fmt.Scanln(&stuName)
	// 2.添加到allStudent中
	newStu := student{
		id:   stuId,
		name: stuName,
	}
	s.allStudent[newStu.id] = newStu
}

// 修改
func (s system) editStu() {
	var stuId int64
	// 1.输入修改学生的学号
	fmt.Printf("请输入删除学生的学号:")
	fmt.Scanln(&stuId)
	// 2.展示学生信息
	stuObj, ok := s.allStudent[stuId]
	if !ok {
		fmt.Println("学生不存在")
		return
	}
	fmt.Printf("要修改的学生是:%s\n", stuObj.name)
	// 3.获取修改后的值
	var newName string
	fmt.Printf("请输入修改后的姓名:")
	fmt.Scanln(&newName)
	// 4.修改
	stuObj.name = newName
	s.allStudent[stuId] = stuObj
}

// 删除
func (s system) deleteStu() {
	var detId int64
	// 1.输入删除学生的学号
	fmt.Printf("请输入删除学生的学号:")
	fmt.Scanln(&detId)
	// 2.查询是否存在
	_, ok := s.allStudent[detId]
	if !ok {
		fmt.Println("学生不存在!")
		return
	}
	// 3.存在删除
	delete(s.allStudent, detId)
	fmt.Println("删除成功!")
}
