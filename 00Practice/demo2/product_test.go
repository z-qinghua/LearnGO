// @program:     LearnGo
// @file:        product_test.go
// @create:      2023-03-17 21:25
// @description:

package demo2

import (
	"fmt"
	"testing"
)

func TestProduct_Create(t *testing.T) {
	product1 := &Product1{}
	product1.SetName("p1")
	fmt.Println(product1.GetName())

	product2 := &Product2{}
	product2.SetName("p2")
	fmt.Println(product2.GetName())
}

func TestProductFactory_Creat(t *testing.T) {
	productFactory := productFactory{}
	product1 := productFactory.Creat(p1)
	product1.SetName("p1")
	fmt.Println(product1.GetName())

	product2 := productFactory.Creat(p2)
	product2.SetName("p2")
	fmt.Println(product2.GetName())
}
