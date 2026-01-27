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
	rectangle := Rectangle{ 10.0, 12.0 }
	got := Area(rectangle)
	want := 120.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

