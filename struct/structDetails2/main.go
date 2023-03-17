package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"` // `json:"name"`是struct tag
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

// 序列化
func main() {
	// 1. 创建一个Monster变量
	monster := Monster{"牛魔王", 500, "芭蕉扇"}

	// 2. 将monster变量序列化为json格式字串
	// json.Marshal函数中使用反射
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("jsonStr", string(jsonStr))
}
