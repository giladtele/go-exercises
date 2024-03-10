package main

import "fmt"

type Shape interface {
	Area() float64
}

type Solid interface {
	Volume() float64
}

type Material interface {
	Material() string
}

type Cube struct {
	Substance string
	Size      float64
}

func (c Cube) Area() float64 {
	return 6 * (c.Size * c.Size)
}

func (c Cube) Material() string {
	return c.Substance
}

func (c Cube) Volume() float64 {
	return (c.Size * c.Size * c.Size)
}

func main() {
	cube := Cube{"titanium", 15}
	fmt.Println("Volume of Cube:", cube.Volume())
	fmt.Println("Area of Cube:", cube.Area())
	fmt.Println("Material of Cube:", cube.Material())
}
