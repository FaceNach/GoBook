package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Testing interface {
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

func main() {

	r := Rectangle{10, 5}
	c := Circle{10}

	fmt.Println("Rectangle area: ", calculateArea(r))
	fmt.Println("Circle area", calculateArea(c))

	var a Shape = Rectangle{}

	_, ok := a.(Rectangle)

	if ok {
		fmt.Println("Its a rectangle")
	} else {
		fmt.Println("Its not")
	}

	_, b := a.(Circle)

	if b {
		fmt.Println("Its a circle")
	} else {
		fmt.Println("Its not")
	}

}
