package main

import (
	"fmt"
	"sync"
	"time"
)

/*
需求：
	现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到 map 中。最后显示出来。
	要求使用 goroutine 完成

分析思路：
	1) 使用 goroutine 来完成，效率高，但是会出现并发/并行安全问题.
	2) 这里就提出了不同 goroutine 如何通信的问题

代码实现
	1) 使用 goroutine 来完成(看看使用 gorotine 并发完成会出现什么问题? 然后我们会去解决)
	2) 在运行某个程序时，如何知道是否存在资源竞争问题。 方法很简单，在编译该程序时，增加一个参数 -race 即可
*/

// 需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到 map 中。
// 最后显示出来。要求使用 goroutine 完成

// 思路
// 1. 编写一个函数，来计算各个数的阶乘，并放入到 map 中.
// 2. 我们启动的协程多个，统计的将结果放入到 map 中
// 3. map 应该做出一个全局的

var (
	myMap = make(map[int]int, 200)
	// 声明一个全局的互斥锁
	// lock是一个全局的互斥锁
	// sync是包：synchornized同步
	// mutex:是互斥
	lock sync.Mutex
)

// test 函数就是计算 n!, 让将这个结果放入到 myMap
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//这里我们将 res 放入到 myMap
	// 加锁
	lock.Lock()
	myMap[n] = res
	// 解锁
	lock.Unlock()
}

func main() {
	// 我们这里开启多个协程完成这个任务[200 个]
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	//休眠 10 秒钟【第二个问题 】
	time.Sleep(time.Second * 10)

	lock.Lock()
	//这里我们输出结果,变量这个结果
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
