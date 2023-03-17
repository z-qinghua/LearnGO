package service

import (
	"LearnGo/0projectExcercise/customerManage/model"
)

// 完成对Customer的操作，包括增删改查

// 声明一个结构体
type CustomerService struct {
	customers []model.Customer
	// 声明一个字段，表示当前切片含有多少个客户
	// 该字段后面作为新客户的id+1
	customerNum int
}

// 编写一个方法，返回*CustomerService
func NewCustomerService() *CustomerService {
	customerservice := &CustomerService{}
	customerservice.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@163.com")
	customerservice.customers = append(customerservice.customers, customer)
	return customerservice
}

// 添加客户到customer切片
func (th *CustomerService) Add(customer model.Customer) bool {
	// 确定一个分配id的规则，就是添加顺序
	th.customerNum++
	customer.Id = th.customerNum
	th.customers = append(th.customers, customer)
	return true
}

// 修改客户切片中的信息
func (th *CustomerService) UpDate(id int, customer model.Customer) bool{
	index := th.FindById(id)
	// 如果index = -1，则没有该客户
	if index == -1 {
		return false
	}
	// 修改用户信息
	customer.Id = id
	th.customers[index] = customer
	return true
}

// 删除客户
func (th *CustomerService) Delete(id int) bool {
	index := th.FindById(id)
	// 如果index = -1，则没有该客户
	if index == -1 {
		return false
	}
	// 从切片中删除该元素
	th.customers = append(th.customers[:index], th.customers[index+1:]...)
	return true
}

// 根据id查找客户在切片中对应的下标，如果没有该客户则返回-1
func (th *CustomerService) FindById(id int) int {
	index := -1
	// 遍历th.Customers切片
	for i := 0; i < len(th.customers); i++ {
		if th.customers[i].Id == id {
			// 找到
			index = i
		}
	}
	return index
}

// 返回客户的切片
func (th *CustomerService) List() []model.Customer {
	return th.customers
}
