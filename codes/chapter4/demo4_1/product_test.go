package demo4_1

import (
	"fmt"
	"testing"
)

func TestProduct(t *testing.T) {
	p := &product1{}
	p.SetName("产品1")
	fmt.Println(p.GetName())

	p2 := &product2{}
	p2.SetName("产品2")
	fmt.Println(p.GetName())
}

func TestProductFactory_CreateProduct(t *testing.T) {
	p := &ProductFactory{}
	product1 := p.Create(p1)
	product1.SetName("产品1")
	fmt.Println(product1.GetName())

	product2 := p.Create(p2)
	product2.SetName("产品2")
	fmt.Println(product2.GetName())
}
