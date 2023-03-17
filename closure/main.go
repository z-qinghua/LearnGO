// @program:     LearnGo
// @file:        main.go
// @create:      2022-10-07 10:08
// @description:

//闭包
package main

import "fmt"

func sub() func() {
	i := 10 //自由变量
	fmt.Printf("%p\n", &i)
	b := func() {
		fmt.Printf("%p\n", &i)
		i--
		fmt.Println(i)
	}
	return b
}

func add(base int) func(int) int {
	return func(i int) int {
		fmt.Printf("%p\n", &base)
		base += i
		return base
	}
}

func main() {
	f := sub()
	f()
	f()
	temp1 := add(10)
	fmt.Println(temp1(1), temp1(2)) //base=11,temp1(1)==base==1 base=13

	temp2 := add(10)                //base重置10
	fmt.Println(temp2(1), temp2(2)) //11，13
}
