package prototype

import "fmt"

//Prototype 原型接口
type Prototype interface {
	DoSomething()
	Clone() Prototype
}

//prototypeManager 原型管理器：用于管理实现原型接口的实例，在需要克隆实例时提供帮助
type prototypeManager struct {
	prototypeMap map[string]Prototype
}

func NewPrototypeManager() *prototypeManager {
	return &prototypeManager{prototypeMap: make(map[string]Prototype)}
}

func (p *prototypeManager) AddPrototype(pName string, prototype Prototype)  {
	p.prototypeMap[pName] = prototype
}

func (p *prototypeManager) GetPrototype(pName string) Prototype {
	return p.prototypeMap[pName]
}

type Cat struct {
	name string
}

func NewCat(name string) *Cat {
	return &Cat{name: name}
}

func (c *Cat) Name() string {
	return c.name
}

func (c *Cat) DoSomething() {
	fmt.Println("i am cat")
}

func (c *Cat) Clone() Prototype {
	return &Cat{name: c.name}
}

type Dog struct {
	Age uint8
}

func (d *Dog) DoSomething() {
	fmt.Println("i am dog")
}

func (d *Dog) Clone() Prototype {
	return &Dog{Age: d.Age}
}
