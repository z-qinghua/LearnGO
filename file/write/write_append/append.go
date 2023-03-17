package main

import (
	"fmt"
	"bufio"
	"os"
)
func main() {
	// 打开一个存在的文件，在原来的内容追加内容 'ABC! ENGLISH!
	// 1. 打开文件 "E:/WorkSpace/goProject/src/LearnGo/file/write/abc.txt"
	filePath := "E:/WorkSpace/goProject/src/LearnGo/file/write/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}

	// 及时关闭file句柄
	defer file.Close()

	// 准备写入10句"ABC! ENGLISH!
	str := "ABC,ENGLISH\r\n"

	// 写入时，使用带缓存的*writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	// 因为writer是带缓存，因此在调用writestring方法时，
	// 其实内容是先写到缓存的，所以需要调用Flush方法，
	// 将缓冲数据真正写入到文件中，否则文件没有数据！！
	writer.Flush()
}