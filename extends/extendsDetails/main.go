package main

import "fmt"

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

type Monster struct {
	Name string
	Age  int
}

type E struct {
	Monster
	int //匿名字段时基本数据类型
	n   int
}

func main() {
	// 嵌套匿名结构体后，也可以在创建结构体变量（实例）时，直接指定各个匿名结构体字段的值
	tv := TV{Goods{"电视机001", 5000.99}, Brand{"海尔", "山东"}}

	tv2 := TV{
		Goods{
			Price: 5000.99,
			Name:  "电视机002",
		},
		Brand{
			Name:    "夏普",
			Address: "北京",
		},
	}
	fmt.Println("tv", tv)
	fmt.Println("tv2", tv2)

	tv3 := TV2{&Goods{"电视机003", 7000.99}, &Brand{"长虹", "湖南"}}
	tv4 := TV2{
		&Goods{
			Name:  "电视机004",
			Price: 9000.99,
		},
		&Brand{
			Name:    "创维",
			Address: "四川",
		},
	}
	fmt.Println("tv3", *tv3.Goods, *tv3.Brand)
	fmt.Println("tv4", *tv4.Goods, *tv4.Brand)
	fmt.Println(tv4.Goods.Name)

	var e E
	e.Name = "狐狸精" //直接访问匿名结构体中字段
	e.Age = 500
	e.int = 20 //匿名字段赋值
	e.n = 30
	fmt.Println("e=", e)

}
