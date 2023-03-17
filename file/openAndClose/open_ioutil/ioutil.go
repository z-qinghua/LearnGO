package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 使用ioutil.ReadFile一次性将文件读取到位
	file := "C:/Users/13237/Desktop/golang_class.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err = %v\n", err)
	}
	// 把读取的内容显示到终端
	// fmt.Printf("%v\n", content)
	fmt.Printf("%v\n", string(content)) //[]byte  //此处的换行符会在最后一行进行换行

	fmt.Println("读取结束")
	// 没有显示的open文件，因此也不需要close文件
	// 因为文件的open和close都被封装到ReadFile内部
}
