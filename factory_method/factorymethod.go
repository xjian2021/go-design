package factory_method

import "fmt"

const (
	RNG_TYPE FactoryType = iota
	EDG_TUPE
)

type FactoryType int

//People 方法集合
type People interface {
	Hand()
	Body()
	Face()
}

type Basis struct {
	handNum  int
	bodySize string
}

func (b *Basis) Hand() {
	fmt.Printf("我有%d只手\n", b.handNum)
}

func (b *Basis) Body() {
	fmt.Printf("我有个%s身体\n", b.bodySize)
}

func (b *Basis) Face() {
	fmt.Println("基础默认脸")
}

type RngFactory struct{}

func (r *RngFactory) Create(bodySize string, handNum int) People {
	return &Rng{
		&Basis{
			bodySize: bodySize,
			handNum:  handNum,
		},
	}
}

type Rng struct {
	*Basis
}

// 可根据需求重载方法
func (r *Rng) Face() {
	fmt.Println("我是rng的脸")
}

type EdgFactory struct{}

func (r *EdgFactory) Create(bodySize string, handNum int) People {
	return &Edg{
		 &Basis{
			bodySize: bodySize,
			handNum:  handNum,
		},
	}
}

type Edg struct {
	*Basis
}

// 可根据需求重载方法
func (e *Edg) Face() {
	fmt.Println("我是edg的脸")
}

type PeopleFactory interface {
	Create(bodySize string, handNum int) People
}

//GetPeople 方式一
//利用工厂模式获取自定义的对象工厂
//优点在于在需要新增业务类型时，只要实现接口即可使用(易于扩展)
func GetPeople(pf PeopleFactory,bodySize string, handNum int) People {
	return pf.Create(bodySize, handNum)
}

//GetPeople2 方式二
//给每种工厂定义标识，在使用时只需要传入标识即可(易于使用)
//优点在于使用者不会直接与具体类型进行互动，而是依靠该构建函数来调用构造方法从而创建接口的实例，仅使用标识参数来控制生产
func GetPeople2(ftype FactoryType, bodySize string, handNum int) People {
	switch ftype {
	case RNG_TYPE:
		return (&RngFactory{}).Create(bodySize, handNum)
	case EDG_TUPE:
		return (&EdgFactory{}).Create(bodySize, handNum)
	}
	return nil
}