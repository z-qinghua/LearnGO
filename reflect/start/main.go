package main

import (
	"fmt"
	"reflect"
)

// 请编写一个案例，演示对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作

// 专门演示反射
func reflectTest01(b interface{}) {
	// 通过反射获取传入变量的type,kind值
	// 1. 先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType =", rType)

	// 2. 获取大reflect.Value
	rVal := reflect.ValueOf(b) //rVal类型为reflect.Value

	n2 := 2 + rVal.Int()
	fmt.Println("n2 =", n2)

	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)

	// 3. 获取变量对应的kind 

	// 下面将rVal转成interface{}
	iV := rVal.Interface()

	// 将interface{}通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Printf("num2=%v num2 type=%T\n", num2, num2)


}

// 专门演示反射（对结构体的反射）
func reflectTest02(b interface{})  {
	// 通过反射获取的传入的变量的type,kind值
	// 1. 先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType =", rType)

	// 2. 获取到reflect.Value
	rVal := reflect.ValueOf(b)

	// 下面将rVal转成interface{}
	iV := rVal.Interface()
	fmt.Printf("iV=%v iV type=%T\n", iV, iV)

	// 将interface{}通过断言转成需要的类型
	// 这里，简单使用带检测的类型断言

	// 使用switch的断言形式更加灵活
	
	if stu, ok := iV.(Student); ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
}

type Student struct {
	Name string
	Age int
}

type Monster struct {
	Name string
	Age int
}


func main()  {
	//请编写一个案例，
	//演示对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作

	// 先定义一个int
	var num int = 3
	reflectTest01(num)

	// 2. 定义一个Student的实例
	stu := Student{
		Name: "tom",
		Age: 20,
	}
	reflectTest02(stu)
}