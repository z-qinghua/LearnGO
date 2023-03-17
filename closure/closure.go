package main

import (
	"fmt"
	"strings"
)

// 累加器
// AddUpper 是一个函数，返回的数据类型是fun(int)int
func AddUpper() func(int) int {

	// 返回的是一个匿名函数，但这个匿名函数引用到函数外的n，
	// 因此这个匿名就和n形成一个整体，构成闭包
	var n int = 10 // n初始化一次，每调用一次，进行累计
	var str string = "hello"
	return func(x int) int {
		n = n + x
		str += string(36) // 36->'$'
		fmt.Println(str)  //hello$-->hello$$-->hello$$$
		return n
	}
}

// 请编写一个程序，具体要求如下
// 1) 编写一个函数 makeSuffix(suffix string) 可以接收一个文件后缀名(比如.jpg)，并返回一个闭包
// 2) 调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg) ,则返回 文件名.jpg , 如果已经有.jpg 后缀，则返回原文件名。
// 3) 要求使用闭包的方式完成
// 4) strings.HasSuffix , 该函数可以判断某个字符串是否有指定的后缀。
func makeSuffix(suffix string) func(name string) string {
	// 如果没有指定后缀，则加上后缀，否则返回原来的名字
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 不用闭包实现上述功能，但此种方法需要每次传后缀
func makeSuffix2(name string, suffix string) string {
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}
	return name
}

func main() {

	// 使用前面的代码
	f1 := AddUpper()
	fmt.Println(f1(1)) // 11
	fmt.Println(f1(2)) // 13
	fmt.Println(f1(3)) // 16

	f2 := makeSuffix(".jpg")
	fmt.Println("文件处理后", f2("winter"))   // winter.jpg
	fmt.Println("文件处理后", f2("bird.jpg")) // bird.jpg

	fmt.Println("文件处理后~", makeSuffix2("winter", ".jpg"))   //winter.jpg
	fmt.Println("文件处理后~", makeSuffix2("bird.jpg", ".jpg")) //bird.jpg
}
