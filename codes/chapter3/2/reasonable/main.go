package unreasonable

type Engine interface {
	Run() string
}

type BMWCar struct {
	engine Engine
}

func (bmw *BMWCar) RunEngine() string {
	return "BMWCar的" + bmw.engine.Run()
}
