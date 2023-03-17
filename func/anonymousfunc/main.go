package main

import "fmt"

var (
	// Fun1是一个全局匿名函数
	Fun1 = func (n1 int, n2 int) int{
		return n1 * n2
	}
)

func main() {
	// 在定义匿名函数时就直接调用，这种方式匿名函数智能调用一次

	// 求两个数的和，用上述匿名函数方式完成
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println(res1)

	// 将匿名函数func(n1 int, n2 int) int赋给a变量
	// 则a 的函数类型就是函数类型，此时，我们就可以通过a完成调用
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(10, 20)

	fmt.Println(res2)

	// 全局匿名函数
	// 匿名函数赋给一个全局变量，在整个程序有效
	res3 := Fun1(10, 20)
	fmt.Println(res3)
}
