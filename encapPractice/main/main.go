/*
	1) 创建程序,在 model 包中定义 Account 结构体：在 main 函数中体会 Golang 的封装性。
	2) Account 结构体要求具有字段：账号（长度在 6-10 之间）、余额(必须>20)、密码（必须是六
	3) 通过 SetXxx 的方法给 Account 的字段赋值。(同学们自己完成
	4) 在 main 函数中测试
*/
package main

import (
	"LearnGo/encapPractice/model"
	"fmt"
)

func main() {
	account := model.NewAccount("zeng6", "888888", 30000)
	if account != nil {
		fmt.Println("创建成功=", *account)
	} else {
		fmt.Println("创建失败")
	}
	// fmt.Println(*account)
}
