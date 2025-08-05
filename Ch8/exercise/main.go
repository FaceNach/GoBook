//1. Write a generic function that doubles the value of any integer or float that’s
//passed in to it. Define any needed generic interfaces.

//2. Define a generic interface called Printable that matches a type that implements
//fmt.Stringer and has an underlying type of int or float64. Define types that
//meet this interface. Write a function that takes in a Printable and prints its value
//to the screen using fmt.Println.

//3. Write a generic singly linked list data type. Each element can hold a comparable
//value and has a pointer to the next element in the list. The methods to implement
//are as follows:
//// adds a new element to the end of the linked list
//Add(T)
//// adds an element at the specified position in the linked list
//Insert(T, int)
//// returns the position of the supplied value, -1 if it's not present
//Index (T) int

package main

import (
	"fmt"
)

type GetDouble interface {
	~int | ~float64
}

func DoublesValue[T GetDouble](value T) T {
	return value * 2
}

//----------------------------------------------------------------------------------------//

type myInt int

type myFloat float64

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

func JustPrint[T Printable](value T) {
	fmt.Println(value)
}

func (i myInt) String() string {
	return fmt.Sprintf("%d", i)
}

func (f myFloat) String() string {
	return fmt.Sprintf("%0.1f", f)
}

func main() {

	a := 4
	b := 5.5

	t := myInt(5)
	c := myFloat(17.5)

	fmt.Println(DoublesValue(a))
	fmt.Println(DoublesValue(b))

	JustPrint(t)
	JustPrint(c)
}
