package main

import (
	"fmt"
)

type AInterface interface {
	Say()
}

type Stu struct {
	Name string 
}

type integer int 

func (i integer) Say() { //只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型。
	fmt.Println("integer Say()")
}


func (stu Stu) Say() {
	fmt.Println("Stu say()")
}

func main() {
	var stu Stu //结构体变量，实现了Say()，实现了AInterface
	
	// 接口本身不能创建实例,但是可以指向一个实现了该接口的自定义类型的变量(实例)
	var a AInterface = stu
	a.Say()

	var i integer
	var b AInterface = i
	b.Say()
}