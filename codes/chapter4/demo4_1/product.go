package demo4_1

// ------------------------------------------------------------------------
// 实现一个抽象的产品
// ------------------------------------------------------------------------
type Product interface {
	SetName(name string)
	GetName() string
}

// ------------------------------------------------------------------------
// 实现具体的产品：产品1
// ------------------------------------------------------------------------

type product1 struct {
	name string
}

func (p *product1) SetName(name string) {
	p.name = name
}

func (p *product1) GetName() string {
	return "产品1的名称：" + p.name
}

// ------------------------------------------------------------------------
// 实现具体的产品：产品2
// ------------------------------------------------------------------------

type product2 struct {
	name string
}

func (p *product2) SetName(name string) {
	p.name = name
}

func (p *product2) GetName() string {
	return "产品2的名称：" + p.name
}

// ------------------------------------------------------------------------
// 实现简单工厂类
// ------------------------------------------------------------------------

type productType int

const (
	p1 = iota
	p2
)

type ProductFactory struct{}

func (pf ProductFactory) Create(pt productType) Product {
	if pt == p1 {
		return &product1{}
	}
	if pt == p2 {
		return &product2{}
	}
	return nil
}
