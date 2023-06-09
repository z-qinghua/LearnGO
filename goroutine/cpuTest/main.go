package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum =", cpuNum)

	// 可以自己设置使用多个cpu
	runtime.GOMAXPROCS(cpuNum - 3)
	fmt.Println("ok!")
}
