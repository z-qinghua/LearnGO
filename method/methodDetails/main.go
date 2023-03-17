package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

// 给*Student实现方法string
func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v] ", (*stu).Name, (*stu).Age)
	return str
}

func main() {

	// 如果一个类型实现了 String()这个方法，那么 fmt.Println 默认会调用这个变量的 String()进行输
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	// 如果你实现了*Student类型的String方法，就会自动调用
	fmt.Println(&stu)
}
