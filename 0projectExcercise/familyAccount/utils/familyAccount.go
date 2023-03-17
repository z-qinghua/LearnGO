package utils

import (
	"fmt"
)

// 声明结构体
type FamilyAccount struct {
	// 声明一个字段保存用户输入的选项
	choose string
	// 声明一个字段控制是否退出循环
	loop bool
	// 定义账户余额
	balance float64
	// 每次收支金额
	money float64
	// 每次收支的说明
	notes string
	// 定义一个字段记录是否有收支的行为
	flag bool
	// 收支的详情使用字符串来记录
	// 当有收支时，只需要对details进行拼接处理
	details string
}

// 编写工厂模式的构造方法，返回一个*FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		choose:  "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		notes:   "",
		flag:    false,
		details: "收支\t账户余额\t收支金额\t说明",
	}
}

// 将显示明细编写成一个方法
func (account *FamilyAccount) showdetails() {
	fmt.Println("----------------当前收支明细记录----------------")
	if !account.flag {
		fmt.Println("没有任何收支，来一笔吧")
	} else {
		fmt.Println(account.details)
	}
}

// 将登记收入写成一个方法，和*FamilyAccount绑定
func (account *FamilyAccount) income() {
	fmt.Print("本次收入金额：")
	fmt.Scanln(&account.money)
	account.balance += account.money //修改账户余额
	fmt.Print("本次收入说明：")
	fmt.Scanln(&account.notes)
	// 将收支情况拼接到details
	account.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", account.balance, account.money, account.notes)
	account.flag = true
}

// 将登记支出写成一个方法，和*FamilyAccount绑定
func (account *FamilyAccount) pay() {
	fmt.Print("本次支出金额：")
	fmt.Scanln(&account.money)
	if account.money > account.balance {
		fmt.Println("余额不足")
	}
	account.balance -= account.money //修改账户余额
	fmt.Print("本次支出说明：")
	fmt.Scanln(&account.notes)
	// 将收支情况拼接到details
	account.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", account.balance, account.money, account.notes)
	account.flag = true
}

// 将退出系统写成一个方法，和*FamilyAccount绑定
func (account *FamilyAccount) exit() {
	fmt.Print("你确定要退出吗? y/n:")
	key := ""
	for {
		fmt.Scanln(&key)
		if key == "y" || key == "n" {
			break
		}
		fmt.Print("你的输入有误，请重新输入：")
	}
	if key == "y" {
		account.loop = false
	}
}

// 显示主菜单
func (account *FamilyAccount) Menu() {
	for {
		fmt.Println("\n----------------家庭收支记账软件----------------")
		fmt.Println("                 1.收支明细")
		fmt.Println("                 2.登记收入")
		fmt.Println("                 3.登记支出")
		fmt.Println("                 4.退出软件")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&account.choose)
		switch account.choose {
		case "1":
			account.showdetails()
		case "2":
			account.income()
		case "3":
			account.pay()
		case "4":
			account.exit()
		default:
			fmt.Println("请输入正确选项...")
		}
		if !account.loop {
			fmt.Println("退出软件...")
			break
		}
	}
}
