package main

type Rectangle struct {
	Width float64
	Height float64
}

func Perimeter(shape Rectangle) (perimeter float64) {
	perimeter = 2 * (shape.Width + shape.Height)
	return 
}

func Area(shape Rectangle) (area float64) {
	area = shape.Width * shape.Height
	return
}
