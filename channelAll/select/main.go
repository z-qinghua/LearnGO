package main

import (
	"fmt"
	"time"
)

var start time.Time

func init()  {
	start = time.Now()
}

// func service1(c chan string)  {
// 	// time.Sleep(time.Second * 3)
// 	c <- "Hello from service 1"
// }

// func service2(c chan string)  {
// 	// time.Sleep(time.Second * 5)
// 	c <- "Hello from service 2"
// }

func main()  {
	fmt.Println("main() started", time.Since(start))
	chan1 := make(chan string, 2)
	chan2 := make(chan string, 2)

	// go service1(chan1)
	// go service2(chan2)

	chan1 <- "value 1"
	chan1 <- "value 2"
	chan2 <- "value 1"
	chan2 <- "value 2"

	select {
	case res := <- chan1:
		fmt.Println("response from chan 1", res, time.Since(start))
	case res := <- chan2:
		fmt.Println("resonse from chan 2", res, time.Since(start))
	}
	fmt.Println("main() stopped", time.Since(start))
}

// 在上面的程序中，两个通道在其缓冲区中都有两个值。
// 因为我们向容量为 2 的缓冲区通道分别发送了两个值，
// 所以这些通道发送操作不会阻塞并且会执行下面的 select 块。 
// select 块中的所有 case 操作都不会阻塞，因为每个通道中都有两个值，
// 而我们的 case 操作只需要取出其中一个值。因此，go 运行时会随机选择一个 case 操作并执行其中的代码。
