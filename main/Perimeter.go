package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func Perimeter(rectangle Rectangle) float64 {
	return 2*rectangle.width + 2*rectangle.height
}
