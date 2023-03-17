package main

import (
	"LearnGo/chatRoom/client/process"
	"fmt"
	"os"
)

// 定义两个变量，一个表示用户id，一个表示用户密码
var userId int
var userPwd string

func main() {
	// 接收用户的选择
	var key int
	// 判断是否还继续选择
	var loop = true

	for loop {
		fmt.Println("-----------欢迎登陆多人聊天系统------------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的id")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanln(&userPwd)
			// 完成登录
			// 创建一个UserProcess的实例
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
			// loop = false
		case 2:
			fmt.Println("注册用户")
			// loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入。。。")
		}
	}
	// 显示新的提示信息
	// if key == 1 {
	// 	// 说明用户要登陆
	// 	// fmt.Println("请输入用户的id")
	// 	// fmt.Scanln(&userId)
	// 	// fmt.Println("请输入用户的密码")
	// 	// fmt.Scanln(&userPwd)
	// 	// 先把登录的函数写到另外一个文件，比如login.go
	// 	login(userId, userPwd)
	// 	// if err != nil {
	// 	// 	fmt.Println("登录失败...")
	// 	// } else {
	// 	// 	fmt.Println("登录成功...")
	// 	// }

	// } else if key == 2 {
	// 	fmt.Println("进行用户注册的逻辑...")
	// }
}
