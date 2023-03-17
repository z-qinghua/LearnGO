package main

import (
	"fmt"
)

// 声明一个接口
type Usb interface {
	Start()
	Stop()
}

// 让Phone实现Usb接口的方法
type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) Call() {
	fmt.Println("phone is calling")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

// 让Camera实现Usb接口的方法
type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {
}

func (computer Computer) Working(usb Usb) {
	usb.Start()
	// 如果usb是指向Phone结构体变量，则还需要调用Call方法
	// 类型断言
	if phone, ok := usb.(Phone); ok {  //类型断言
		phone.Call()
	}
	usb.Stop()
}

func main() {

	// 定义一个Usb接口数组，可以存放Phone和Camera结构体变量
	// 这里就体现出多态多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"xiaomi"}
	usbArr[2] = Camera{"nikang"}

	for _, val := range usbArr {
		fmt.Println(val)
	}
	// fmt.Println(usbArr)

	//Phone 还有一个特有的方法 call()，请遍历 Usb 数组，如果是 Phone 变量，
	//除了调用 Usb 接口声明的方法外，还需要调用 Phone 特有方法 call. =》类型断言
	fmt.Println("---------------")
	var computer Computer
	for _, val := range usbArr {
		computer.Working(val)
		fmt.Println()
	}
}
