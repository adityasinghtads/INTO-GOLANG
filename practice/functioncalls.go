package main

import "fmt"

func main() {
	var firstReturn int
	firstReturn = first()

	fmt.Println("THis is first return ")
	fmt.Println(firstReturn)

	youth := second(firstReturn)
	fmt.Println(youth)

}

func first() int {
	return 1
}

func second(firstReturn int) int {
	firstReturn = 10
	return firstReturn
}
