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
		name string
		shape Shape
		hasArea float64
	}{
		{ name: "Rectangle", shape: Rectangle{ Width: 10.0, Height: 12.0 }, hasArea: 120.0 },
		{ name: "Circle", shape: Circle{ Radius: 10.0 }, hasArea: 314.1592653589793 },
		{ name: "Triangle", shape: Triangle{ Base: 12, Height: 6.0 }, hasArea: 36 },
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v - got %.2f want %.2f", tt.shape, got, tt.hasArea)
		}
	}
}

