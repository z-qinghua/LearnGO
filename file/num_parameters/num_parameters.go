package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 定义一个结构体，用于保存统计结果
type CharCount struct {
	ChCount    int //英文字符个数
	NumCount   int //数字个数
	SpaceCount int //空格个数
	OtherCount int //其它字符个数
}

func main() {
	// 思路: 打开一个文件, 创一个 Reader
	//每读取一行，就去统计该行有多少个 英文、数字、空格和其他字符
	//然后将结果保存到一个结构体

	// 声明结构体变量
	var count CharCount

	fileName := "E:/WorkSpace/goProject/src/LearnGo/file/write/abc.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Open file err = %v\n", err)
		return
	}
	// 及时关闭文件
	defer file.Close()

	reader := bufio.NewReader(file)

	// 循环读取
	for {
		str, err := reader.ReadString('\n')
		for _, val := range str {
			switch {
			case val >= 'a' && val <= 'z':
				fallthrough //穿透，强制执行下一case,并忽略下一case的匹配，并执行语句
			case val >= 'A' && val <= 'Z':
				count.ChCount++
			case val >= '0' && val <= '9':
				count.NumCount++
			case val == ' ' || val == '\t':
				count.SpaceCount++
			default:
				count.OtherCount++
			}
		}
		if err == io.EOF { //读到文件末尾退出
			break
		}
	}
	fmt.Printf("英文字符=%v 数字字符=%v 空格字符=%v 其他字符=%v\n",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
