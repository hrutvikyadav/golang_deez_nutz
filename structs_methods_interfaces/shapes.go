package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() (area float64) {
	area = r.Width * r.Height
	return
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() (area float64) {
	area = math.Pi * c.Radius * c.Radius
	return
}

func Perimeter(shape Rectangle) (perimeter float64) {
	perimeter = 2 * (shape.Width + shape.Height)
	return 
}

