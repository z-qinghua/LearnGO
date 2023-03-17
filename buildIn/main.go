package main

import "fmt"

func main() {
	// 介绍内置函数使用

	num1 := 100
	fmt.Printf("num1=%v num1 address = %v num1 type = %T\n", num1, &num1, num1)

	// 1. len：用来求长度，比如 string、array、slice、map、channel
	str := "hello,world!"
	fmt.Printf("str的长度=%v\n", len(str))

	// 2. new函数，用于值类型，如：int、float、struct
	num2 := new(int)
	fmt.Printf("num2自己的值=%v num2地址=%v num2类型=%T num2存储地址对应的值=%v\n", num2, &num2, num2, *num2) // 新分配的内存存储默认值
	*num2 = 100
	// 自己的地址、存储的地址无法指定，由系统指定
	fmt.Printf("num2自己的值=%v num2地址=%v num2类型=%T num2存储地址对应的值=%v\n", num2, &num2, num2, *num2)

	// 3. make，用于引用类型，如：channel、map、slice

}
