package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(10)
	fmt.Println(n)

	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big")
	} else {
		fmt.Println("That's a good number")
	}

	if m := rand.Intn(10); m == 0 {
		fmt.Println("Thats too low")
	} else if m > 5 {
		fmt.Println("Thats too big")
	} else {
		fmt.Println("Thats a good number")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for j := 0; j < 10; {
		fmt.Println(j)
		j++
	}

	evenVals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	for _, v := range evenVals {
		fmt.Println(v)
	}

	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}

	for k := range uniqueNames {
		fmt.Println(k)
	}

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for i := 0; i < len(m); i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	samples := []string{"Hello", "apple:n!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
	}

	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				break
			}
		}
	}

	var words = []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}

	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, " is too short")
		case 5:
			fmt.Println(word, " is exactly the length", len(word))
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word")
		}
	}
}
