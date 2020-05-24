package demo4_3

// ------------------------------------------------------------------------
// 定义一个抽象的产品
// ------------------------------------------------------------------------
type Product interface {
	SetName(name string)
	GetName() string
}

// ------------------------------------------------------------------------
// 实现具体的产品：产品1
// ------------------------------------------------------------------------

type Product1 struct {
	name string
}

func (p *Product1) SetName(name string) {
	p.name = name
}

func (p *Product1) GetName() string {
	return "产品1的名称：" + p.name
}

// ------------------------------------------------------------------------
// 实现具体的产品：产品2
// ------------------------------------------------------------------------

type Product2 struct {
	name string
}

func (p *Product2) SetName(name string) {
	p.name = name
}

func (p *Product2) GetName() string {
	return "产品2的名称：" + p.name
}

// ------------------------------------------------------------------------
// 定义一个抽象的产品工厂
// ------------------------------------------------------------------------

type ProductFactory interface {
	Create() Product
}

// ------------------------------------------------------------------------
// 实现具体的产品工厂
// ------------------------------------------------------------------------

type Product1Factory struct{}

func (p *Product1Factory) Create() Product {
	return &Product1{}
}

type Product2Factory struct{}

func (p *Product2Factory) Create() Product {
	return &Product2{}
}
