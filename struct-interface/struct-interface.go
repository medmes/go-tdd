package struct_interface

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

// Rectangle methods, Rectangle implement Shape interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle methods
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle methods
func (t Triangle) Area() float64 {
	return (t.Height * t.Base) * 0.5
}
