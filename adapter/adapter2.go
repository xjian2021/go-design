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

func NewSquarePegAdapter(squarePeg SquarePeg) RoundPeg {
	return &SquarePegAdapter{SquarePeg: squarePeg}
}

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

func (r *roundHole) Fits(peg RoundPeg) bool {
	return r.GetRadius() >= peg.GetRadius()
}

type roundPeg struct {
	radius float64
}

func NewRoundPeg(radius float64) *roundPeg {
	return &roundPeg{radius: radius}
}

func (r *roundPeg) GetRadius() float64 {
	return r.radius
}

type squarePeg struct {
	width float64
}

func NewSquarePeg(width float64) *squarePeg {
	return &squarePeg{width: width}
}

func (s *squarePeg) GetWidth() float64 {
	return s.width
}

