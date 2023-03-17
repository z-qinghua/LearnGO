package main

import "fmt"

func main() {
	//Golang没有专门的字符类型，单个字符用byte
	//Golang用的是UTF-8编码，兼容ASCII编码，没有乱码困扰

	var c1 byte = 'c'
	var c2 byte = '0'

	//直接输出时，输出的是对应的UTF-8码值
	fmt.Println(c1, c2)

	//格式化输出对应字符
	fmt.Printf("c1=%c c2=%c\n", c1, c2)

	//var c3 byte = '北' //overflow溢出
	var c3 int = '北'
	fmt.Printf("c3=%c\n", c3)
	fmt.Printf("c3对应的码值:%d\n", c3)

	//给变量赋值一个数字，格式化输出%c,输出数值对应的字符
	var c4 int = 333
	fmt.Printf("c3对应的字符:%c\n", c4)

	//字符类型可以进行运算，相当于数值
	var n1 = 10 + 'a'
	fmt.Printf("n1=%d", n1)
}
