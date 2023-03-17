package main

import (
	"fmt"
	"strconv"
)

func main() {
	//练习基本数据转string
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var mychar byte = 'h'
	var str string //空的str

	//使用第一种方式fmt.Sprintf转string
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type %T str=%q\n", str, str)

	//使用第二种方式strconv包转string
	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	//str = strconv.FormatFloat(num2, 'f', 10, 64)	
	//'f':格式 10；表示小数位保留10位 64：表示这个小数是float64
	str = strconv.FormatFloat(num2, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(b)
	fmt.Printf("str type %T str=%q\n", str, str)

	//strconv包中有个Itoa函数
	var num3 int64 = 234
	str = strconv.Itoa(int(num3))//int64要转换成int
	fmt.Printf("str type %T str=%q\n", str, str)
}
