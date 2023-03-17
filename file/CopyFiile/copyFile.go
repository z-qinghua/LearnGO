package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 自己编写一个函数，接收两个文件路径 srcFileName dstFileName
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
	}

	defer srcFile.Close()

	//通过 srcfile ,获取到 Reader
	reader := bufio.NewReader(srcFile)

	//打开 dstFileName,不存在时创建
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Open file err = %v\n", err)
		return
	}

	defer dstFile.Close()

	//通过 dstFile, 获取到 Writer
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)
}

func main() {
	// 将"E:/pictures/壁纸/small164756fSeTQ1665478076.jpg"文件拷贝到"E:/WorkSpace/goProject/src/LearnGo/file/write/picture.jpg"
	//调用 CopyFile 完成文件拷贝
	srcFile := "E:/pictures/壁纸/small164756fSeTQ1665478076.jpg"
	dstFile := "E:/WorkSpace/goProject/src/LearnGo/file/write/picture.jpg"

	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Println("拷贝成功")
	} else {
		fmt.Printf("拷贝失败,err=%v\n", err)
	}
}
