package main

import "fmt"

func main() {
	add(1, 2)
	a := 5
	b := 9
	callByRefrence(&a, &b)
	callByValue(5, 6)
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
