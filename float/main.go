package main

import "fmt"

//Golang浮点型有固定长度和范围，不受os的影响。
func main() {

	//尾数部分可能造成精度损失
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=", num3, "num4=", num4)

	//Golang浮点型默认为float64，开发中推荐使用。
	num5 := 1.1
	fmt.Printf("浮点型默认为：%T\n",num5)

	//十进制形式，如：5.12、.12(必须有小数点)
	num7 := 5.12
	fmt.Println(num7)

	//科学计数法
	num8 := 5.1234e2
	fmt.Println(num8)
}
