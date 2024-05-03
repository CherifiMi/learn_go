package _struct

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("test rect", func(t *testing.T) {
		rect := Rect{
			W: 10,
			H: 10,
		}
		g := rect.Perimeter()
		w := 40.0
		if g != w {
			t.Errorf("error")
		}
	})
	t.Run("test circle", func(t *testing.T) {
		circle := Circle{
			R: 10,
		}
		g := circle.Perimeter()
		w := 62.83185307179586
		if g != w {
			t.Errorf("error")
		}
	})
}
func TestArea(t *testing.T) {
	t.Run("test rect", func(t *testing.T) {
		rect := Rect{
			W: 12.0,
			H: 6.0,
		}
		g := rect.Area()
		w := 72.0
		if g != w {
			t.Errorf("error")
		}
	})
	t.Run("test rect", func(t *testing.T) {
		circle := Circle{
			R: 10.0,
		}
		g := circle.Area()
		w := 314.1592653589793
		if g != w {
			t.Errorf("error")
		}
	})
}
