package utils

import "fmt"

var Age int
var Name string

// Age 和 Name全局变量，我们需要在main.go使用
// 但是我们需要初始化age 和 Name

// init完成初始化工作
func init() {
	fmt.Println("utils包的init")
	Age = 23
	Name = "zeng"
}
