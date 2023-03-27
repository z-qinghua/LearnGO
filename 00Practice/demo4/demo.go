// @program:     LearnGo
// @file:        demo.go
// @create:      2023-03-19 20:07
// @description:

package demo4

import (
	"fmt"
)

//定义一个抽象的组件

type Component interface {
	Operate()
}

//实现一个具体的组件

type Component1 struct {
}

func (c1 *Component1) Operate() {
	fmt.Println("c1 operate")
}

//定义一个抽象的装饰者

type Decorator interface {
	Component
	Do()
}

//实现一个具体的装饰者

type Decorator1 struct {
	c Component
}

func (c1 *Decorator1) Do() {
	fmt.Println("c1 发生了装饰行为")
}

func (c1 *Decorator1) Operate() {
	c1.Do()
	c1.c.Operate()
}
