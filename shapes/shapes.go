package shapes

import "math"

type Rectangle struct {
	w float64
	h float64
}
type Circle struct {
	r float64
}

func Perimeter(r Rectangle) float64 {
	return 2*r.h + 2*r.w
}

func (r Rectangle) Area() float64 {
	return r.h * r.w
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

type Shape interface {
	Area() float64
}