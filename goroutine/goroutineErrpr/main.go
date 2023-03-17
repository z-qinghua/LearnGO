// goroutine 中使用 recover，解决协程中出现 panic，导致程序崩溃问题

package main

import (
	"fmt"
	"time"
)

// function

func syHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

// 函数
func test() {
	// 这里使用defer + recover
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang"

	// 此处没有分配空间，会出现错误
}

func main() {
	go syHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok =", i)
		time.Sleep(time.Second)
	}
}
