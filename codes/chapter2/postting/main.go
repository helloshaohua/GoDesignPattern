package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) Eat() {
	fmt.Println(p.name + "在吃饭")
}
