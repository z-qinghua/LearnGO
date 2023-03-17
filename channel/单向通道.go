// @program:     LearnGo
// @file:        单向通道.go
// @create:      2022-10-24 20:13
// @description:

//我们已经学习到可以双向传递数据的通道，或者说，我们可以对通道做读操作和写操作。
//但是事实上我们也可以创建单向通道。比如只读通道只允许读操作，只写通道只允许写操作。
//
//单向通道也可以使用 make 函数创建，不过需要额外加一个箭头语法

package main

import "fmt"

//把双向通道转化为单向通道
func greets(roc <-chan string) {
	fmt.Println("Hello " + <-roc + "!")
	//从通道中读取数据，通道上的任何写入操作将会发生错误
}

func main() {
	//只读通道
	roc := make(<-chan int)
	//只写通道
	soc := make(chan<- int)

	c := make(chan string)
	go greets(c)
	c <- "John"

	fmt.Printf("Data type of roc is '%T'\n", roc)
	fmt.Printf("Data type of soc is '%T'\n", soc)

	fmt.Printf("main() stopped")
}

//使用单项通道可以提高程序的类型安全性，使得程序不容易出错
