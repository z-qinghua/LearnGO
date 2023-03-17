// @program:     LearnGo
// @file:        fib.go
// @create:      2022-10-06 13:25
// @description:

package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}
