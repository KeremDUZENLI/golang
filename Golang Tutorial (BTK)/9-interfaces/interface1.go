package interfaces

import "fmt"

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return c.radius * c.radius * (3.14)
}

func school(s shape) {
	fmt.Printf("Area: %v\n", s.area())
}

func Interface1() {
	r := rectangle{width: 10, height: 6}
	c := circle{radius: 10}
	school(r)
	school(c)
}
