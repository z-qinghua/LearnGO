// @program:     LearnGo
// @file:        closeChannel.go
// @create:      2022-10-24 15:57
// @description:

//关闭通道
package main

import "fmt"

func greet(c chan string) {
	<-c //for john
	<-c //for mike
	//第一个<-c是非阻塞的，因为此时有数据可读，
	//第二个<-将被阻塞，因为通道没有数据可读
}

func main() {
	fmt.Println("main() started")
	c := make(chan string, 1)

	go greet(c)
	c <- "john" //将阻塞协程直到有其他协程从此通道中读取数据，
	// greet被调度器执行

	//main协程将被激活并且程序执行close(c)关闭通道
	close(c) //closing channel

	//for {
	//	val, ok := <-c
	//	if ok == false {
	//		fmt.Printf("[%v],%v\n", val, ok)
	//		//从已经关闭的通道接收数据或者正在接收数据数据时，会收到通道类型的零值
	//
	//		break
	//	} else {
	//		fmt.Println(val, ok)
	//	}
	//}

	//for val := range c { //通道关闭会主动退出循环
	//	fmt.Println(val)
	//}

	c <- "mike" //从一个已关闭的通道读取数据，程序会异常
	fmt.Println("main() stopped")
}
