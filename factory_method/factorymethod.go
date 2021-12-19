package factory_method

import "fmt"

// 基础功能，也就是每个对象都拥有的技能
type BasisFunction interface {
	Hand()
	Body()
}

// 差异功能，每个对象都有差异的功能
type UniqueFunction interface {
	Face()
}

// 功能总和
type People interface {
	BasisFunction
	UniqueFunction
}

type PeopleFactory interface {
	Create(bodySize string, handNum int) People
}

func CreateFactory(f PeopleFactory, bodySize string, handNum int) People {
	return f.Create(bodySize, handNum)
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

type RngFactory struct{}

func (r *RngFactory) Create(bodySize string, handNum int) People {
	return &Rng{
		&Basis{
			bodySize: bodySize,
			handNum:  handNum,
		},
	}
}

// 如果嵌入的是类型 那么该类型额外实现的方法 Rng内部也可以使用
type Rng struct {
	*Basis
}

//Foot Basis额外实现的方法
func (b *Basis) Foot() {
	fmt.Println("我是脚")
}

func (r *Rng) Face() {
	fmt.Println("我是rng的脸")
	//使用Basis的Foot方法
	r.Foot()
}

type EdgFactory struct{}

func (r *EdgFactory) Create(bodySize string, handNum int) People {
	return &Edg{
		BasisFunction: &Basis{
			bodySize: bodySize,
			handNum:  handNum,
		},
	}
}

// 如果嵌入的是接口 那么就Edg就只能使用接口规定的方法
type Edg struct {
	BasisFunction
}

func (r *Edg) Face() {
	fmt.Println("我是edg的脸")
}
