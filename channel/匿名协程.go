// @program:     LearnGo
// @file:        匿名协程.go
// @create:      2022-10-24 20:30
// @description:

package main

import "fmt"

func main() {
	fmt.Println("main() started")
	c := make(chan string)

	//launch anonymous goroutine
	go func(c chan string) {
		fmt.Println("Hello " + <-c + "!")
	}(c)
	c <- "John"
	fmt.Println("main() stopped")
}
