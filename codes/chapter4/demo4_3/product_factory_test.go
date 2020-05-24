package demo4_3

import (
	"fmt"
	"testing"
)

func TestProduct1Factory_Create(t *testing.T) {
	var productFactory1 ProductFactory
	productFactory1 = &Product1Factory{}
	product := productFactory1.Create()
	product.SetName("IPhone")
	fmt.Println(product.GetName())
}

func TestProduct2Factory_Create(t *testing.T) {
	var productFactory2 ProductFactory
	productFactory2 = &Product2Factory{}
	product := productFactory2.Create()
	product.SetName("iMac")
	fmt.Println(product.GetName())
}
