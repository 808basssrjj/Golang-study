package service

import (
	"go_code/customer/model"
)

// CustomerService 对客户进行操作
type CustomerService struct {
	customers   []model.Customer
	customerNum int // 客户数量as
}

func NewCustomerService() *CustomerService {
	customer := model.NewCustomer(1, 18, "张三", "男", "110", "110@qq.com")
	CustomerService := &CustomerService{
		customers:   nil,
		customerNum: 1,
	}
	CustomerService.customers = append(CustomerService.customers, customer)
	return CustomerService
}

// List 返回列表
func (c *CustomerService) List() []model.Customer {
	return c.customers
}

// Add 添加客户
func (c *CustomerService) Add(cus model.Customer) bool {
	c.customerNum++
	cus.Id = c.customerNum
	c.customers = append(c.customers, cus)
	return true
}

// findById 根据id返回切片下标
func (c *CustomerService) findById(id int) (index int) {
	index = -1
	for i, v := range c.customers {
		if id == v.Id {
			index = i
		}
	}
	return
}

// Delete 删除客户
func (c *CustomerService) Delete(id int) bool {
	index := c.findById(id)
	if index == -1 {
		// 客户不存在
		return false
	}
	c.customers = append(c.customers[:index], c.customers[index+1:]...)
	return true
}

// Update 修改客户
func (c *CustomerService) Update(id int, cus model.Customer) bool {
	index := c.findById(id)
	if index == -1 {
		// 客户不存在
		return false
	}
	cus.Id = c.customers[index].Id
	c.customers[index] = cus
	return true
}
