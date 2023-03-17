package main

import (
	"fmt"
)

/*
	1. 编写一个 Student 结构体，包含 name、gender、age、id、score 字段，分别为 string、string、int、
	int、float64 类型。
	2. 结构体中声明一个 say 方法，返回 string 类型，方法返回信息中包含所有字段值。
	3. 在 main 方法中，创建 Student 结构体实例(变量)，并访问 say 方法，并将调用结果打印输出。
*/
type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (stu *Student) say() string {
	infoStr := fmt.Sprintf("student的信息 name=[%v], gender=[%v], age=[%v], id=[%v], score=[%v]",
		stu.name, stu.gender, stu.age, stu.id, stu.score)
	return infoStr
}

/*
	1) 一个景区根据游人的年龄收取不同价格的门票，比如年龄大于 18，收费 20 元，其它情况门票免
	费.
	2) 请编写 Visitor 结构体，根据年龄段决定能够购买的门票价格并输出
*/
type Visitor struct {
	Name string
	Age  int
}

func (visitor *Visitor) showPrice() {
	if visitor.Age >= 90 || visitor.Age <= 8 {
		fmt.Println("考虑到安全，你不适合游玩")
		return
	}
	if visitor.Age > 18 {
		fmt.Printf("游客的名字为%v,年龄为%v,收费20元\n", visitor.Name, visitor.Age)
	} else {
		fmt.Printf("游客的名字为%v,年龄为%v,免票\n", visitor.Name, visitor.Age)
	}
}

func main() {
	stu := Student{
		name:   "tome",
		gender: "male",
		age:    18,
		id:     1000,
		score:  99.98,
	}
	fmt.Println((&stu).say())

	var v Visitor
	for {
		fmt.Println("输入你的名字")
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("退出程序...")
			break
		}
		fmt.Println("输入你的年龄")
		fmt.Scanln(&v.Age)
		v.showPrice()
	}
}
