package adapter

import "math"

//RoundHole 圆孔接口
type RoundHole interface {
	RoundPeg
	Fits(RoundPeg) bool
}

//RoundPeg 圆钉接口
type RoundPeg interface {
	GetRadius() float64
}

//SquarePeg 方钉接口
type SquarePeg interface {
	GetWidth() float64
}

//SquarePegAdapter 方钉适配器类
//会让方钉也适用于圆孔
type SquarePegAdapter struct {
	SquarePeg
}

//NewSquarePegAdapter 在适配器类中添加一个成员变量用于保存对于服务对象的引用。 通常情况下会通过构造函数对该成员变量进行初始化， 但有时在调用其方法时将该变量传递给适配器会更方便。
func NewSquarePegAdapter(squarePeg SquarePeg) RoundPeg {
	return &SquarePegAdapter{SquarePeg: squarePeg}
}

//GetRadius 依次实现适配器类客户端接口的所有方法。 适配器会将实际工作委派给服务对象， 自身只负责接口或数据格式的转换。
func (s *SquarePegAdapter) GetRadius() float64 {
	return s.GetWidth() * math.Sqrt(2) / 2
}

//roundHole 圆孔类
type roundHole struct {
	radius float64
}

func NewRoundHole(radius float64) RoundHole {
	return &roundHole{radius: radius}
}

func (r *roundHole) GetRadius() float64 {
	return r.radius
}

//Fits 客户端必须通过客户端接口(指接口类型)使用适配器。这样一来， 你就可以在不影响客户端代码的情况下修改或扩展适配器。
func (r *roundHole) Fits(peg RoundPeg) bool {
	return r.GetRadius() >= peg.GetRadius()
}

type roundPeg struct {
	radius float64
}

func NewRoundPeg(radius float64) RoundPeg {
	return &roundPeg{radius: radius}
}

func (r *roundPeg) GetRadius() float64 {
	return r.radius
}

type squarePeg struct {
	width float64
}

func NewSquarePeg(width float64) SquarePeg {
	return &squarePeg{width: width}
}

func (s *squarePeg) GetWidth() float64 {
	return s.width
}

