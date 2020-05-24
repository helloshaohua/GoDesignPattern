# OOP基础

- 类与对象。
- 三大基本特性。
- 五大基本原则。

## Go中的类与对象

Go中用结构体模拟类与对象。

```go
// Bike 表示一个类型或者说一个OOP类
type Bike struct {
	color string // 首字母小写表示属性私有
	Name  string // 首字母大小表示属性公有
}

// 首字母大写表示方法对外公开
func (b *Bike) Move() string {
	return b.color + "的自行车在前进🚲"
}
```

## Go中的三大基本特性

- **_封装_**：首字母大小写方式代表公有私有权限。

```go
type Person struct {
	name string
}

func (p *Person) Eat() {
	fmt.Println(p.name + "在吃饭")
}
```

- **_继承_**：使用内嵌的方式，对结构体struct进行组合。

```go
type Person struct {
	name string
}

func (p *Person) Eat() {
	fmt.Println(p.name + "在吃饭")
}

type Chinese struct {
	Person
	skin string
}

func (c *Chinese) GetSkin() string {
	return c.skin
}
```

- **_多态_**：使用interface来实现。

```go
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
```

## 回顾下五大基本原则

- 单一功能原则。
对象应该具有单一功能。比如说人类，不能将`走路`和`跳跃`定义为一个方法。
- 开闭原则。
类或函数，扩展我们是开放的，修改我们是要关闭的。
- 里氏替换原则。
子类对象可以替换父类对象，而程序逻辑不变。在Go语言中要面向接口去开发，接口是很好去替换的。
- 接口隔离原则。
就是要保证接口的最小化，接口不能滥用，保持接口定义的合理性，需要根据业务来划分接口。
- 依赖反转原则。
就是不要依赖具体的实现，要去面向接口去开发。

### 单一功能原则和开闭原则示例

- Tom的设计

```go
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

func (p *Person) By(action string) {
	switch action {
	case "bike":
		bike := Bike{}
		bike.Move()
	case "car":
		car := Car{}
		car.Move()
	}
}
```
看到上面Tom设计的`Person`类的`By`方法了没有，它通过传入action参数，去调用不同交通工具的移动Move方法，其实这样的设计是违反了`单一功能原则`，就是这个`By`方法它有两层意思，也违反了`开闭原则`，就是说我要增加一个交通工具如`地铁`，还要再增加一个`case`，可以通过下面Kitty的设计来避免。

- Kitty的设计

```go
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
```

Kitty的设计就很好，它是将一个人是开车还是骑自行车的交通方式类型方法定义拆分开了，这样就可以很好的扩展，比如说我要为这个Person类再添加一个坐地铁方式，是不是直接再添加一个方法就可以了，而不用再去修改By方法了，对吧，这样就满足`了单一原则`以及`开闭原则`。

### 依赖反转原则示例

- Tom的设计

```go
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
```

Tom的设计其实违反了`依赖反转原则`，它依赖具体的实现，这么说吧，现在这个宝马汽车需要将V6引擎升级为V8引擎，你是不是还要修改V6Engine类型到V8Engine类型，可以通过下面的方式去避免。

- Kitty的设计

```go
type Engine interface {
	Run() string
}

type BMWCar struct {
	engine Engine
}

func (bmw *BMWCar) RunEngine() string {
	return "BMWCar的" + bmw.engine.Run()
}
```

Kitty的设计就很好，将汽车引擎定义成接口，就避免了依赖具体的实现，比如说上面的宝马汽车之前是V6引擎，现在要升级为V8引擎，那这里的代码其实不用改动了，添加一个V8引擎并将Engine接口实现即可。

很多时候我们的设计原则，就是要去更好的面对业务的变化。Tom的设计不是说不好，在面对需求变化的时候，可能不能很好的应对。