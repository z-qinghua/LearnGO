package main

import (
	"fmt"
	"time"
)

func squares(c chan int) {
	for i := 1; i <= 9; i++ {
		c <- i * i
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	close(c)//关闭通道
}

func main() {
	fmt.Println("main() started!")

	c := make(chan int)

	go squares(c) //started goroutine

	time.Sleep(time.Second)

	// for {
	// 	val, ok := <-c
	// 	if ok {
	// 		fmt.Println(val, ok)
	// 	}else {
	// 		fmt.Println(val, ok, "<-- loop broke")
	// 		// 因为是无限循环，所以我们需要在一个阶段使用 break 去中断循环
	// 		break
	// 	}
	// }

	for val := range c {
		fmt.Println(val)
	}

	fmt.Println("main() stopped!")
}

// 两个协程交替进行
// 当通道被关闭后我们在主线程中还可以读取 0 值数据。因为这个通道是用来传输 int 类型数据，默认情况下 int 的空值是 0 被返回。
// （注：从已经关闭的通道接收数据或者正在接收数据时，将会接收到通道类型的零值）
