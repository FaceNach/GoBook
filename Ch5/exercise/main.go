package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

//1. The simple calculator program doesn’t handle one error case: division by zero.
//Change the function signature for the math operations to return both an int and
//an error. In the div function, if the divisor is 0, return errors.New("division
//by zero") for the error. In all other cases, return nil. Adjust the main function
//to check for this error.

//2. Write a function called fileLen that has an input parameter of type string and
//returns an int and an error. The function takes in a filename and returns the
//number of bytes in the file. If there is an error reading the file, return the error.
//Use defer to make sure the file is closed properly.

//3. Write a function called prefixer that has an input parameter of type string
//and returns a function that has an input parameter of type string and returns a
//string. The returned function should prefix its input with the string passed into
//prefixer. Use the following main function to test prefixer:

//func main() {
//	helloPrefix := prefixer("Hello")
//	fmt.Println(helloPrefix("Bob")) // should print Hello Bob
//	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
//}

func division(num, denom int) (int, error) {
	if denom == 0 {
		divideByZero := errors.New("you cannot divide by zero")
		return 0, errors.New(divideByZero.Error())
	}

	return num / denom, nil
}

func fileLen(filename string) (int, error) {

	if filename == "" {
		fmt.Println("the file name is empty")
		return 0, nil
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	defer file.Close()

	data := make([]byte, 512)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return count, nil

}

func prefixer(word string) func(string) string {
	return func(s string) string {
		return word + " " + s
	}
}

func main() {

	a, b := 9, 3
	c, d := 5, 0
	fmt.Println(a, b)
	fmt.Println(c, d)

	result, _ := division(a, b)
	fmt.Println(result)

	_, err := division(c, d)
	if err != nil {
		fmt.Println(err)
	}

	length, _ := fileLen("justText.txt")
	fmt.Println(length)

	testing, err := fileLen("")
	fmt.Println(testing, err)

	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}
