package main

import "fmt"

// defer的价值在于及时释放函数创建的资源
func sum(n1 int, n2 int) int {
	// 当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈（defer栈）
	// 当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行

	// 将defer语句压入栈时，会同时将相关的值拷贝同时入栈
	defer fmt.Println("ok1, n1=", n1)
	defer fmt.Println("ok2 n2=", n2)

	// 增加一句话
	n1++
	n2++
	n := n1 + n2 // 32
	fmt.Println("ok3 n", n)
	return n
}

func main() {

	res := sum(10, 20)
	fmt.Println("sum=", res)
}
