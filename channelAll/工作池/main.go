/*
顾名思义，一个工作池 并发执行某项工作的协程集合。 在 WaitGroup 章节，
我们已经用到的多个协程执行一个任务，但是他们并没有执行特定的工作，只是 sleep 了一下。
如果你向协程中传一个通道，他们可以去完成一些工作，变成一个工作池。

所以工作池其实就是维护了多个工作协程，这些协程的功能是可以收到任务，执行任务并返回结果。
他们完成任务后我们就可以收到结果。这些协程使用相同的通道来达到自己的目的。
*/
package main

import (
	"fmt"
	"time"
)

func sqrWorker(tasks <-chan int, results chan<- int, instance int) {
	for num := range tasks {
		time.Sleep(time.Millisecond) // simulating block task 使用了 sleep 操作来模拟阻塞操作
		fmt.Printf("[worker %v] Sending result by worker %v\n", instance, instance)
		results <- num * num
	}
}

func main() {
	fmt.Println("main() started")

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// lauching 3 worker goroutines
	for i := 0; i < 3; i++ {
		go sqrWorker(tasks, results, i)
	}

	// passing 5 tasks
	for i := 0; i < 5; i++ {
		tasks <- i * 2 // no-blocking as buffer capacity is 10
	}
	fmt.Println("[main] wrote 5 tasks")

	// closing tasks
	close(tasks)

	// receving results from all workers
	for i := 0; i < 5; i++ {
		result := <-results // blocking because buffer is empty
		fmt.Println("[main] Result", i, ":", result)
	}

	fmt.Println("main() stopped")
}
