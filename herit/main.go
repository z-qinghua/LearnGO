// @program:     LearnGo
// @file:        main.go
// @create:      2022-10-08 08:51
// @description:

package main

import "fmt"

//继承
type Plane struct {
	Color string
}

func (Plane) fly() {
	fmt.Println("Plane fly()")
}

type Bird struct { //继承Plane的方法
	Plane
}

//重写
func (b Bird) fly() {
	b.Plane.fly() //中间Plane不能省略，否则会产生歧义（调用的自身的fly），陷入循环，调用栈会爆掉
	fmt.Println("Bird fly()")
}

func main() {
	b := Bird{}
	fmt.Println(b.Plane.Color)
	fmt.Println(b.Color) //省略Plane
	b.Plane.fly()
	b.fly() //调用的Bird自己重写的方法

}
