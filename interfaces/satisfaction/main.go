package main

import "fmt"

type Fulfill interface {
	Name() string
}

type Object struct {
	Identifier string
}

func (o Object) Name() string {
	return o.Identifier
}

func isSatisfactory(i interface{}) {
	switch v := i.(type) {
	case Fulfill:
		fmt.Printf("%+v - The passed parameter implements the correct interface successfully\n", v)
	default:
		fmt.Printf("%+v - Passed parameter does NOT implement the correct interface\n", v)
	}
}

func main() {
	var full []Fulfill
	full = append(full, Object{"John"})
	for _, fulls := range full {
		isSatisfactory(fulls)
	}
	isSatisfactory(22)
}
