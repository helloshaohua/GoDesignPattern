package polymorphisn

import "fmt"

type Human interface {
	Speck()
}

type American struct {
	name     string
	Language string
}

func (a *American) Speck() {
	fmt.Println("美国人" + a.name + "说" + a.Language)
}

type Chinese struct {
	name     string
	Language string
}

func (c *Chinese) Speck() {
	fmt.Println("美国人" + a.name + "说" + a.Language)
}
