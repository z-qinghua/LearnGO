// @program:     LearnGo
// @file:        func.go
// @create:      2022-09-19 11:59
// @description:

package main

import "fmt"

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation:" + op)
	}
}

func div(a, b int) (q int, r int) { //函数可以返回多个值
	//q = a / b
	//r = a % b

	return a / b, a % b
}

//交换数值
/*
func swap(a, b *int) {
	*a, *b = *b, *a
}
*/
func swap(a, b int) (int, int) {
	return b, a
}

// go支撑可变参数
// 编写一个函数，可求1到多个int的和

// return返回值只有一个时，返回值类型列表可省略()
// 形参中有可变参数，则可变参数需要放到形参列表最后
func sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum = sum + args[i] // 表示取出args切片第i个元素值
	}
	return sum
}

func main() {
	fmt.Println(eval(3, 4, "-"))
	q, r := div(3, 4)
	fmt.Println(q, r)

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)

	fmt.Println("1到多个int的和=", sum(1, 2, 3, 4, 5))
}
