package main

import (
	"fmt"
)

func main() {
	// 第二种方式（推荐）
	cities := make(map[string]string)
	cities["no1"] = "beijing"
	cities["no2"] = "tianjin"
	cities["no3"] = "shanghai"
	fmt.Println(cities)

	// 如果no3这个key已经存在，就是修改，如果没有，就是增加
	cities["no3"] = "shanghai~"
	fmt.Println(cities)

	// 删除操作
	delete(cities, "no1")
	fmt.Println(cities)

	// 当delete指定的key不存在时，删除不会操作，也不会报错
	delete(cities, "no4")
	fmt.Println(cities)

	fmt.Println("-------------------------")
	// 如果希望一次性删除所有的key
	// 1. 遍历所有的key，如何逐一删除[遍历]
	// 2. 直接make一个新的空间
	cities = make(map[string]string, 10)
	fmt.Println(cities)
}