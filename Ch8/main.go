package main

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	vals []T
}

func (p *Stack[T]) Push(v T) {

	p.vals = append(p.vals, v)
}

func (p *Stack[T]) Pop() (T, error) {
	if len(p.vals) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}

	top := len(p.vals) - 1
	elem := p.vals[top]
	p.vals = p.vals[:top]

	return elem, nil

}

func Contains[T comparable](slice []T, v T) bool {
	for _, value := range slice {
		if value == v {
			return true
		}
	}

	return false
}

func main() {

	a := Stack[int]{vals: []int{1, 2, 3, 4}}
	b := Stack[string]{vals: []string{"Hola", "Chau", "Testing"}}

	a.Push(5)
	fmt.Println(a)
	v, _ := a.Pop()
	fmt.Printf("The value: %v was poped from %v\n", v, a)

	b.Push("Hola Como Estas?")
	fmt.Println(b)
	b.Pop()
	fmt.Println(b)

	slice1 := []int{1, 2, 3, 4}
	slice2 := []string{"Probando", "Probando2", "Probando3"}

	fmt.Println(Contains(slice1, 3))
	fmt.Println(Contains(slice2, b.vals[0]))

}
