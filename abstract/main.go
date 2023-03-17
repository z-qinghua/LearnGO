package main

import (
	"fmt"
)

// 定义一个结构体Account
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// 存款
func (account *Account) Deposite(pwd string, money float64) {
	// 查询输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	// 看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}
	// 余额变动
	account.Balance += money
	fmt.Println("存款成功")
}

// 取款
func (account *Account) WithDraw(pwd string, money float64) {
	// 查询密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	// 看取款金额是否正确
	if money <= 0 || money > account.Balance {
		fmt.Println("你输入的金额不正确")
		return
	}
	// 取款
	account.Balance -= money
	fmt.Println("取款成功")
}

// 查询余额
func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	fmt.Printf("账户:%v 余额:%v \n", account.AccountNo, account.Balance)
}

func main() {
	account := Account{
		AccountNo: "zeng666",
		Pwd:       "888888",
		Balance:   100.0,
	}
	account.Query("888888")
	account.Deposite("888888", 10000)
	account.WithDraw("888888", 2000)
	account.Query("888888")
}
