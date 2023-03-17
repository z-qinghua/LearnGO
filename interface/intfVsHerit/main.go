package main

import (
	"fmt"
)

// Monkey结构体
type Monkey struct {
	Name string
}

// 声明接口
type BirdAble interface {
	Flying()
}
type FishAble interface {
	Swimming()
}

func (th *Monkey) Climbing() {
	fmt.Println(th.Name, "生来会爬树...")
}

// LittleMonkey结构体
type LittleMonkey struct {
	Monkey //继承
}

// 让LittleMonkey实现BirdAble
func (th *LittleMonkey) Flying() {
	fmt.Println(th.Name, "通过学习，会飞翔...")
}

func (th *LittleMonkey) Swimming() {
	fmt.Println(th.Name, "通过学习，会游泳...")
}

func main() {
// 创建一个littleMonkey实例
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}

	monkey.Climbing()
	monkey.Flying()
	monkey.Swimming()

	var a BirdAble = &monkey
	a.Flying()
}

/*
	对上面代码的小结
	1) 当 A 结构体继承了 B 结构体，那么 A 结构就自动的继承了 B 结构体的字段和方法，并且可以直
		接使用
	2) 当 A 结构体需要扩展功能，同时不希望去破坏继承关系，则可以去实现某个接口即可，因此我
		们可以认为：实现接口是对继承机制的补充.  实现接口可以看作是对 继承的一种补充
*/