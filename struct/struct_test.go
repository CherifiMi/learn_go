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

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%f %f", got, want)
		}

	}

	t.Run("test rect", func(t *testing.T) {
		rect := Rect{12.0, 6.0}
		checkArea(t, rect, 72.0)
	})
	t.Run("test rect", func(t *testing.T) {
		circle := Circle{
			R: 10.0,
		}
		checkArea(t, circle, 314.1592653589793)
	})
}
