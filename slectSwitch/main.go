package main

import (
	"fmt"
)

// 一个简单函数
func test(b byte) byte {
	return b + 1
}

func main() {
	// 1.定义一个变量接收符
	// 2.使用switch完成
	var key byte
	fmt.Println("输入一个字符:a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)

	switch test(key) { // switch后接一个表达式
	case 'a':
		fmt.Println("monday,猴子穿新衣")
	case 'b':
		fmt.Println("Tuesday,猴子当小二")
	default:
		fmt.Println("输入有误")
	}

	var n1 int = 5
	var n2 int = 10
	switch n1 {
	case n2, 10, 5: // case后面可以有多个表达式，如果是常量则要求不能重复
		fmt.Println("ok1")
	default:
		fmt.Println("没有匹配到")
	}

	// switch后面可以不带表达式，类似于if-else使用
	var age int = 23
	switch {
	case age == 23:
		fmt.Println("age=23")
	case age == 22:
		fmt.Println("age=22")
	default:
		fmt.Println("没有匹配到")
	}

	// case可以对范围进行判断
	var score int = 30
	switch {
	case score > 90:
		fmt.Println("A")
	case score >= 70 && score <= 90:
		fmt.Println("B")
	default:
		fmt.Println("没有匹配到")
	}

	// switch后也可以直接声明/定义一个变量，分号结束（不推荐）
	switch grade := 90; {
	case grade > 90:
		fmt.Println("A")
	case grade >= 70 && grade <= 90:
		fmt.Println("B")
	default:
		fmt.Println("没有匹配到")
	}

	// switch的穿透，fallthrough
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough // 默认只能穿透一层
	case 20:
		fmt.Println("ok2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到")
	}

	// Type Switch:switch语句还可以被用于type-switch来判断某个interface变量实际指向的变量类型
	
}
