// @program:     LearnGo
// @file:        main.go
// @create:      2022-10-07 16:13
// @description:

// 类型断言，不仅可以判定已有类型，还可以判定自定义类型
package main

import "fmt"

func assert(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%d\n", v)
	case float64:
		fmt.Printf("%f\n", v)
	case byte, uint16, string:
		fmt.Printf("%T  %v\n", v, v) //%T：输出变量类型，%v：输出变量本来值
	}
}

func assert2(i interface{}) {
	switch i.(type) {
	case int:
		v := i.(int)
		fmt.Printf("%d\n", v)
	case float64:
		v := i.(float64)
		fmt.Printf("%f\n", v)
	case byte, uint16, string:
		fmt.Printf("%T  %v\n", i, i) //%T：输出变量类型，%v：输出变量本来值
	}
}

func main() {
	var i interface{}
	var a int
	var b float64
	var c byte
	//分别赋值给i
	i = a
	assert2(i)

	i = b
	assert2(i)

	i = c
	assert2(i)
}
