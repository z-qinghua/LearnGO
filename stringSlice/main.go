package main

import (
	"fmt"
)

func main() {
	// string就是一个byte数组，因此string也可进行切片处理
	str := "hello@atguigu"
	// 使用切片获取 atguigu
	slice := str[6:]
	fmt.Println("slice=", slice)

	// string是不可变的，也就说不能通过str[0] = 'z'方式来修改字符串
	// str[0] = 'z'【编译不会通过，报错，原因是string不可变】

	// 如果需要修改字符串，可以先将string --> []byte / 或者[]rune --> 修改 --> 重写为string
	// "hello@atguigu" 改写 "zello@atguigu"
	// arr1 := []byte(str)
	// arr1[0] = 'z'
	// str = string(arr1) // 重写
	// fmt.Println("str=", str)

	// 细节，我们转成[]byte后，可以处理英文和数字，但是不能处理中文
	// 原因是 []byte按字节处理，而一个汉字是3个字节，因此会出现乱码
	// 解决方法：将string转成[]rune即可，因为[]rune按字符处理，兼容汉字
	arr1 := []rune(str)
	arr1[0] = '北'
	str = string(arr1) // 重写
	fmt.Println("str=", str)
}
