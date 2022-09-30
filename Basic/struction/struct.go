package struction

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Length float64
	Height float64
}

type Shape interface {
	Area() float64
}

func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

func Area(width, height float64) float64 {
	return width * height
}

func RectanglePerimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func RectangleArea(rectangle Rectangle) float64 {
	return (rectangle.Width * rectangle.Height)
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (t Triangle) Area() float64 {
	return t.Height * t.Length / 2
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
