// @program:     LearnGo
// @file:        defer.go
// @create:      2022-10-06 13:46
// @description:

package main

import "fmt"

func tryDefer() {
	defer fmt.Println(1) //函数快关闭时执行
	defer fmt.Println(2)
	defer fmt.Println(3)
	//output:3 2 1,相当于有个栈，先进后出

	fmt.Println(4)
}

//defer执行顺序
func basic() {
	fmt.Println("A")
	defer fmt.Println(1)
	fmt.Println("B")
	//多个defer，后注册的先执行
	defer fmt.Println(2)
	fmt.Println("C")
}

func defer_exe_time() (i int) {
	i = 9
	defer func() {
		fmt.Printf("i=%d\n", i) //打印5，而非9
	}() //5
	defer func(i int) {
		fmt.Printf("i=%d\n", i)
	}(i) //9
	defer fmt.Printf("i=%d\n", i) //i=5 变量在注册defer时被拷贝或计算
	return 5
}

func defer_panic() {
	defer fmt.Println("111")
	n := 0
	defer func() {
		fmt.Println(2 / n)
		defer fmt.Println("2222")
	}() //不影响下一个defer执行
	defer fmt.Println("3333")
}
func main() {
	tryDefer()
	basic()
	//panic(0)//panic内部会去调用os.Exit(2)
	defer_exe_time()
	defer_panic()
}
