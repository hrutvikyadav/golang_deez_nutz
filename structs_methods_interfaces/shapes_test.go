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
	areaTests := []struct{
		shape Shape
		want float64
	}{
		{ Rectangle{ 10.0, 12.0 }, 120.0 },
		{ Circle{ 10.0 }, 314.1592653589793 },
		{ Triangle{ 12, 6.0 }, 36 },
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}
}

