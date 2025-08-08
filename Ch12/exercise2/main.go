package main

import (
	"fmt"
	"sync"
)

func callFunc() <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			out <- fmt.Sprintf("The value is %v from first Goroutine", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			out <- fmt.Sprintf("The value is %v from second Goroutine", i)
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	channel := callFunc()

	for v := range channel {
		fmt.Println(v)
	}
}
