// @program:     LearnGo
// @file:        main.go
// @create:      2022-10-08 09:19
// @description:

package main

import "fmt"

func add4int(a, b int) int {
	return a + b
}

func add4string(a, b string) string {
	return a + b
}

//泛型
type Addable interface { //定义一组类型
	int | int8 | int16 | int32 | int64 | uint16 | string | complex128 | float32 | float64
}

func add[T Addable](a, b T) T {
	return a + b
}
func main() {
	var a, b int = 1, 2
	var c, d string = "qinghua", "zeng"
	fmt.Println(add4int(a, b))
	fmt.Println(add4string(c, d))
	fmt.Println(add(a, b))
	fmt.Println(add(c, d))
}
