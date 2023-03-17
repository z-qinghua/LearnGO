// @program:     LearnGo
// @file:        main.go
// @create:      2022-10-09 10:57
// @description:

package main

import "fmt"

type Student struct {
	Name   string
	string //匿名字段，直接使用数据类型作为字段名
	int    //省略了变量，匿名字段中同一类型只能省略一次
}

func NewStudent(name, city string, age int) *Student {
	//var S Student
	//S.Name = name
	//S.string = city
	//S.int = age
	//return &S

	//另一种写法
	s := Student{Name: name, string: city, int: age}
	sPtr := &s
	fmt.Println(sPtr.Name)
	return &s

}
func main() {
	var s Student
	s.Name = "zeng"
	s.string = "hunan"
	s.int = 23
	//s := Student{Name: "zeng", string: "ld", int: 23}
	NewStudent("zeng", "HN", 23)
}
