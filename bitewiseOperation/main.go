package main

import "fmt"

func main() {
	// 位运算
	fmt.Println(2 & 3)  // 2
	fmt.Println(2 | 3)  // 3
	fmt.Println(2 ^ 3)  // 1
	fmt.Println(-2 ^ 2) // -4

	// 右移运算符>>：低位溢出，符号位不变，并用符号位补溢出的高位
	// 左移运算符<<：符号位不变，低位补0
	a := 1 >> 2 // 0
	b := 1 << 2 // 4
	fmt.Println(a)
	fmt.Println(b)
}
