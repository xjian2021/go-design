package abstract_factory

import "fmt"

const (
	CAT_TYPE FactoryType = iota
	DOG_TUPE
)

type FactoryType int

type BasisActionsInterface interface {
	Run()
	Jump()
}

type SpecialActionsInterface interface {
	DoSomething()
}

type ActionsSet interface {
	BasisActionsInterface
	SpecialActionsInterface
}

//ActionsFactory 行为抽象模式工厂接口
type ActionsFactory interface {
	CreateActions(name string) ActionsSet
}

//GetActionsSet 方式一
//利用抽象工厂模式获取自定义的对象工厂
//优点在于在需要新增业务类型时，只要实现接口即可使用(易于扩展)
func GetActionsSet(af ActionsFactory, name string) ActionsSet {
	return af.CreateActions(name)
}

//GetActionsSet2 方式二
//给每种工厂定义标识，在使用时只需要传入标识即可(易于使用)
//优点在于使用者不会直接与具体类型进行互动，而是依靠该构造函数来创建接口的实例，仅使用标识参数来控制生产
func GetActionsSet2(ftype FactoryType, name string) ActionsSet {
	switch ftype {
	case CAT_TYPE:
		return (&CatActions{}).CreateActions(name)
	case DOG_TUPE:
		return (&DogActions{}).CreateActions(name)
	}
	return nil
}

type BasisActions struct {
	ObjectName string
}

func (b *BasisActions) Run() {
	fmt.Printf("%s run\n", b.ObjectName)
}

func (b *BasisActions) Jump() {
	fmt.Printf("%s jump\n", b.ObjectName)
}

type CatActions struct{}

func (c *CatActions) CreateActions(name string) ActionsSet {
	return &Cat{
		&BasisActions{
			ObjectName: name,
		},
	}
}

type Cat struct {
	*BasisActions
}

func (c *Cat) DoSomething() {
	fmt.Printf("%s sleep\n", c.ObjectName)
}

type DogActions struct{}

func (c *DogActions) CreateActions(name string) ActionsSet {
	return &Dog{
		&BasisActions{
			ObjectName: name,
		},
	}
}

type Dog struct {
	*BasisActions
}

func (c *Dog) DoSomething() {
	fmt.Printf("%s play\n", c.ObjectName)
}
