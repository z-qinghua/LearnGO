/*
使用通道，我们可以更好地实现生成器。如果生成器的计算开销很大，
那么我们也可以并发生成数据。这样，生成数据的逻辑就不会阻塞主程序。比如生成斐波那契数列.
*/

package main

import (
	"fmt"
)

// fib returns a channel which transports fibonacci numbers
func fib(length int) <-chan int {
	// make buffered channel
	c := make(chan int, length)

	// run generation concurrently
	go func() { // 匿名协程
		for i, j := 0, 1; i < length; i, j = i+j, i {
			c <- i
		}
		close(c)
	}()
	return c // 转化为单向通道返回
}

func main() {
	// read 10 fibonacci numbers from channel returned by 'fib' function
	for fn := range fib(10) {
		fmt.Println("Current fibonacci number is", fn)
	}
}
