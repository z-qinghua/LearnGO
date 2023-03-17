// @program:     LearnGo
// @file:        Hello.go
// @create:      2022-09-17 14:02
// @description:

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func trangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func euler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	//fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func consts() {
	const (
		a, b     = 3, 4
		filename = "abc.txt"
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp    = 1
		java   = 2
		python = 3
	)
	fmt.Println(cpp, java, python)

}

func main() {
	fmt.Println("Hello World！")
	euler()
	trangle()
	consts()
	enums()
}
