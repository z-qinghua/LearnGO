// @program:     LearnGo
// @file:        main.go
// @create:      2022-10-08 14:55
// @description:

package main

import (
	"fmt"
)

func main() {
	var ch chan int
	//未初始化时默认值是nil
	fmt.Printf("ch is nill :%t\n", ch == nil)
	fmt.Printf("len of ch is %d\n", len(ch))

	ch = make(chan int, 10) //初始化，容量为10
	fmt.Printf("len of ch is %d\n", len(ch))
	fmt.Printf("cap of ch is %d\n", cap(ch))

	//写入数据
	for i := 0; i < 10; i++ {
		ch <- 3
	}
	fmt.Println("len of ch is ", len(ch))

	//管道要关闭才能遍历，且遍历时会取出数据

	//此方法不需close(ch)
	L := len(ch)
	for i := 0; i < L; i++ {
		ele := <-ch
		fmt.Println(ele)
	}
	fmt.Println("len of ch is ", len(ch))

	//下面方法，需先close(ch),效果如上
	//close(ch)
	//关闭后不能再写数据
	//for ele := range ch {
	//	fmt.Println(ele)
	//}
	//fmt.Println("len of ch is ", len(ch))
}
