package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	var x = []int{1, 2, 3, 4}
	var s = []int{1, 4: 5, 20, 35, 25: 25, 30, 40}
	fmt.Println(x)
	fmt.Println(s)

	var t []int

	fmt.Println(t)

	fmt.Println(slices.Equal(x, s)) //compares two slices
	fmt.Println(len(x))             //length of slice
	x = append(x, 5, 6, 7, 8)
	fmt.Println(x)

	x = append(x, s...) //kinds of works like destructuration in JS ??
	fmt.Println(x)
	fmt.Println(cap(x))

	u := make([]int, 5, 10)

	fmt.Println(len(u))
	fmt.Println(cap(u))
	u = append(u, 5, 6, 7, 8, 9, 10)
	fmt.Println(u)
	clear(u)
	fmt.Println(u)

	y := x[:2] //sharing memory,
	num := copy(x, y)
	fmt.Println(num)

	fmt.Println(y)

	testing := make([]string, 0, 5)
	testing = append(testing, "a", "b", "c", "d")
	secondTesting := testing[:2]
	thirdTesting := testing[2:]
	fmt.Println(cap(testing), cap(secondTesting), cap(thirdTesting))

	fourthtesting := testing[2:4:4]

	fmt.Println(testing)
	fmt.Println(secondTesting)
	fmt.Println(thirdTesting)
	fmt.Println(fourthtesting)

	fmt.Println("-------------------Maps--------------------")

	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)

	m := map[string]int{
		"hello": 5,
		"world": 10,
	}

	n := map[string]int{
		"hello": 10,
		"world": 5,
	}

	fmt.Println(m)

	fmt.Println(maps.Equal(m, n)) // used to compare maps

	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, r := range vals {
		intSet[r] = true
		fmt.Printf("key: %v bool: %v\n", vals[r], intSet[r])

	}

	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

	//method to create a set with maps in go. Basically its takes the key on the map and assigns the value true to every key
	//if we have duplicates this algorithm overwrites the key and the bool value. That's why the length of vals != intSet
	// You cannot have duplicated keys in a map.

	fmt.Println("------------------------------------------------------")

	type person struct {
		name string
		age  int
		pet  string
	}

	julia := person{
		"julia",
		40,
		"cat",
	}

	julia.name = "Julia"

	fmt.Println(julia.name)

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "Dog",
	}

	fmt.Println(pet)

	//Write a program that defines a variable named greetings of type slice of
	//strings with the following values: "Hello", "Hola", "नमस्कार", "こんにちは",
	//and "Привіт". Create a subslice containing the first two values; a second subslice
	//with the second, third, and fourth values; and a third subslice with the fourth and
	//fifth values. Print out all four slices.

	var firstSlice = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	secondSlice := firstSlice[:2]
	thirdSlice := firstSlice[2:4]
	fourthSlice := firstSlice[3:5]

	fmt.Println(firstSlice)
	fmt.Println(secondSlice)
	fmt.Println(thirdSlice)
	fmt.Println(fourthSlice)

	//	3. Write a program that defines a struct called Employee with three fields:
	//		firstName, lastName, and id. The first two fields are of type string, and the
	//		last field (id) is of type int. Create three instances of this struct using whatever
	//		values you’d like. Initialize the first one using the struct literal style without
	//		names, the second using the struct literal style with names, and the third with a
	//		var declaration. Use dot notation to populate the fields in the third struct. Print
	//		out all three structs

	type Employee struct {
		firstName  string
		secondName string
		id         int
	}

	employee1 := Employee{
		firstName:  "John",
		secondName: "Smith",
		id:         1,
	}

	employee2 := Employee{
		"Jane",
		"Doe",
		2,
	}

	var employee3 = Employee{}

	employee3.firstName = "Mike"
	employee3.secondName = "Peter"
	employee3.id = 3

	fmt.Println(employee1)
	fmt.Println(employee2)
	fmt.Println(employee3)
}
