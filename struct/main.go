package main

import "math"

type Form interface {
	area() float64
	perimeter() float64
}

type Retangle struct {
	height float64
	width  float64
}

func (r Retangle) area() float64 {
	return r.height * r.width
}

func (r Retangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
