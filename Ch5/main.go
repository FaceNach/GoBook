package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sort"
)

func divNumbers(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("Cannot divide by zero")
	}

	return num / denom, num % denom, nil
}

func Testing(myFuncOpts MyFuncOpts) error {
	fmt.Println(myFuncOpts.FirstName)
	fmt.Println(myFuncOpts.LastName)
	fmt.Println(myFuncOpts.Age)

	return nil
}

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}

	return out
}

func returningFunction(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func multiplyNumWithRandom(base int) func() int {
	return func() int {
		rndNumber := rand.Intn(10)
		fmt.Println(rndNumber)
		return rndNumber * base
	}
}

func main() {

	result, rest, err := divNumbers(15, 4)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println(result)
	fmt.Println(rest)

	var myFuncOpts = MyFuncOpts{
		FirstName: "John",
		LastName:  "Doe",
		Age:       18,
	}

	output := Testing(myFuncOpts)
	if output != nil {
		fmt.Println(err)
	}

	fmt.Println("---------------------------------------------")

	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))

	fmt.Println("Anonymous functions")

	anonymousFunction := 20
	f := func() {
		fmt.Println(anonymousFunction)
		anonymousFunction := 30
		fmt.Println(anonymousFunction)
	}
	f()
	fmt.Println(anonymousFunction)

	type Person struct {
		Name     string
		LastName string
		Age      int
	}

	people := []Person{
		{"John", "Doe", 18},
		{"Tracy", "Dylan", 30},
		{"Fred", "Stan", 25},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Println(people)

	twoBase := returningFunction(2)
	threeBase := returningFunction(3)

	for i := 0; i < 5; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}

	rnd := multiplyNumWithRandom(5)

	fmt.Println(rnd())

}
