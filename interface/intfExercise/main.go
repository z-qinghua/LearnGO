package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1. 声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

// 2. 声明一个Hero结构体切片类型
type HeroSlice []Hero

// 3. 实现Interface接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less方法就是决定你使用什么标准进行排序
// 1. 按Hero年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
	// 也可以对name进行排序
	// return hs[i].Name < hs[j].Name
}

// 交换
func (hs HeroSlice) Swap(i, j int) {
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	hs[i],hs[j] = hs[j], hs[i]
}

func main() {
	// 先定义一个数组切片
	var intSlice = []int{0, -1, 10, 7, 90}

	// 要求对intSlice进行排序
	// 1. 冒泡排序
	// 2. 使用系统提供的方法
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	// 对结构体进行排序
	// 1. 冒泡排序
	// 2. 使用系统提供的方法

	// 测试是否可以对结构体进行排序
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}

	// 排序前的顺序
	fmt.Println("------排序前------")
	for _, val := range heros {
		fmt.Println(val)
	}

	// 调用sort
	sort.Sort(heros)

	// 排序后的顺序
	fmt.Println("------排序后------")
	for _, val := range heros {
		fmt.Println(val)
	}
}
