package main

import "fmt"

// 获取用户终端输入
func main() {
	// fmt.scanln方式
	var name string
	var age int
	var sal float64
	var isPass bool

	// 方式一（推荐） fmt.Scanln()
	fmt.Println("输入姓名：")
	// 程序执行到fmt.Scanln(&name)会暂停，等待用户输入，并回车
	fmt.Scanln(&name)

	fmt.Println("输入年龄；")
	fmt.Scanln(&age)

	fmt.Println("输入薪水；")
	fmt.Scanln(&sal)

	fmt.Println("输入是否通过考试：")
	fmt.Scanln(&isPass)

	// 方式二 fmt.scanf(),可以使用指定的格式输入
	fmt.Println("输入你的姓名，年龄，薪水，是否通过考试，使用空格符隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)

	fmt.Printf("姓名是 %v\n 年龄是 %v\n 薪水是 %v\n 是否通过考试 %v\n", name, age, sal, isPass)

}
