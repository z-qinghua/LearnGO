package main

import (
	"LearnGo/mypkg/utils"
	"fmt"
)

func main() {
	// 包的三大作用：
	// 1. 区分相同名字的函数、变量等标识符
	// 2. 当程序文件很多时，可以很好的管理项目
	// 3. 控制“函数、变量”等访问范围，即作用域
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '+'
	result := utils.Cal(n1, n2, operator)
	fmt.Println(result)

	// 访问其它包中的变量
	fmt.Println("utils.go num=", utils.Num1)
}
