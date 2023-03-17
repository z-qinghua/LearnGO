package main

import (
	"encoding/json" 
	"fmt"
)

// 定义一个结构体

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	// 演示
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}

	// 将monster序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化的结果
	fmt.Printf("monster序列化后=%v\n", string(data))
}

// 将Map进行序列化
func testMap() {
	// 定义一个Map
	var a map[string]interface{}

	// 使用map需要make
	a = make(map[string]interface{})

	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"

	// 将这个mao进行序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化的结果
	fmt.Printf("a map 序列化后=%v\n", string(data))
}

// 演示对切片进行序列化，我们这个切片[]map[string]interface{}
func testSlice() {
	var slice []map[string]interface{}

	var m1 map[string]interface{}
	// 使用mao前先make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 7
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	// 使用map前先make
	m2 = make(map[string]interface{})
	m2["make"] = "tom"
	m2["age"] = 20
	m2["address"] = [2]string{"墨西哥", "夏威夷"}

	slice = append(slice, m2)

	// 将切片进行序列化操作
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("slice 序列化后=%v\n", string(data))
}

// 对基本数据类型序列化，对基本数据类型进行序列化意义不大
func testFloat64() {
	var num1 float64 = 2345.67

	// 对num1进行序列化
	data, err := json.Marshal(&num1)

	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("num1 序列化后=%v\n", string(data))
}

func main() {
	testStruct()
	testMap()
	testSlice()   //演示对切片的序列化
	testFloat64() //演示对基本数据类型的序列化
}
