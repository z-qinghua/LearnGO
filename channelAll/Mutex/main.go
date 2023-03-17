/*
互斥是 Go 中一个简单的概念。在我解释它之前，先要明白什么是竞态条件。
goroutines 都有自己的独立的调用栈，因此他们之间不分享任何数据。
但是有一种情况是数据存放在堆上，并且被多个 goroutines 使用。
多个 goroutines 试图去操作一个内存区域的数据会造成意想不到的后果.。
*/
package main

import (
	"fmt"
	"sync"
)

var i int // i == 0
// goroutine increment gloabl variable i
func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	// 此处不加锁会引起竞态条件
	m.Lock()
	i = i + 1
	m.Unlock() // release lock
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &m)
	}
	// wait until all 1000 goroutine are done
	wg.Wait()

	// value of i should be 1000
	fmt.Println("value of i after 1000 operators is", i)
}

// 在任何情况下，都不应该依赖 Go 的调度算法，而应该实现自己的逻辑来同步不同的 goroutine.

// 实现方法之一就是使用我们上面提到的互斥锁。互斥锁是一个编程概念，
// 它保证了在同一时间只能有一个线程或者协程去操作同一个数据。当一个协程想要操作数据的时候，
// 必须获取该数据的一个锁，操作完成后必须释放锁，如果没有获取到该数据的锁，那么就不能操作这个数据。
