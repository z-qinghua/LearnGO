package main

import (
	"fmt"
)

func main() {
	// 第一种方式
	var a map[string]string
	// 在使用之前需要先make，make的作用就是给map分配
	a = make(map[string]string, 10)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no3"] = "武松"
	fmt.Println(a)

	// 第二种方式（推荐）
	cities := make(map[string]string)
	cities["no1"] = "beijing"
	cities["no2"] = "tianjin"
	cities["no3"] = "shanghai"
	fmt.Println(cities)

	// 第三种方式
	heroes := map[string]string{
		"hero1": "宋江",
		"hero2": "卢俊义",
		"hero3": "吴用",
	}
	fmt.Println(heroes)

	// 	演示一个 key-value 的 value 是 map 的案例
	// 比如：我们要存放 2 个学生信息, 每个学生有 name 和 sex 信息
	// 思路：map[string]map[string]string
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "Tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "shanghai"

	studentMap["stu02"] = make(map[string]string)
	studentMap["stu02"]["name"] = "marry"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "beijing"

	fmt.Println(studentMap)
	fmt.Println(studentMap["stu02"])
	fmt.Println(studentMap["stu02"]["address"])
}
