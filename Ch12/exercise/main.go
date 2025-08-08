package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func callFunction() <-chan int {
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan int)

	out := make(chan int)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i * rand.Intn(10)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i * rand.Intn(10)
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	go func() {
		for v := range ch {
			out <- v
		}
		close(out)
	}()

	return out
}

func main() {

	channel := callFunction()

	for v := range channel {
		fmt.Println(v)
	}

}
