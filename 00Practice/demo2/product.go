// @program:     LearnGo
// @file:        product.go
// @create:      2023-03-17 21:25
// @description:

package demo2

// 实现一个抽象的产品

type Product interface {
	SetName(name string)
	GetName() string
}

// 实现具体的产品1

type Product1 struct {
	name string
}

func (p1 *Product1) SetName(name string) {
	p1.name = name
}

func (p1 *Product1) GetName() string {
	return "产品1的name为：" + p1.name
}

// 实现具体的产品2

type Product2 struct {
	name string
}

func (p2 *Product2) SetName(name string) {
	p2.name = name
}

func (p2 *Product2) GetName() string {
	return "产品2的name为：" + p2.name
}

type productType int

const (
	p1 productType = iota
	p2             //1
)

// 实现简单工厂类
type productFactory struct {
}

func (pf productFactory) Creat(productType productType) Product {
	if productType == p1 {
		return &Product1{}
	}
	if productType == p2 {
		return &Product2{}
	}
	return nil
}
