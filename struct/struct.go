package _struct

import "math"

type Rect struct {
	W float64
	H float64
}

func (rect *Rect) Perimeter() float64 {
	return 2 * (rect.W + rect.H)
}
func (rect *Rect) Area() float64 {
	return rect.H * rect.W
}

type Circle struct {
	R float64
}

func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.R
}
func (circle Circle) Area() float64 {
	return math.Pi * circle.R * circle.R
}
