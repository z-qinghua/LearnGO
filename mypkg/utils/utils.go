package utils // 定义包名时尽量和文件夹名相同，方便调用

import "fmt"

// 同一个包下面不同重复定义函数名和变量名

// 将计算的功能，放到一个函数中，需要使用使调用
var Num1 int = 300

func Cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作复有误...")
	}
	return res
}
