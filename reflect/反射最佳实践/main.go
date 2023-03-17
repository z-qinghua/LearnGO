// 使用反射来遍历结构体的字段，调用结构体的方法，并获取结构体标签的值

package main

import (
	"fmt"
	"reflect"
)

// 定义一个Monster结构体
type Monster struct {
	Name string `json:"name"`
	Age int `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex string
}

// 方法：返回两个数的和
func (s Monster) GetSum(n1, n2 int) int{
	return n1 + n2 
}

// 方法：接收四个值，给s赋值
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

// 方法：显示s的值
func (s Monster) Print() {
	fmt.Println("--start--")
	fmt.Println(s)
	fmt.Println("--end--")
}

func TestStruct(a interface{})  {
	// 获取reflect.Type类型
	typ := reflect.TypeOf(a)

	// 获取reflect.Value类型
	val := reflect.ValueOf(a)

	// 获取到a对应的类别
	kd := val.Kind()

	// 如果传入的不是struct，就退出
	if kd != reflect.Struct {
		fmt.Println("Expect struct")
		return
	}

	// 获取到该结构体有几个字段
	num := val.NumField()

	fmt.Printf("struct has %d fields\n", num)// 4

	// 变量结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值=%v\n", i, val.Field(i))
		// 获取到struct标签，注意需要通过reflect.Type来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")

		// 如果该字段有tag标签就显示，否则就不显示
		if tagVal != "" {
			fmt.Printf("Field %d:tag=%v\n", i, tagVal)
		}
	}
	// 获取该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	// 方法的排序默认是按照函数名的排序（ASCII码）
	val.Method(1).Call(nil)//获取到第二个方法，调用它

	// 调用结构体的第一个方法Method(0)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) //传入的参数是[]reflect.Value，返回[]reflect.Value
	fmt.Println("res=", res[0].Int()) //返回结果，返回的结果是[]reflect.Value

}


func main()  {
	// 创建一个Monster实例
	var a Monster = Monster{
		Name: "黄鼠狼精",
		Age:400,
		Score:30.8,
	}
	// 将Monster实例传递给TestStruct函数
	TestStruct(a)
}