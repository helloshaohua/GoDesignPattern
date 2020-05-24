package unreasonable

type V6Engine struct{}

func (v6 *V6Engine) Run() string {
	return "V6Engine run"
}

type BMWCar struct {
	engine V6Engine
}

func (bmw *BMWCar) RunEngine() string {
	return "BMWCar的" + bmw.engine.Run()
}
