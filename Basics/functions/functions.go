package main

import (
	"fmt"
	"strings"
)

func main() {
	add(1, 2)
	a := 5
	b := 9
	callByRefrence(&a, &b)
	callByValue(5, 6)

	// ************* Anonymous function *************

	func(element int) {
		fmt.Println(element)
	}(8)

}

func add(a, b int) int {
	c := a + b
	fmt.Println(c)
	return c
}

func sub(a, b int) int {
	c := a - b
	fmt.Println(c)
	return c
}

// ********** Call by Value ************

func callByValue(a, b int) int {
	c := a
	b = a
	b = c
	return c
}

// ********** Call by Refrence ************
// This is used to change the value of the parameters passed in the main function. The same is not done in call by value.

func callByRefrence(a, b *int) int {
	c := *a
	*b = *a
	*b = c
	return c
}

// *************** Varadic Function **************
// to pass zero or multiple arguments in the function the number or the variables names dont need to be provided.

func varadicFunc(element ...string) string {
	return strings.Join(element, "-")
}

// *************** Init Function ****************
// The min purpse of init function is to run before main function in Golang, apart from this there can be multiple init functions in a program.

func init() {
	fmt.Println("Into Init function")
}

// ************ Underscore Identifier in Golang ************

func Underscore() {
	mulDiv(5, 5) // returns the correct value

	a, b := mulDiv(6, 6)
	println(a, b)

	_, c := mulDiv(7, 7)
	println(c)
}
func mulDiv(a, b int) (int, int) {
	return a * b, a / b
}
