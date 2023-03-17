// @program:     LearnGo
// @file:        多协程协同工作.go
// @create:      2022-10-24 17:03
// @description:

package main

import "fmt"

//写两个协程，一个计算数字的平方，一个计算数字的立方

func square(c chan int) {
	fmt.Println("[square] reading")
	num := <-c
	c <- num * num
}

func cube(c chan int) {
	fmt.Println("[cube] reading")
	num := <-c
	c <- num * num * num
}

func main() {
	fmt.Println("main() started")

	squareChan := make(chan int) //无缓存通道
	cubeChan := make(chan int)

	//分别运行square、cube协程
	go square(squareChan)
	go cube(cubeChan)

	testNum := 3 //调度权还在主线程，所以执行赋值操作
	fmt.Println("[main] sent testNum to squarechan")

	squareChan <- testNum //主线程阻塞直到通道数据被读取，一旦被读取，主线程继续执行

	fmt.Println("[main] resuming")
	fmt.Println("[main] sent testNum to cubChan")

	cubeChan <- testNum

	fmt.Println("[main] resuming")
	fmt.Println("[main] reading from channels")

	//试图从两个通道读取数据，此时线程可能阻塞直到有数据写入通道
	squareVal, cubeVal := <-squareChan, <-cubeChan
	//当数据被写入通道，主线程继续运行，最后打印
	sum := squareVal + cubeVal

	fmt.Println("[main] sum of square and cube of", testNum, "is", sum)
	fmt.Println("[main] main() stopped")
}
