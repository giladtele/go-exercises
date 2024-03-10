package main

type Shape interface {
	Area() float64
}

type Embed interface {
	Shape
}

type Cube struct {
	Face float64
}

func (c Cube) Area() float64 {
	return 6 * (c.Face * c.Face)
}

func main() {
	var cubixes []Embed
	cubixes = append(cubixes, Cube{12})
	for _, cubix := range cubixes {
		cubix.Area()
	}
}
