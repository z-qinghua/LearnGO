package main

import (
	"bufio"
	"fmt"
	"io"
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

	// 函数退出时及时关闭文件
	defer file.Close()

	// 创建一个*Reader，带缓冲的
	/*
		const (
		defaultBufSize = 4096 //默认的缓冲区为 4096
		)
	*/
	reder := bufio.NewReader(file)
	// 循环读取文件的内容
	for {
		str, err := reder.ReadString('\n')
		if err == io.EOF { //EOF表示文件结尾
			break
		}
		// 输出内容
		fmt.Print(str)
	}
	fmt.Println("文件读取结束...")
}
