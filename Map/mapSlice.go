package main

import (
	"fmt"
)

func main() {
	// 使用一个 map 来记录 monster 的信息 name 和 age, 也就是说一个 monster 对应一个 map,并
	// 且妖怪的个数可以动态的增加=>map 切片

	// 1. 声明一个切片
	var monster []map[string]string
	monster = make([]map[string]string, 2) // 准备放入两个妖怪
	// 2. 增加第一个妖怪的信息
	if monster[0] == nil {
		monster[0] = make(map[string]string, 2)
		monster[0]["name"] = "牛魔王"
		monster[0]["age"] = "500"
	}

	if monster[1] == nil {
		monster[1] = make(map[string]string, 2)
		monster[1]["name"] = "玉兔精"
		monster[1]["age"] = "400"
	}

	// 下面的写法越界
	// if monster[2] == nil {
	// 	monster[2] = make(map[string]string, 2)
	// 	monster[2]["name"] = "狐狸精"
	// 	monster[2]["age"] = "300"
	// }

	// 这里我们需要使用到切片的append函数，可以动态增加monster
	// 1. 先定义一个monser信息
	newMonster := map[string]string{
		"name": "新妖怪~火云邪神",
		"age":  "200",
	}
	monster = append(monster, newMonster)
	fmt.Println(monster)
}
