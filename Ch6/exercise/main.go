//1. Create a struct named Person with three fields: FirstName and LastName of
//type string and Age of type int. Write a function called MakePerson that
//takes in firstName, lastName, and age and returns a Person. Write a second
//function MakePersonPointer that takes in firstName, lastName, and age and
//returns a *Person. Call both from main. Compile your program with go build
//-gcflags="-m". This both compiles your code and prints out which values
//escape to the heap. Are you surprised about what escapes?

//2. Write two functions. The UpdateSlice function takes in a []string and a
//string. It sets the last position in the passed-in slice to the passed-in string. At
//the end of UpdateSlice, print the slice after making the change. The GrowSlice
//function also takes in a []string and a string. It appends the string onto the
//slice. At the end of GrowSlice, print the slice after making the change. Call these
//functions from main. Print out the slice before each function is called and after
//each function is called. Do you understand why some changes are visible in main
//and why some changes are not?

//3. Write a program that builds a []Person with 10,000,000 entries (they could all
//be the same names and ages). See how long it takes to run. Change the value
//of GOGC and see how that affects the time it takes for the program to complete.
//Set the environment variable GODEBUG=gctrace=1 to see when garbage collections
//happen and see how changing GOGC changes the number of garbage collections.
//What happens if you create the slice with a capacity of 10,000,000?

package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {

	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

func UpdateSlice(slice []string, word string) []string {
	position := len(slice)
	slice[position-1] = word

	fmt.Println(slice)
	return slice
}

func GrowSlice(slice []string, word string) []string {
	slice = append(slice, word)
	fmt.Println(slice)
	return slice
}
func main() {

	firstName, firstlastName := "John", "Smith"
	firstAge := 25

	secondName, secondLastName := "Evan", "Doe"
	secondAge := 30

	firstPerson := MakePerson(firstName, firstlastName, firstAge)
	secondPerson := MakePersonPointer(secondName, secondLastName, secondAge)

	fmt.Println(firstPerson)
	fmt.Println(secondPerson)

	fmt.Println("------------------------------------------")

	slice1 := []string{"Hello", "From", "Slice1"}
	slice2 := []string{"Bye", "From", "Slice2"}

	fmt.Println(slice1)
	fmt.Println(slice2)

	word1 := "Word1"
	word2 := "Word2"

	fmt.Println(UpdateSlice(slice1, word1))
	fmt.Println(slice1)

	fmt.Println("---------------------")

	fmt.Println(GrowSlice(slice2, word2))
	fmt.Println(slice2)

}
