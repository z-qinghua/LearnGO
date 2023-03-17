package main

import (
	"LearnGo/func/utils"
	"fmt"
)

// 如果一个文件同时包含全局变量、init函数、main函数，则执行的流程 全局变量-->init函数-->main函数

var age int = test()

// 为了看到全局变量先被初始化，写个函数显示
func test() int {
	fmt.Println("test()")
	return 23
}

// init函数，通常可以在init()完成初始化工作
func init() {
	fmt.Println("init()")
}

func main() {
	fmt.Println("main(),age=", age)
	fmt.Printf("Age=%v Name=%v\n", utils.Age, utils.Name)
}
