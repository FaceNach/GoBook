package main

import "fmt"

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (l *LinkedList[T]) Add(value T) {
	n := &Node[T]{Value: value}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = l.Tail.Next
}

func (l *LinkedList[T]) Insert(value T, pos int) {
	n := &Node[T]{Value: value}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	if pos == 0 {
		n.Next = l.Head
		l.Head = n
		return
	}

	curNode := l.Head

	for i := 1; i < pos; i++ {
		if curNode.Next == nil {
			curNode.Next = n
			l.Tail = curNode.Next
			return
		}
		curNode = curNode.Next
	}

	n.Next = curNode.Next
	curNode.Next = n
	if l.Tail == curNode {
		l.Tail = n
	}
	return
}

func (l *LinkedList[T]) Index(value T) int {
	index := 0

	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		if curNode.Value == value {
			return index
		}
		index++

	}

	return -1

}

func main() {

	l := &LinkedList[int]{}
	l.Add(5)
	l.Add(10)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))

	l.Insert(100, 0)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	l.Insert(200, 1)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(200))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Insert(300, 10)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Add(400)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Insert(500, 6)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

}
