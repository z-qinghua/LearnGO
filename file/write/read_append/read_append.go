package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// /打开一个存在的文件，将原来的内容读出显示在终端，并且追加 5 句"hello,北京!"
	// 1. 打开文件 "E:/WorkSpace/goProject/src/LearnGo/file/write/abc.txt"
	filePath := "E:/WorkSpace/goProject/src/LearnGo/file/write/abc.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Open file err = %v\n", err)
		return
	}

	// 及时关闭file句柄
	defer file.Close()

	//先读取原来文件的内容，并显示在终端.
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //读取到文件末尾
			break
		}
		// 显示到终端
		fmt.Print(str)
	}

	// 准备写入5句"hello,北京!"
	str := "hello,北京！\r\n"

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
