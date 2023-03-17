package main

import (
	"LearnGo/0projectExcercise/customerManage/model"
	"LearnGo/0projectExcercise/customerManage/service"
	"fmt"
)

type CustomerView struct {
	Key  string //接收用户输入
	loop bool   //表示是否循环显示主菜单
	// 增加一个字段customerService
	customerservice *service.CustomerService
}

func (th *CustomerView) MainMenu() {
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("                  1 添 加 客 户")
		fmt.Println("                  2 修 改 客 户")
		fmt.Println("                  3 删 除 客 户")
		fmt.Println("                  4 客 户 列 表")
		fmt.Println("                  5 退 出")
		fmt.Print("输入你的选项(1-5):")
		fmt.Scanln(&th.Key)

		switch th.Key {
		case "1":
			th.add()
		case "2":
			th.upDate()
		case "3":
			th.Delete()
		case "4":
			th.List()
		case "5":
			th.exit()
		default:
			fmt.Println("输入有误，重新输入...")
		}
		if !th.loop {
			fmt.Println("退出系统...")
			break
		}
	}
}

// 显示所有客户信息
func (th *CustomerView) List() {
	// 获取当前所有客户的信息（在切片中）
	customers := th.customerservice.List()

	// 显示
	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t 姓名\t 性别\t 年龄\t 电话\t 邮箱")
	for i := 0; i < len(customers); i++ {
		//fmt.Println(customers[i].Id,"\t", customers[i].Name...)
		fmt.Println(customers[i].GetInfo())
	}

	fmt.Printf("\n-------------------------客户列表显示完成-----------------------\n\n")
}

// 得到用户的输入信息，构建新的客户，并完成添加
func (th *CustomerView) add() {
	fmt.Println("---------------------添加客户---------------------")
	fmt.Print("姓名:")
	name := ""
	fmt.Scanln(&name)

	fmt.Print("性别:")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Print("年龄:")
	age := 0
	fmt.Scanln(&age)

	fmt.Print("电话:")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Print("邮箱:")
	email := ""
	fmt.Scanln(&email)

	// 构建一个新的Customer实例
	// 注意id号，没有让用户输入，id是唯一的，需要系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)

	// 调用
	if th.customerservice.Add(customer) {
		fmt.Println("---------------------添加完成---------------------\n\n")
	} else {
		fmt.Println("---------------------添加失败---------------------\n\n")
	}
}

// 得到用户Id，修改Id对应客户信息
func (th *CustomerView) upDate() {
	fmt.Println("---------------------修改客户信息---------------------")
	fmt.Print("请选择修改的客户编号(-1退出)")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("放弃修改操作")
		return
	}

	// 修改信息
	fmt.Print("姓名:")
	name := ""
	fmt.Scanln(&name)

	fmt.Print("性别:")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Print("年龄:")
	age := 0
	fmt.Scanln(&age)

	fmt.Print("电话:")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Print("邮箱:")
	email := ""
	fmt.Scanln(&email)

	// 构建一个新的Customer实例
	// 注意id号，id号唯一标识客户，不能修改
	customer := model.NewCustomer2(name, gender, age, phone, email)

	// 调用
	for {
		fmt.Print("确认是否修改(y/n):")
		choice := ""
		fmt.Scanln(&choice)

		if choice == "y" || choice == "Y" {
			// 调用customerService的UpDate方法
			if th.customerservice.UpDate(id, customer) {
				fmt.Println("---------------------修改完成---------------------\n\n")
				break
			} else {
				fmt.Println("---------------------修改失败,Id号不存在---------------------\n\n")
				break
			}
		} else if choice == "n" || choice == "N" {
			fmt.Println("---------------------放弃修改---------------------\n\n")
			break
		}
	}
}

// 得到用户的输入Id，删除Id对应的客户
func (th *CustomerView) Delete() {
	fmt.Println("---------------------删除客户---------------------")
	fmt.Print("请选择删除的客户编号(-1退出)")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("放弃删除操作")
		return
	}

	// 循环判断，直到输入y/n才退出
	for {
		fmt.Print("确认是否删除(y/n)")
		choice := ""
		fmt.Scanln(&choice)

		if choice == "y" || choice == "Y" {
			// 调用customerService的Delete方法
			if th.customerservice.Delete(id) {
				fmt.Println("---------------------删除成功---------------------\n\n")
				break
			} else {
				fmt.Println("---------------------删除失败,id号不存在------------\n\n")
				break
			}
		} else if choice == "n" || choice == "N" {
			fmt.Println("---------------------放弃删除---------------------\n\n")
			break
		}
	}
}

// 退出软件
func (th *CustomerView) exit() {
	for {
		fmt.Print("确认是否退出(y/n):")
		fmt.Scanln(&th.Key)

		if th.Key == "y" || th.Key == "Y" {
			th.loop = false
			break
		}

		if th.Key == "n" || th.Key == "N" {
			break
		}
	}
}

func main() {
	//在 main 函数中，创建一个 customerView,并运行显示主菜单
	customerView := CustomerView{
		Key:  "",
		loop: true,
	}
	//这里完成对 customerView 结构体的 customerService 字段的初始化
	customerView.customerservice = service.NewCustomerService()
	//显示主菜单..
	customerView.MainMenu()
}
