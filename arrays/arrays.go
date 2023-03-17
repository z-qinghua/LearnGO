// @program:     LearnGo
// @file:        arrays.go
// @create:      2022-09-19 15:38
// @description:

package main

import "fmt"

func printarray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{5, 6, 7, 8, 9}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3, grid)

	//for i := 0; i < len(arr3); i++ {
	//	fmt.Println(arr3[i])
	//}
	for _, v := range arr3 {
		fmt.Println(v)
	}

	fmt.Println("printarray(arr1)")
	printarray(arr1)

	fmt.Println("printarray(arr3)")
	printarray(arr3)

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)
}
