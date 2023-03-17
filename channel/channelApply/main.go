/*
	请完成goroutine和channel协同工作的案例，具体要求:
	1. 开启一个writeData协程，向管道intChan中写入50个整数
	2. 开启一个readData协程，从管道intChan中读取writeData写入的数据
	3. 注意: writeData和readDate操作的是同一个管道
	4. 主线程需要等待writeData和readDate协程都完成工作才能退出[管道
*/

package main

import (
	"fmt"
)

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		// 放入数据
		intChan <- i
		fmt.Printf("writeData 数据=%v\n", i)
	}
	// 关闭
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 数据=%v\n", v)
	}
	// 读完数据，即任务完成
	exitChan <- true
	close(exitChan)
}

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
