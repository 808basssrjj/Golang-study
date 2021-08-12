package main

import (
	"fmt"
	"go_code/customer/model"
	"go_code/customer/service"
)

type customerView struct {
	key             string //接受用户呼入
	loop            bool
	customerService *service.CustomerService
}

// list 客户列表
func (c *customerView) list() {
	// 获取客户信息
	customers := c.customerService.List()

	fmt.Println("\n----------客户信息列表----------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱\t")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("----------客户列表完成----------")
}

// add 获得用户输入完成添加客户
func (c *customerView) add() {
	var name, gender, phone, email string
	var age int
	fmt.Println("\n----------添加客户----------")
	fmt.Print("姓名:")
	_, _ = fmt.Scanln(&name)
	fmt.Print("年龄:")
	_, _ = fmt.Scanln(&age)
	fmt.Print("性别:")
	_, _ = fmt.Scanln(&gender)
	fmt.Print("电话:")
	_, _ = fmt.Scanln(&phone)
	fmt.Print("邮箱:")
	_, _ = fmt.Scanln(&email)

	customer := model.NewCustomer1(age, name, gender, phone, email)
	c.customerService.Add(customer)
	fmt.Println("----------添加成功----------")
}

// delete 删除客户
func (c customerView) delete() {
	fmt.Println("\n----------删除客户----------")
	fmt.Print("输入要删除客户的编号(-1退出):")
	id := -1
	_, _ = fmt.Scanln(&id)
	if id == -1 {
		return //放弃操作
	}

	fmt.Print("确认删除(y/n)")
	choice := ""
	for {
		_, _ = fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" {
			if c.customerService.Delete(id) {
				fmt.Println("----------删除成功----------")
			} else {
				fmt.Println("----------删除失败,id不存在----------")
			}
			break
		}
		if choice == "n" || choice == "N" {
			break
		}
		fmt.Print("请输入(y/n)")
	}
}

// update 删除客户
func (c customerView) update() {
	fmt.Println("\n----------修改客户----------")
	fmt.Print("输入要修改客户的编号(-1退出):")
	id := -1
	_, _ = fmt.Scanln(&id)
	if id == -1 {
		return //放弃操作
	}

	var name, gender, phone, email string
	var age int
	fmt.Print("姓名:")
	_, _ = fmt.Scanln(&name)
	fmt.Print("年龄:")
	_, _ = fmt.Scanln(&age)
	fmt.Print("性别:")
	_, _ = fmt.Scanln(&gender)
	fmt.Print("电话:")
	_, _ = fmt.Scanln(&phone)
	fmt.Print("邮箱:")
	_, _ = fmt.Scanln(&email)
	customer := model.NewCustomer1(age, name, gender, phone, email)
	if c.customerService.Update(id, customer) {
		fmt.Println("----------修改成功----------")
	}else {
		fmt.Println("----------修改失败,id不存在----------")
	}
}

// 主菜单
func (c *customerView) mainMenu() {
	for {
		fmt.Println("\n----------客户信息管理软件----------")
		fmt.Println("          1.添 加 客 户")
		fmt.Println("          2.修 改 客 户")
		fmt.Println("          3.删 除 客 户")
		fmt.Println("          4.客 户 列 表")
		fmt.Println("          5.退      出")
		fmt.Print("请选择(1-5)")

		_, _ = fmt.Scanln(&c.key)
		switch c.key {
		case "1":
			c.add()
			//fmt.Println("添 加 客 户")
		case "2":
			c.update()
			//fmt.Println("修 改 客 户")
		case "3":
			c.delete()
			//fmt.Println("删 除 客 户")
		case "4":
			c.list()
			//fmt.Println("客 户 列 表")
		case "5":
			c.loop = false
		default:
			fmt.Println("输入有误,请重新输入")
		}

		if !c.loop {
			break
		}
	}
	fmt.Println("退出成功!")
}
func main() {
	cv := &customerView{
		key:             "",
		loop:            true,
		customerService: service.NewCustomerService(),
	}
	cv.mainMenu()
}
