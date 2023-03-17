package main


import (
	"fmt"
)

// 使用内置函数 close 可以关闭 channel, 当 channel 关闭后，就不能再向 channel 写数据了，但是仍然
// 可以从该 channel 读取数据

func main()  {
	intChan := make(chan int, 3)

	intChan <- 100
	intChan <- 200

	// 关闭管道,然后就不能再写入数据到channel
	close(intChan)

	fmt.Println("okok~")

	// 管道关闭后，读取数据是可以的
	n1 := <- intChan
	fmt.Println("n1 =", n1)

	intChan2 := make(chan int, 100)

	for i := 0; i < 100; i++ {
		intChan2 <- i * 2	// 放100个数据到channel
	}
	// 遍历管道不能使用普通的for循环
	
/*
	channel 支持 for--range 的方式进行遍历，请注意两个细节
	1) 在遍历时，如果 channel 没有关闭，则回出现 deadlock 的错误
	2) 在遍历时，如果 channel 已经关闭，则会正常遍历数据，遍历完后，就会退出遍历。
*/
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v =", v)
	}
}