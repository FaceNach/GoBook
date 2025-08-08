package main

import (
	"fmt"
	"math"
	"sync"
)

func callFunc() map[int]float64 {

	var a = map[int]float64{}

	for i := 0; i < 100000; i++ {
		a[i] = math.Sqrt(float64(i))
	}

	return a
}

func main() {
	a := sync.OnceValue(callFunc)

	b := a()

	var wg sync.WaitGroup

	wg.Add(1)

	for i := 0; i < 100000; i += 1000 {
		fmt.Printf("sqrt(%d) = %f\n", i, b[i])
	}

}
