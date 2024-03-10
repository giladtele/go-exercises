package main

import "fmt"

type GeometricOperations interface {
	Area() float64
	Volume() float64
	Perimeter() float64
}

type GeometricShapes struct {
	Height, Width, Radius, Face float64
}

func (g GeometricShapes) Area() float64 {
	if g.Radius == 0 && g.Face == 0 && g.Height != 0 && g.Width != 0 {
		return (g.Height * g.Width)
	} else if g.Radius != 0 && g.Face == 0 && g.Height == 0 && g.Width == 0 {
		return (g.Radius * g.Radius) * 3.14
	} else if g.Face != 0 && g.Radius == 0 && g.Height == 0 && g.Width == 0 {
		return 6 * (g.Face * g.Face)
	} else {
		fmt.Println("Error! Improper values were passed")
		return 0
	}
}

func (g GeometricShapes) Volume() float64 {
	if g.Face != 0 && g.Height == 0 && g.Width == 0 && g.Radius == 0 {
		return g.Face * g.Face * g.Face
	} else {
		fmt.Println("Provided shape does not have the Volume geometrical property. Otherwise, improper values were passed")
		return 0
	}
}

func (g GeometricShapes) Perimeter() float64 {
	if g.Radius != 0 && g.Height == 0 && g.Width == 0 && g.Face == 0 {
		return 2 * (g.Radius * 3.14)
	} else {
		return 0
	}
}

func main() {
	var shapes []GeometricOperations
	var face, rad, len, width float64
	fmt.Printf("Please input values in the following order for:\n1.Face (cube)\n2. Radius (circle)\n3. Height (rectangular)\n4.Width (rectangular)\n")
	fmt.Scanln(&face)
	fmt.Scanln(&rad)
	fmt.Scanln(&len)
	fmt.Scanln(&width)
	shapes = append(shapes, GeometricShapes{Face: face, Radius: rad, Height: len, Width: width})
	for _, shape := range shapes {
		fmt.Println("Area of shape:", shape.Area())
		if shape.Perimeter() == 0 {
			fmt.Println("Provided shape does not have the Perimeter geometrical property. Otherwise, improper values were passed")
		} else {
			fmt.Println("Perimeter of circle is:", shape.Perimeter())
		}
		if shape.Volume() == 0 {
			fmt.Println("Provided shape does not have the Volume geometrical property. Otherwise, improper values were passed")
		} else {
			fmt.Println("Volume of cube is:", shape.Volume())
		}
	}
}
