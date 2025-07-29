package main

import (
	"fmt"
)

func main() {

	a := 10
	b := true

	c := &a
	d := &b

	fmt.Println(*c)
	fmt.Println(c)

	fmt.Println(*d)
	fmt.Println(d)

	var e *int

	num := 35
	e = &num

	*e = 40

	fmt.Println(*e + 5)
	fmt.Println(*e)

}
