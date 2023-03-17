// 扇入是一种多路复用的策略，把几个通道的输入整合到一个输出的通道。扇出是一种多路分解策略，将单个通道的数据分散到多个通道。

package main

import (
	"fmt"
	"sync"
)

// 从切片中读取元素，写入input channel
func getInputChan() <-chan int {
	// 实例化容量为100，int类型的input channel
	input := make(chan int, 100)

	// 用来写道通道的数据
	numbers := []int{0,1,2,3,4,5,6,7,8,9}

	// 启动协程，把数字写到input channel
	go func () {
		for num := range numbers {
			input <- num
		}
		// 数据写入完毕，关闭channel
		close(input)
	}()

	return input
}

// 把input channel读到的数字，做平方运算，再写入output channel
func getSquareChan(input <-chan int) <-chan int {
	// 实例化容量为100的int类型的output channel
	output := make(chan int, 100)

	// 启动协程
	go func() {
		// 遍历input channel，把读到的数字平方运算，写入output
		for num := range input {
			output <- num * num
		}
		// 关闭output channel
		close(output)
	}()

	return output
}

// 返回对'outputsChan'通道合并之后的通道
// 这会产生扇入通道
// 这是一个可变参数的函数
func merge(outputsChan ...<-chan int) <-chan int{
	// 申明WaitGroup
	var wg sync.WaitGroup

	// 实例化merged 通道
	merged := make(chan int, 100)

	// 增加一个计数器，计数器的参数为outputsChan的长度
	// 因为我们将会创建多个goroutine，其中goroutine的数量就是要准备进行合并的通道的数量
	wg.Add(len(outputsChan))

	// 从sc channel读取数据，写入到merged通道
	output := func(sc <-chan int) {
		// 遍历
		for sqr := range sc {
			merged <- sqr
		}
		// 一旦sc通道关闭
		// 在'WaritGroup'上调用'Done'来递减计数器
		wg.Done()
	}

	//把`output`函数运行为groutines，
    // 启动n个协程
    //其中n等于作为函数参数接收的通道数
    //这里我们在`outputsChan`上使用`for range`循环，因此无需手动告诉`n`
	for _, optChan := range outputsChan {
		go output(optChan)
	}

	// 一旦完成，运行goroutine关闭merged通道
	go func() {
		// 等到WaitGroup结束
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	// 步骤1：获取数字通道
	// 通过调用'getInputChan'函数，运行一个goroutine，它将数字发送到返回的通道
	chanInputNums := getInputChan()

	// 步骤2：对多个goroutine进行'扇出'平方操作
	// 这通过多次调用'getSquareChan'函数来完成，其中单个函数调用返回一个通道，该通道发送由'chanInputNums'通道提供的数字的平方
	// 'getSquareChan'函数在内部运行goroutine，同时运行平方操作
	chanOptSqr1 := getSquareChan(chanInputNums)
	chanOptSqr2 := getSquareChan(chanInputNums)

	//步骤3：扇入（合并）`chanOptSqr1`和`chanOptSqr2`输出到合并频道
    // 这是通过调用`merge`函数实现的，该函数将多个通道作为参数
    // 并使用`WaitGroup`和多个goroutines来接收平方数，我们可以发送平方数
    //到 `merged` 通道，并关闭
	chanMergedSqr := merge(chanOptSqr1, chanOptSqr2)

	//步骤4：计算0到9之间的所有整数的平方再求和，大约是'285`
    //这是通过在`chanMergedSqr`上使用`for range`循环来完成的
	sqrSum := 0

	//运行直到`chanMergedSqr`或合并频道关闭
    //当所有goroutines推送到合并频道完成时，在`merge`函数中发生
	for num := range chanMergedSqr {
		sqrSum += num
	}

	//步骤5：当`for循环'完成执行时，在`chanMergedSqr`通道关闭之后打印总和
	fmt.Println("Sum of squares between 0-9 is", sqrSum)
}