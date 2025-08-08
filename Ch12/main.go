package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		ch <- 5
	}()

	go func() {
		defer wg.Done()
		ch <- 6
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for results := range ch {
		fmt.Println(results)
	}

	select {
	case test, ok := <-ch:
		if !ok {
			fmt.Println("Channel is closed")
			break
		}
		fmt.Println(test)
	}

	a := []int{1, 2, 3, 4, 5, 6, 8}
	ch4 := make(chan int, len(a))

	for _, vals := range a {
		go func() {
			ch4 <- vals * 2
		}()
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch4)
	}

}
