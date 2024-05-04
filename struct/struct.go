package _struct

import "math"

type Shape interface {
	Area() float64
}

type Rect struct {
	W float64
	H float64
}

func (rect Rect) Perimeter() float64 {
	return 2 * (rect.W + rect.H)
}
func (rect Rect) Area() float64 {
	return rect.H * rect.W
}

type Circle struct{ R float64 }

func (circle Circle) Perimeter() float64 { return 2 * math.Pi * circle.R }
func (circle Circle) Area() float64 {
	return math.Pi * circle.R * circle.R
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 { return (t.Base * t.Height) * 0.5 }
