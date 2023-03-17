package model

import (
	"fmt"
)

type account struct {
	accountNo string
	pwd       string
	balance   float64
}

// 工厂模式的构造函数
func NewAccount(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号长度不对...")
		return nil
	}

	if len(pwd) != 6 {
		fmt.Println("密码长度不对...")
		return nil
	}
	if balance < 20 {
		fmt.Println("余额数目不对...")
		return nil
	}

	return &account{
		accountNo: accountNo,
		pwd:       pwd,
		balance:   balance,
	}
}
