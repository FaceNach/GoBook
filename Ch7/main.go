package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func main() {

	miRectangle := Rectangle{10, 5}

	fmt.Println(miRectangle.Area())

}
