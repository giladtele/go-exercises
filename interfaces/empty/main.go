package main

import "fmt"

type Shape interface {
	Area() float64
}

type Object struct {
	Length float64
}

func (o Object) Area() float64 {
	return o.Length * o.Length * o.Length
}

func identityType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("This is an int:", v)
	case string:
		fmt.Println("This is a string:", v)
	case float64:
		fmt.Println("This is a float64:", v)
	case float32:
		fmt.Println("This is a float32:", v)
	case Shape:
		fmt.Println("This is an interface of type Shape:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	identityType(42)
	identityType("hello")
	identityType(Object{Length: 2})
	identityType(4.2)
}
