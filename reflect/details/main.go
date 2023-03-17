package main

import (
	"fmt"
	"reflect"
)

// 通过反射修改
// num int的值
// 修改student的值

func reflect01(b interface{}) {
	// 获取到reflect.Value
	rVal := reflect.ValueOf(b)

	// 看看rVal的Kind是
	fmt.Printf("rVal Kind=%v\n", rVal.Kind()) //指针类型
	// 3. rVal
	rVal.Elem().SetInt(20)

	// rVal.Elem()用于获取指针指向变量，类似如下
	// var num = 10
	// var c *int = &num
	// *c = 20
}

func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println("num = ", num)
}
