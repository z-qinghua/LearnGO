package main

import "fmt"

func main() {
	// 传统遍历方式，默认按字节取值(此方式无法正确输出汉字)
	// 解决办法：将str转化成[]rune切片
	// var str string = "Hello,World!"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c\n", str[i])
	// }

	var str string = "Hello,World!北京"
	str1 := []rune(str)
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c\n", str1[i])
	}

	// for-range遍历方式(推荐)
	var str2 string = "Hello,World!北京" //汉字占3个字节
	for index, value := range str2 {
		fmt.Printf("index=%d value=%c\n", index, value)
	}

	// while and do-while 实现（go本身没有，用for实现）

	// while

	// 循环变量初始化
	var i int = 0

	for {
		if i > 10 {
			break
		}
		fmt.Println("Hello,world!", i)
		i++
	}

	// do-while，循环体至少执行一次

	// 循环变量初始化
	var j int = 1

	for {
		fmt.Println("Hello,world!", j)
		j++ // 循环变量的迭代
		if j > 10 {
			break
		}
	}
}
