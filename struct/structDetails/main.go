package main

import (
	"fmt"
)

// 结构体
type Point struct {
	x int
	y int
}

// 结构体
type Rect struct {
	leftUp, rightDown Point
}

// 结构体
type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	// r1有四个int，在内存中是连续分布
	// 打印地址
	fmt.Printf("r1.leftUp.x 地址=%v r1.leftUp.y 地址=%v r1.rightDown.x 地址=%v r1.rightDown.y 地址=%v \n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	// r2有两个*point类型，这两个*point类型的本身地址也是连续的，
	// 但是他们指向的地址不一定是连续的

	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}

	// 打印地址
	fmt.Printf("r1.leftUp 本身地址=%v r1.rightDown 本身地址=%v \n",
		&r2.leftUp, &r2.rightDown)

	fmt.Printf("r1.leftUp 地址=%p r1.rightDown 地址=%p \n",
		r2.leftUp, r2.rightDown)
}
