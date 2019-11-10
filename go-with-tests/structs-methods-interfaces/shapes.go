package structs_methods_interfaces

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Height float64
	Base   float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	result := math.Pi * c.Radius * c.Radius
	return math.Round(result*10) / 10
}

func (t Triangle) Area() float64 {
	result := (t.Height * t.Base) / 2
	return result
}
