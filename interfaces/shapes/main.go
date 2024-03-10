package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) Area() float64 {
	return 3.14 * (c.radius * c.radius)
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	var shapes []Shape

	// as we're passing the values like this,
	// it creates a copy of the elements in the `shapes` slice
	// if we wanted to avoid that, we woud've done something like so:
	// `shapes = append(shapes, &Circle{5}, &Rectangle{5, 5})`
	shapes = append(shapes, Circle{5}, Rectangle{5, 5})

	for _, shape := range shapes {
		fmt.Println(shape.Area())
	}
}
