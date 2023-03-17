package main

import (
	"fmt"
)

type Person struct {
	Name string
}

// 给Person绑定一个方法
func (person Person) test() {
	fmt.Println("test() name=", person.Name)
}

// 1. 给 Person 结构体添加 speak 方法,输出 xxx 是一个好人
func (p Person) speak() {
	fmt.Println(p.Name, "是一个好人")
}

// 2. 给 Person 结构体添加 jisuan 方法,可以计算从 1+..+1000 的结果, 说明方法体内可以函数一样，进行各种运算
// 说明方法体内可以和函数一样进行各种运算
func (p Person) jisuan() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算的结果是=", res)
}

// 3. 给 Person 结构体 jisuan2 方法,该方法可以接收一个数 n，计算从 1+..+n 的结果
func (p Person) jisuan2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算的结果是=", res)
}

// 4. 给 Person 结构体添加 getSum 方法,可以计算两个数的和，并返回结果
func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	// Golang中的方法是作用在指定的数据类型上的（即：和指定的数据类型绑定），因此自定义类型，都可以有方法，而不仅仅是struct
	var person Person
	person.Name = "tom"
	person.test() // 调用方法

	// 1. test方法和Person类型绑定
	// 2. test方法只能通过Person类型的变量来调用，而不能直接调用，也不能使用其它类型变量来调用
	// 3. func(p Person) test(){}...p表示哪个Person变量调用，这个p就是它的副本，这点和函数传参非常相似

	person.speak()
	person.jisuan()
	person.jisuan2(9)
	res := person.getSum(10, 20)
	fmt.Println("res=", res)
	// fmt.Println(person.getSum(10, 20))
}
