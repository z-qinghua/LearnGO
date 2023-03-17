package main

import (
	"fmt"
	"time"
)

func main() {
	// 看看日期和时间相关函数和方法使用

	// 1. 获取当前时间
	now := time.Now() // 返回的是一个结构体
	fmt.Printf("now=%v now type=%T\n", now, now)

	// 2. 获取年、月、日、时、分、秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	// 格式化日期时间

	// 1. 第一种方式：使用printf或者Sprintf
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	datestr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("datestr=%v\n", datestr)

	// 2. 第二种方式：使用time.Format()
	fmt.Printf(now.Format("2006-1-6"))
	fmt.Println()

	fmt.Printf(now.Format("15-04-05"))
	fmt.Println()

	fmt.Printf(now.Format("2006-1-6 15:04:05")) // 数字固定，可自由组合
	fmt.Println()

	// unix和unixNano的使用
	fmt.Printf("unix时间戳=%v uniNanos时间戳=%v\n", now.Unix(), now.UnixNano())
}
