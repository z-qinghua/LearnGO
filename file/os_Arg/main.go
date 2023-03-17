// @program:     LearnGo
// @file:        main.go.go
// @create:      2023-03-15 20:37
// @description:

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "Alice"
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], "")
		//参数会放置在切片 os.Args[] 中（以空格分隔），
		//从索引 1 开始（os.Args[0] 放的是程序本身的名字
	}
	fmt.Println("Good Morning", who)
}
