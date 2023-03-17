package main

import (
	"fmt"
)

func modify(map1 map[int]int) {
	map1[10] = 900
}

// 声明一个结构体
type Stu struct {
	name    string
	age     int
	address string
}

func main() {

	// map是引用类型，遵守引用类型传递的机制，在一个函数接收map,
	// 修改后，会直接修改原来的map

	map1 := make(map[int]int)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	modify(map1)
	// 看结果，map1[10] = 900，说明map是引用类型
	fmt.Println(map1)

	// map的value也经常使用struct类型
	// 更适合管理复杂的数据（比前面value是一个map更好），
	// 比如value为Student结构体
	// 1. map的key为学生的学号，是唯一的
	// 2. map的value为结构体，包含学生的名字，年龄，地址

	students := make(map[string]Stu, 10)
	stu1 := Stu{"tom", 18, "beijing"}
	stu2 := Stu{"marry", 28, "shanghai"}
	students["no1"] = stu1
	students["no2"] = stu2

	fmt.Println(students)

	// 遍历各个学生信息
	for k, v := range students {
		fmt.Printf("学生的学号是%v \n", k)
		fmt.Printf("学生的名字是%v \n", v.name)
		fmt.Printf("学生的年龄是%v \n", v.age)
		fmt.Printf("学生的地址是%v \n", v.address)
		fmt.Println() // 每个学生信息后面换行
	}

}
