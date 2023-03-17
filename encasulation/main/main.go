/*
	请大家看一个程序(person.go),不能随便查看人的年龄,工资等隐私，并对输入的年龄进行合理的验证。
	设计: model 包(person.go) main 包(main.go 调用 Person 结构体)
*/

package main

import (
	"LearnGo/encasulation/model"
	"fmt"
)

func main() {
	p := model.NewPerson("zeng")
	p.SetAge(23)
	p.SetSal(30000)
	fmt.Println(*p)
	fmt.Println("name=", p.Name, "age=", p.GetAge(), "sal=", p.GetSal())
}
