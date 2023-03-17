package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    // 将"C:/Users/13237/Desktop/golang_class.txt" 文件内容导入到 "E:/WorkSpace/goProject/src/LearnGo/file/write/abc.txt"
    // 1. 首先将"C:/Users/13237/Desktop/golang_class.txt" 文件内容读取到内存
    filePath1 := "E:/WorkSpace/goProject/src/LearnGo/file/write/abc.txt"
    filePath2 := "E:/WorkSpace/goProject/src/LearnGo/file/write/def.txt"

    content, err := ioutil.ReadFile(filePath1)
    if err != nil {
        // 读取文件有错误
        fmt.Printf("read file err =  %v\n", err)
        return
    }

    fmt.Printf("%s", string(content))
    // 2. 将读取到的文件内容导入到 "E:/WorkSpace/goProject/src/LearnGo/file/write/abc.txt"
    // 原有内容会被覆盖掉
    err = ioutil.WriteFile(filePath2, content, 0666)
    if err != nil {
        fmt.Printf("write file err = %v\n", err)
    }
}
