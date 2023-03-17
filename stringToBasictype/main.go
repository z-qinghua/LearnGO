package main

import (
	"fmt"
	"strconv"
)

// string转成基本数据类型
// 要确保能转换为有效数据，如果失败，会设置为默认值
func main() {
	// string to bool
	var str string = "true"
	var b bool

	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n", b, b)

	// string to int
	var str1 string = "1234"
	var num1 int64
	num1, _ = strconv.ParseInt(str1, 10, 64)
	fmt.Printf("num1 type %T num1=%v\n", num1, num1)

	// string to float
	var str2 string = "12.34"
	var f1 float64
	f1, _ = strconv.ParseFloat(str2, 64)
	// 如果希望得到float32，可再做转换
	fmt.Printf("num2 type %T num2=%v", f1, f1)
}
