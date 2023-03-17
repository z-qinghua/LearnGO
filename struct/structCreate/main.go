package main

import (
	"fmt"
)

// 声明一个结构体
type Person struct {
	Name string
	Age  int
}

func main() {
	// 方式1-直接声明
	var person Person
	person.Name = "zeng"
	fmt.Println(person.Name)

	// 方式2-{}(推荐)
	p2 := Person{"marry", 20}
	fmt.Println(p2)

	// 方式3-&
	var p3 *Person = new(Person)
	(*p3).Name = "smith"
	// p3.Name = "john"

	(*p3).Age = 30
	// p3.Age = 40

	fmt.Println(*p3)

	// 方式4-{}
	var p4 *Person = &Person{}
	(*p4).Name = "scott"
	// p4.Name = "scott-"

	(*p4).Age = 88
	// p4.Age = 100

	fmt.Println(*p4)

}
