package main

import (
	"fmt"
)

// 编写一个学生考试系统

type Student struct {
	Name  string
	Age   int
	Score int
}

// 将Pupil和Graduate共有的方法绑定到*student
func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v \n", stu.Name, (*stu).Age, (*&stu.Score))
}

func (stu *Student) SetScore(score int) {
	// 业务判断

	(*stu).Score = score
}

// 小学生
type Pupil struct {
	Student //嵌入Student匿名结构体
}

// Pupil特有的结构体方法保留
func (p *Pupil) testing() {
	fmt.Printf("小学生正在考试中...\n")
}

// 大学生
type Graduate struct {
	Student //嵌入Student匿名结构体
}

// Graduate特有的结构体方法保留
func (p *Graduate) testing() {
	fmt.Printf("大学生正在考试中...\n")
}

func main() {
	// 当我们对结构体嵌入了匿名结构体，使用方法会发送变化
	pupil := &Pupil{}
	pupil.Student.Name = "tom~"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()

	graduate := Graduate{}
	graduate.Student.Name = "mary~"
	graduate.Student.Age = 22
	graduate.testing()
	graduate.Student.SetScore(90)
	graduate.Student.ShowInfo()
}
