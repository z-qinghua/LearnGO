/*
需求：
	1. 要求统计 1-200000 的数字中，哪些是素数？这个问题在本章开篇就提出了，现在我们有 goroutine
	和 channel 的知识后，就可以完成了 [测试数据: 80000]
分析思路：
	1. 传统的方法，就是使用一个循环，循环的判断各个数是不是素数【ok】。
	2. 使用并发/并行的方式，将统计素数的任务分配给多个(4 个)goroutine 去完成，完成任务时间短

*/

package main

import (
	"fmt"
	"time"
)

//向 intChan 放入 1-8000 个数
func putNum(intChan chan int)  {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}

	// 关闭intChan
	close(intChan)
}

// 从 intChan 取出数据，并判断是否为素数,如果是，就放入到 primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool)  {
	// 使用fot循环
	var flag bool

	for {
		time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan

		if !ok {//intChan 取不到
			break
		}
		flag = true //假设是素数
		// 判断num是不是素数
		for i := 2; i < num; i++ {
			if num % i ==0 {
				// 不是素数
				flag = false
				break
			}
		}
		if flag {
			// 素数放到primeChan
			primeChan <- num
		}
		
	}
	fmt.Println("有一个primeNum协程因为取不到数据，退出")
	// 这里还不能关闭primChan
	// 向exitChan写入true
	exitChan<- true
}

func main()  {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) // 放入结果
	
	// 标识退出的管道
	exitChan := make(chan bool, 4)

	//开启一个协程，向 intChan 放入 1-8000 个数
	go putNum(intChan)

	//开启 4 个协程，从 intChan 取出数据，并判断是否为素数,如果是，就放入到 primeChan
	for i := 1; i <= 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//这里我们主线程，进行处理
	//直接
	go func ()  {
		for i := 0; i < 4; i++ {
			<- exitChan
		}
		//当我们从 exitChan 取出了 4 个结果，就可以放心的关闭 primeChan
		close(primeChan)
	}()

	//遍历我们的 primeChan ,把结果取出
	for {
		val, ok := <- primeChan
		if !ok {
			break
		}
		// 输出结果
		fmt.Println("素数 =", val)
	}
	fmt.Println("main 线程退出")
}

// 结论：使用 go 协程后，执行的速度，比普通方法提高至少 4 倍