package main

import (
	"fmt"
)

func main() {
	// 管道可以声明为只读或只写

	// 1. 在默认情况下，管道是双向的
	// var chan1 chan int

	// 2. 声明为只写

	chan2 := make(chan<- int, 3)
	chan2 <- 20
	// num := <- chan2 //error

	// 3. 声明为只读

	chan3 := make(<-chan int, 3)
	// num2 := <-chan3

	// chan3 <- 30 //error

	// fmt.Println("num2 = ", num2)
	fmt.Printf("Data type of chan2 is '%T'\n", chan2)
	fmt.Printf("Data type of chan3 is '%T'\n", chan3)
}
