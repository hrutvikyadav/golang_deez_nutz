package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{ 10.0, 12.0 }
	got := Perimeter(rectangle)
	want := 2 * (10.0 + 12.0)

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea (t *testing.T) {
	t.Run("for rectangles", func(t *testing.T) {
		rectangle := Rectangle{ 10.0, 12.0 }
		got := rectangle.Area()
		want := 120.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("for circles", func(t *testing.T) {
		circle := Circle{ 10.0 }
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}

