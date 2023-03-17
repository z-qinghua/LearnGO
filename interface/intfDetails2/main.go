package main

import (
	"fmt"
	_ "fmt"
)

type BInterface interface {
	test01()
}

type CInterface interface {
	test02()
}

type AInterface interface {
	BInterface
	CInterface
	test03()
}

type usb interface {
	test04()
}

// 一个接口(比如 A 接口)可以继承多个别的接口(比如 B,C 接口)，这时如果要实现 A 接口，也必
// 须将 B,C 接口的方法也全部实现。

type Stu struct {
}

func (stu Stu) test01() {

}

func (stu Stu) test02() {

}

func (stu Stu) test03() {

}

func (stu *Stu) test04() {

}

// 空接口 interface{} 没有任何方法，所以所有类型都实现了空接口, 即我们可以把任何一个变量
// 赋给空接口
type T interface {
}

func main() {
	var stu Stu
	var a AInterface = stu
	a.test01()

	var t T = stu
	fmt.Println(t)

	// 我们可以把任何一个变量赋给空接口。
	var t2 interface{} = stu
	var num float64 = 8.8
	t = num
	t2 = num

	fmt.Println(t)
	fmt.Println(t2)

	var c usb = &stu
	c.test04()
}
