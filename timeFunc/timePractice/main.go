package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	start := time.Now().Unix() //秒
	test03()
	end := time.Now().Unix() //秒
	fmt.Printf("test03()耗费时间=%v秒\n", end-start)
}
