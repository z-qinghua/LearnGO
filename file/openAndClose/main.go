package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	// 概念说明：file的叫法
	// 1. file叫file对象
	// 2. file叫file指针
	// 3. file叫文件句柄
	file, err := os.Open("C:/Users/13237/Desktop/golang_class.txt")
	if err != nil {
		fmt.Printf("err=%v\n", err)
	}
	// 输出下文件，看文件是什么？其实file就是一个指针*File
	fmt.Printf("file=%v\n", file)
	
	// 关闭文件
	err = file.Close()
	if err != nil {
		fmt.Printf("err=%v\n", err)
	}
}
