package main

import (
	"fmt"
)

// 声明结构体,结构体中可以没有任何字段
type MethodUtils struct {
}

// 定义小小计算器结构体(Calcuator)，实现加减乘除四个功能
type Calcuator struct {
	Num1 float64
	Num2 float64
	// operator byte
}

// 1.编写结构体(MethodUtils)，编程一个方法，方法不需要参数，在方法中打印一个 10*8 的矩形，在 main 方法中调用该方法
func (mu MethodUtils) print() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

// 2. 编写一个方法，提供 m 和 n 两个参数，方法中打印一个 m*n 的矩形
func (mu MethodUtils) print2(m int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

// 3. 编写一个方法算该矩形的面积(可以接收长 len，和宽 width)， 将其作为方法返回值。在 main方法中调用该方法，接收返回的面积值并打印
func (mu MethodUtils) area(len float64, width float64) float64 {
	return len * width
}

// 4. 编写方法：判断一个数是奇数还是偶数
func (mu MethodUtils) judgeNum(num int) {
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}
}

// 5. 根据行、列、字符打印 对应行数和列数的字符，比如：行：3，列：2，字符*,则打印相应的效果
func (mu MethodUtils) print3(m int, n int, key string) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

// 6. 定义小小计算器结构体(Calcuator)，实现加减乘除四个功能

// 实现形式 1：分四个方法完成:
func (calcuator *Calcuator) getSum() float64 {
	return calcuator.Num1 + calcuator.Num2
}

func (calcuator *Calcuator) getSub() float64 {
	return calcuator.Num1 - calcuator.Num2
}

// 实现形式 2：用一个方法搞定
func (calcuator *Calcuator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = calcuator.Num1 + calcuator.Num2
	case '-':
		res = calcuator.Num1 - calcuator.Num2
	case '*':
		res = calcuator.Num1 * calcuator.Num2
	case '/':
		res = calcuator.Num1 / calcuator.Num2
	default:
		fmt.Println("输出出现错误")
	}
	return res
}

func main() {
	var mu MethodUtils
	mu.print()

	fmt.Println()
	mu.print2(5, 4)

	fmt.Println()
	areaRes := mu.area(4.5, 3.5)
	fmt.Println("矩形面积=", areaRes)

	mu.judgeNum(3)
	mu.print3(5, 5, "*")

	var calcuator Calcuator
	calcuator.Num1 = 1.2
	calcuator.Num2 = 2.2

	fmt.Println("实现方式1:")
	fmt.Println("sum=", calcuator.getSum())
	fmt.Println("sub=", calcuator.getSub())

	fmt.Println("实现方式2:")
	res1 := calcuator.getRes('+')
	res2 := calcuator.getRes('-')
	fmt.Println("sum~=", res1)
	fmt.Println("sub~=", res2)
}
