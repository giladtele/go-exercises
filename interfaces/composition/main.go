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

type Shp struct {
	Material string
	Size     float64
}

func (c Cube) Area() float64 {
	return 6 * (c.Size * c.Size)
}

func (m Cube) Material() string {
	return m.Substance
}

func (v Cube) Volume() float64 {
	return (v.Size * v.Size * v.Size)
}

func main() {
	cube := Cube{"titanium", 15}
	fmt.Println("Volume of Cube:", cube.Volume())
	fmt.Println("Area of Cube:", cube.Area())
	fmt.Println("Material of Cube:", cube.Material())
}
