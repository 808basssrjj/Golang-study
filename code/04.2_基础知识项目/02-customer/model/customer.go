package model

import "fmt"

// Customer 表示一个客户信息
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// NewCustomer 返回Customer实例
func NewCustomer(id, age int, name, gender, phone, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func NewCustomer1(age int, name, gender, phone, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}
// GetInfo 获取用户信息
func (c Customer) GetInfo() (info string) {
	info = fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", c.Id, c.Name, c.Gender, c.Age, c.Phone, c.Email)
	return
}
