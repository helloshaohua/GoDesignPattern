package unreasonable

import "fmt"

type Bike struct{}

func (b *Bike) Move() {
	fmt.Println("骑自行车")
}

type Car struct{}

func (b *Car) Move() {
	fmt.Println("开车")
}

type Person struct {
	name string
}

func (p *Person) RideBike(bike Bike) {
	bike.Move()
}

func (p *Person) DriveCar(car Car) {
	car.Move()
}
