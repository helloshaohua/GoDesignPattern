package demo5_1

import "testing"

func TestComponent1_Operate(t *testing.T) {
	component1 := &Component1{}
	component1.Operate()
}

func TestDecorator1_Operate(t *testing.T) {
	d := &Decorator1{&Component1{}}
	d.Operate()
}
