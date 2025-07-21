package main

import (
	"fmt"
	"math"
)

func main() {

	//1. Write a program that declares an integer variable called i with the value 20.
	//Assign i to a floating-point variable named f. Print out i and f.

	//2. Write a program that declares a constant called value that can be assigned to
	//both an integer and a floating-point variable. Assign it to an integer called i and a
	//floating-point variable called f. Print out i and f.

	//3. Write a program with three variables, one named b of type byte, one named
	//smallI of type int32, and one named bigI of type uint64. Assign each variable
	//the maximum legal value for its type; then add 1 to each variable. Print out their
	//values.

	//1

	i := 20
	var f float64

	f = float64(i)

	fmt.Println(f)

	fmt.Println("---------------------------------------------------\n")

	const value = 10

	var t int = value
	var j float64 = value

	fmt.Println(value)
	fmt.Printf("%T, %v\n \n", t, t)
	fmt.Printf("%T, %v\n \n", j, j)

	fmt.Println("---------------------------------------------------\n")

	var (
		b      byte
		smallI int32
		bigI   uint64
	)

	b = math.MaxInt8
	smallI = math.MaxInt32
	bigI = math.MaxUint64

	fmt.Printf("%T, %v \n", b, b)
	fmt.Printf("%T, %v \n", smallI, smallI)
	fmt.Printf("%T, %v \n", bigI, bigI)

	fmt.Printf("%T, %v, %v \n", b, b, b+1)
	fmt.Printf("%T, %v, %v \n", smallI, smallI, smallI+1)
	fmt.Printf("%T, %v, %v \n", bigI, bigI, bigI+1)

}
