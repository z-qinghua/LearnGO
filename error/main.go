package main

import (
	"errors"
	"fmt"
	_ "time"
)

func test() {
	// 使用defer + recover 来捕获和处理异常
	defer func() {
		err := recover() // recover()内置函数，可以捕获到异常
		if err != nil {
			fmt.Println("test() err", err)
			fmt.Println("发送邮件给admin@")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2 //此处会出现错误
	fmt.Printf("res=%v\n", res)
}

// 自定义错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		// 读取...
		return nil
	} else {
		// 返回一个自定义错误
		return errors.New("读取文件错误")
	}
}

func test02() {
	err := readConf("config.in")
	// 如果读取文件发生错误，就输出这个错误，并终止程序
	if err != nil {
		fmt.Println("err info =", err.Error())
	}
	fmt.Println("test02()继续执行")
}

func main() {

	// 测试
	test()
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("继续执行main下面的代码...\n")
	// 	time.Sleep(time.Second)
	// }

	// 测试自定义错误
	test02()
	fmt.Printf("继续执行main下面的代码...\n")
}
