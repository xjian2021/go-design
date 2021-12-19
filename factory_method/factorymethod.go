package factory_method

import "fmt"

type BasisFunction interface {
	Hand()
	Body()
}

type UniqueFunction interface {
	Face()
}

type People interface {
	BasisFunction
	UniqueFunction
}

type PeopleFactory interface {
	Create(bodySize string, handNum int) People
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
		Basis: &Basis{
			bodySize: bodySize,
			handNum:  handNum,
		},
	}
}

type Rng struct {
	*Basis
}

func (r *Rng) Face() {
	fmt.Println("我是rng的脸")
}

type EdgFactory struct{}

func (r *EdgFactory) Create(bodySize string, handNum int) People {
	return &Edg{
		Basis: &Basis{
			bodySize: bodySize,
			handNum:  handNum,
		},
	}
}

type Edg struct {
	*Basis
}

func (r *Edg) Face() {
	fmt.Println("我是edg的脸")
}
