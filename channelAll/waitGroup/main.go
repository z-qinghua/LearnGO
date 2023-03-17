/*
有一种业务场景是你需要知道所有的协程是否已执行完成他们的任务。
这个和只需要随机选择一个条件为 true 的 select 不同，他需要你满足所有的条件都是 `true` 才可以激活主线程继续执行。
这里的 条件 指的是非阻塞的通道操作。

WaitGroup 是一个带着计数器的结构体，这个计数器可以追踪到有多少协程创建，
有多少协程完成了其工作。当计数器为 0 的时候说明所有协程都完成了其工作
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func service(wg *sync.WaitGroup, instance int){
	time.Sleep(time.Second * 2)
	fmt.Println("Service called on instance", instance)
	// Done 方法可以降低计数器的值。他不接受任何参数，因此，它每执行一次计数器就减 1
	wg.Done() //decrement counter 
}

func main() {
	fmt.Println("main() started")

	var wg sync.WaitGroup //create waitgoup (empty struct)

	for i := 1; i <= 3; i++ {
		// Add 方法的参数是一个变量名叫 delta 的 int 类型参数，主要用来内部计数。 内部计数器默认值为 0. 
		// 它用于记录多少个协程在运行。当 WaitGroup 创建后，计数器值为 0，
		// 我们可以通过给 Add 方法传 int 类型值来增加它的数量。 
		// 记住， 当协程建立后，计数器 的值不会自动递增 ，因此需要我们手动递增它。
		wg.Add(1) //increment counter
		go service(&wg, i)
	}
	// Wait 方法用来阻塞当前协程。一旦 计数器 为 0, 协程将恢复运行。 因此，我们需要一个方法去降低计数器的值。
	wg.Wait() //block here
	fmt.Println("main() stopped")
}