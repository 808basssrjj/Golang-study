package models

import (
	"bubble/dao"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// 创建
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	if err != nil {
		return err
	}
	return
}

// 获取全部
func GetAllTodo() (todos []*Todo, err error) {
	if err := dao.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return
}

// 获取某条
func GetATode(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err := dao.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

// 更新某条
func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

// 删除某条
func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
