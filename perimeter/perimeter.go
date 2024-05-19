package perimeter

import "math"

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (c Triangle) Area() float64 {
	return 0.5 * c.Base * c.Height
}
