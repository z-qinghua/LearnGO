package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age  int
}

func main() {
	// 定义一个存放任意数据类型的管道3个数据
	// var  allChan chan interface{}
	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{"小花猫", 4}
	allChan <- cat

	// 我们希望得到管道中的第三个元素，先将前两个推出

	<-allChan
	<-allChan

	newCat := <-allChan //从管道中取出的Cat是什么

	fmt.Printf("newCat=%T , newCat=%v\n", newCat, newCat)
	// 下面的写法错误，编译不通过
	// fmt.Printf("newCat.Name=%v", newCat.Name)

	a := newCat.(Cat) // 类型断言

	fmt.Printf("newCat.Name=%v\n", a.Name)
}
