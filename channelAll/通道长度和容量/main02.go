package main

import (
	"fmt"
	"runtime"
)
func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}
func main() {
	fmt.Println("main() started!")
	c := make(chan int, 3)
	go squares(c)
	
	fmt.Println("active goroutine", runtime.NumGoroutine())
	
	c <- 1
	c <- 2
	c <- 3
	c <- 4  // block here

	fmt.Println("active goroutine", runtime.NumGoroutine())

	go squares(c)

	fmt.Println("active goroutine", runtime.NumGoroutine())

	c <- 5
	c <- 6
	c <- 7
	c <- 8  // block here
	fmt.Println("active goroutine", runtime.NumGoroutine())

	fmt.Println("main() stopped!")

}