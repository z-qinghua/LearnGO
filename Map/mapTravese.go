package main

import (
	"fmt"
)

func main() {
	// 使用for-range遍历map
	// 第二种方式（推荐）
	cities := make(map[string]string)
	cities["no1"] = "beijing"
	cities["no2"] = "tianjin"
	cities["no3"] = "shanghai"

	for k, v := range cities {
		fmt.Printf("k=%v v=%v\n", k, v)
	}

	// 使用for-range遍历一个结构比较复杂的map
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "Tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "shanghai"

	studentMap["stu02"] = make(map[string]string)
	studentMap["stu02"]["name"] = "marry"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "beijing"

	for k1, v1 := range studentMap {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v v2=%v\n", k2, v2)
		}
	}
	fmt.Printf("studentMap length= %v\n", len(studentMap))
}
