// @program:     LearnGo
// @file:        slice.go
// @create:      2022-09-19 16:22
// @description:

package main

import (
	"fmt"
)

func updateSlice(s []int) {
	s[0] = 100
}

// 切片是引用类型，所以在传递时，遵守引用传递机制。
func test(slice []int) {
	slice[0] = 100 // 这里修改slice[0]，会改变实参
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1=", s1)

	s2 := arr[:]
	fmt.Println("s2=", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr[2:])

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr[:])

	fmt.Println("Extending slice")
	arr[0] = 0
	arr[2] = 2
	fmt.Println("arr = ", arr)
	s1 = arr[2:6]
	s2 = s1[3:5] //s1[3],s1[4]
	fmt.Printf("s1=%v,len(s1) = %d,cap(s1) = %d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v,len(s2) = %d,cap(s2) = %d\n",
		s2, len(s2), cap(s2))
	fmt.Println("s1[3:5] = ", s1[3:5])

	fmt.Println("_____________________________________")
	s3 := append(s2, 11)
	s4 := append(s3, 20)
	s5 := append(s4, 30)
	fmt.Println("s3,s4,s5 = ", s3, s4, s5)
	//s4 and s5 no longer view arr.
	fmt.Println("arr = ", arr)

	// 拷贝注意事项
	fmt.Println("------------------------")
	var a []int = []int{1, 2, 3, 4, 5}
	var slice = make([]int, 1)
	fmt.Println(slice) // [0]
	// 下面代码没有问题，可以运行
	copy(slice, a)
	fmt.Println(slice) //[1]

	var slice3 = []int{1, 2, 3, 4}
	fmt.Println("slice=", slice3) // [1,2,3,4]
	test(slice3)
	fmt.Println("slice=", slice3) // [100,2,3,4]
}
