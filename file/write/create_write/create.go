package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 创建一个新文件，写入内容 5句"hello,world"
	// 1. 打开文件 "E:/WorkSpace/goProject/src/LearnGo/file/write/abc.txt"
	filePath := "E:/WorkSpace/goProject/src/LearnGo/file/write/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("file err = %v\n", err)
		return
	}

	// 及时关闭file句柄
	defer file.Close()

	// 准备写入5句"hello,world"
	str := "hello,world\n"

	// 写入时，使用带缓存的*writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	// 因为writer是带缓存，因此在调用writestring方法时，
	// 其实内容是先写到缓存的，所以需要调用Flush方法，
	// 将缓冲数据真正写入到文件中，否则文件没有数据！！
	writer.Flush()
}
