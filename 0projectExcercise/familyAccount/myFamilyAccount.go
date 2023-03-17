package main

import (
	"fmt"
)

func main() {
	// 声明一个变量保存用户输入的选项
	var choose string
	// 声明一个变量控制是否退出循环
	var loop bool = true

	// 定义账户余额
	balance := 10000.0
	// 每次收支金额
	money := 0.0
	// 每次收支的说明
	notes := ""
	// 定义一个变量记录是否有收支的行为
	flag := false
	// 收支的详情使用字符串来记录
	// 当有收支时，只需要对details进行拼接处理
	details := "收支\t账户余额\t收支金额\t说明"

	// 主菜单的显示
	for {
		fmt.Println("\n----------------家庭收支记账软件----------------")
		fmt.Println("                 1.收支明细")
		fmt.Println("                 2.登记收入")
		fmt.Println("                 3.登记支出")
		fmt.Println("                 4.退出软件")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&choose)
		switch choose {
		case "1":
			fmt.Println("----------------当前收支明细记录----------------")
			if !flag {
				fmt.Println("没有任何收支，来一笔吧")
			} else {
				fmt.Println(details)
			}
		case "2":
			fmt.Print("本次收入金额：")
			fmt.Scanln(&money)
			balance += money //修改账户余额
			fmt.Print("本次收入说明：")
			fmt.Scanln(&notes)
			// 将收支情况拼接到details
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, notes)
			flag = true
		case "3":
			fmt.Print("本次支出金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			balance -= money //修改账户余额
			fmt.Print("本次支出说明：")
			fmt.Scanln(&notes)
			// 将收支情况拼接到details
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, notes)
			flag = true
		case "4":
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
				loop = false
			}

		default:
			fmt.Println("请输入正确选项...")
		}
		if !loop {
			fmt.Println("退出软件...")
			break
		}
	}
}
