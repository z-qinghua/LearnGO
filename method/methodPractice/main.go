package main

import (
	"fmt"
)

// 声明结构体
type Circle struct {
	radius float64
}

// 声明方法
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	/*
		请编写一个程序，要求如下：
		1) 声明一个结构体 Circle, 字段为 radius
		2) 声明一个方法 area 和 Circle 绑定，可以返回面积。
	*/
	var c Circle
	c.radius = 4.0
	res := c.area()
	fmt.Println("面积=", res)
}
