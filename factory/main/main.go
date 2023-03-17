package main

import (
	"LearnGo/factory/model" // 使用go mod 管理各种依赖，当导入自己定义的包的时候，需要从go.mod文件作为绝对路径
	"fmt"
)

func main() {
	// var stu = model.Student{
	// 	Name:  "tom",
	// 	Score: 78.8,
	// }
	// fmt.Println(stu)

	// 给定的结构体首字母小写，我们可以通过工厂模式解决
	var stu = model.NewStudent("tom~", 78.9) // 返回的是一个结构体指针
	fmt.Println(*stu)
	fmt.Println("Name=", (*stu).Name, "Score=", (*stu).GetScore())
	stu.SetScore(88.8)
	fmt.Println("Name=", (*stu).Name, "Score~=", (*stu).GetScore())

}
