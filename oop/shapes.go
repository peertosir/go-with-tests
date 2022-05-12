package oop

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	a float64
	b float64
	c float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * math.Pi
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) Area() float64 {
	semiPerimeter := t.Perimeter() / 2
	return math.Sqrt(semiPerimeter * (semiPerimeter - t.a) * (semiPerimeter - t.b) * (semiPerimeter - t.c))
}
